package main

import (
	"awesomeProject/models"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
)

const dbUrl = "postgres://postgres:root@localhost:5432/movies_db_sqlboiler_demo?sslmode=disable"

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

func main() {
	db := connectToDb()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	findMovies(ctx, db, 5)
	findGenres(ctx, db, 5)
}
