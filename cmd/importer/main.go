package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/lib/pq"

	"io"
	"log"
	"os"
	"strings"
)

var dbUrl = "postgres://postgres:root@localhost:5432/movies_db_sqlboiler_demo?sslmode=disable"

func fail(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	adult = iota
	belongs_to_collection
	budget
	genres
	homepage
	id
	imdb_id
	original_language
	original_title
	overview
	popularity
	poster_path
	production_companies
	production_countries
	release_date
	revenue
	runtime
	spoken_languages
	status
	tagline
	title
	video
	vote_average
	vote_count
)

type RawGenre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var ErrIsDuplicate error = errors.New("pq: duplicate key value violates unique constraint")

func setGenreToMovie(ctx context.Context, tx *sql.Tx, movieId int, genreId int) error {
	rows, err := tx.QueryContext(ctx, `INSERT INTO movie_genres (movie_id, genre_id) VALUES ($1, $2)`, movieId, genreId)
	defer rows.Close()
	return err
}

func findGenreIdByName(ctx context.Context, tx *sql.Tx, name string) (int, error) {
	id := 0
	rows, err := tx.QueryContext(ctx, `SELECT id FROM genres WHERE name = $1 LIMIT 1`, name)
	defer rows.Close()

	if err != nil {
		return 0, err
	}

	for rows.Next() {
		rows.Scan(&id)
	}

	return id, nil
}

func insertGenres(ctx context.Context, tx *sql.Tx, row []string, movieId int) error {
	var rawGenres []RawGenre
	parsedGenres := strings.ReplaceAll(row[genres], "'", `"`)

	if err := json.Unmarshal([]byte(parsedGenres), &rawGenres); err != nil {
		return err
	}

	for _, value := range rawGenres {
		existentGenreId, err := findGenreIdByName(ctx, tx, value.Name)

		if err != nil {
			return err
		}

		if existentGenreId != 0 {
			log.Printf("Found duplicate genre %s. Skipping", value.Name)

			err = setGenreToMovie(ctx, tx, movieId, existentGenreId)

			if err != nil {
				return err
			}
		} else {
			id := 0
			rows, err := tx.QueryContext(ctx, `INSERT INTO genres (name) VALUES ($1) RETURNING id`, value.Name)

			for rows.Next() {
				rows.Scan(&id)
			}
			rows.Close()

			err = setGenreToMovie(ctx, tx, movieId, id)
			return err
		}
	}

	return nil
}

func insertMovie(ctx context.Context, tx *sql.Tx, row []string) (int, error) {
	id := 0
	query := `INSERT INTO movies (title, is_adult) VALUES ($1, $2) RETURNING id`
	rows, err := tx.QueryContext(ctx, query, row[title], row[adult])
	defer rows.Close()

	if err != nil {
		return 0, err
	}

	for rows.Next() {
		rows.Scan(&id)
	}

	return id, err
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := sql.Open("postgres", dbUrl)
	tx, txErr := db.BeginTx(ctx, &sql.TxOptions{})

	if txErr != nil {
		log.Fatalf("could not beging transaction", err)
	}

	workdir, err := os.Getwd()
	fail(err)

	file, err := os.Open(fmt.Sprintf("%s/dataset/movies_metadata.csv", workdir))
	defer file.Close()
	fail(err)

	reader := csv.NewReader(file)
	currentLine := 0

	for {
		row, err := reader.Read()

		if err != nil || err == io.EOF {
			fail(err)
		}

		if currentLine == 10000 {
			break
		}

		if currentLine > 0 {
			movieId, err := insertMovie(ctx, tx, row)
			if err != nil {
				tx.Rollback()
				log.Fatalf(err.Error())
			}

			err = insertGenres(ctx, tx, row, movieId)

			if err != nil {
				tx.Rollback()
				log.Fatalf(err.Error())
			}

			log.Print("Processed movie id", movieId)
		}

		currentLine++
	}

	log.Print("About to commit transaction")
	err = tx.Commit()

	if err != nil {
		log.Fatalf("shit. could not commit due to %v", err)
	}
	log.Print("Committed successfully")
}
