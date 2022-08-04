package main

import (
	"awesomeProject/models"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
)

const dbUrl = "postgres://postgres:root@localhost:5432/movies_db_sqlboiler_demo?sslmode=disable"

type Genre struct {
	Name string
}

type Movie struct {
	Id     int
	Title  string
	Genres []Genre
}

func connectToDb() *sql.DB {
	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func findMovies(ctx context.Context, db *sql.DB, limit int) {
	log.Print("--- findMovies query --- \n\n")
	m, err := models.Movies(qm.Limit(limit)).All(ctx, db)

	if err != nil {
		log.Fatal(err)
	}

	for _, movie := range m {
		log.Println(movie.Title.String)
	}

	log.Print("\n\n--- End of findMovies query ---")
}

func findGenres(ctx context.Context, db *sql.DB, limit int) {
	log.Print("--- findGenres query --- \n\n")
	g, err := models.Genres(qm.Limit(limit)).All(ctx, db)

	if err != nil {
		log.Fatal(err)
	}

	for _, genre := range g {
		log.Println(genre.Name.String)
	}

	log.Print("\n\n--- End of findGenres query ---")
}

func findMovieByIdWithGenres(ctx context.Context, db *sql.DB, id int) {
	log.Print("--- findMovieByIdWithGenres query --- \n\n")

	movie, err := models.Movies(
		models.MovieWhere.ID.EQ(id),
		qm.Load(qm.Rels(models.MovieRels.MovieGenres, models.MovieGenreRels.Genre)),
	).One(ctx, db)

	if err != nil {
		log.Fatal(err)
	}

	// for _, movie := range m {
	fmt.Printf("Title: %s", movie.Title.String)
	fmt.Println("Categories: ")
	tmpMovie := Movie{Id: movie.ID, Title: movie.Title.String, Genres: make([]Genre, 0)}
	var tmpGenres []Genre
	for _, movieGenre := range movie.R.MovieGenres {
		fmt.Printf("%s ", movieGenre.R.Genre.Name.String)
		tmpGenres = append(tmpGenres, Genre{movieGenre.R.Genre.Name.String})
	}

	tmpMovie.Genres = tmpGenres

	fmt.Println("\n", tmpMovie)

	log.Print("\n\n--- End of findMovieByIdWithGenres query ---")
}

func main() {
	db := connectToDb()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	boil.DebugMode = true

	findMovies(ctx, db, 5)
	findGenres(ctx, db, 5)
	findMovieByIdWithGenres(ctx, db, 8)
}
