package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"io"
	"log"
	"os"
	"strconv"
)

var dbUrl = "postgres://postgres:root@localhost:5432/movies_db_sqlboiler_demo?sslmode=disable"

func fail(err error) {
	if err != nil {
		panic(err)
	}
}

func setup() {
	fmt.Println("...")

	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		panic(err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	workdir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	if err := goose.Up(db, fmt.Sprintf("%s/sql/migrations", workdir)); err != nil {
		panic(err)
	}
	if err := goose.Up(db, fmt.Sprintf("%s/sql/seeds", workdir)); err != nil {
		panic(err)
	}
}

func old() {
	workdir, err := os.Getwd()
	fail(err)

	file, err := os.Open(fmt.Sprintf("%s/dataset/movies_metadata.csv", workdir))
	defer file.Close()
	fail(err)

	scanner := bufio.NewScanner(file)
	currentLine := 0

	for scanner.Scan() {
		if currentLine == 2 {
			break
		}

		fmt.Println(scanner.Text())
		currentLine++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("err at %v: %v", currentLine, err)
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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := sql.Open("postgres", dbUrl)

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

		if currentLine == 1000 {
			break
		}

		if currentLine > 0 {
			query := `INSERT INTO movies (title, is_adult) VALUES ($1, $2)`
			result, err := db.QueryContext(ctx, query, row[title], row[adult])

			if err != nil {
				result.Close()
				fail(fmt.Errorf("failed at line %s due to: %s", strconv.Itoa(currentLine), err.Error()))
			}

			result.Close()
			// fmt.Printf("Movie %s processed", row[title])
		}

		// fmt.Println("title", row[title])

		// fmt.Println("row col", row[0])
		/* for column := range row {
			fmt.Println()
		}*/

		currentLine++
	}
}
