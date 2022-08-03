// Code generated by SQLBoiler 4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Genres", testGenres)
	t.Run("MovieGenres", testMovieGenres)
	t.Run("Movies", testMovies)
}

func TestDelete(t *testing.T) {
	t.Run("Genres", testGenresDelete)
	t.Run("MovieGenres", testMovieGenresDelete)
	t.Run("Movies", testMoviesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Genres", testGenresQueryDeleteAll)
	t.Run("MovieGenres", testMovieGenresQueryDeleteAll)
	t.Run("Movies", testMoviesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Genres", testGenresSliceDeleteAll)
	t.Run("MovieGenres", testMovieGenresSliceDeleteAll)
	t.Run("Movies", testMoviesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Genres", testGenresExists)
	t.Run("MovieGenres", testMovieGenresExists)
	t.Run("Movies", testMoviesExists)
}

func TestFind(t *testing.T) {
	t.Run("Genres", testGenresFind)
	t.Run("MovieGenres", testMovieGenresFind)
	t.Run("Movies", testMoviesFind)
}

func TestBind(t *testing.T) {
	t.Run("Genres", testGenresBind)
	t.Run("MovieGenres", testMovieGenresBind)
	t.Run("Movies", testMoviesBind)
}

func TestOne(t *testing.T) {
	t.Run("Genres", testGenresOne)
	t.Run("MovieGenres", testMovieGenresOne)
	t.Run("Movies", testMoviesOne)
}

func TestAll(t *testing.T) {
	t.Run("Genres", testGenresAll)
	t.Run("MovieGenres", testMovieGenresAll)
	t.Run("Movies", testMoviesAll)
}

func TestCount(t *testing.T) {
	t.Run("Genres", testGenresCount)
	t.Run("MovieGenres", testMovieGenresCount)
	t.Run("Movies", testMoviesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Genres", testGenresHooks)
	t.Run("MovieGenres", testMovieGenresHooks)
	t.Run("Movies", testMoviesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Genres", testGenresInsert)
	t.Run("Genres", testGenresInsertWhitelist)
	t.Run("MovieGenres", testMovieGenresInsert)
	t.Run("MovieGenres", testMovieGenresInsertWhitelist)
	t.Run("Movies", testMoviesInsert)
	t.Run("Movies", testMoviesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("MovieGenreToGenreUsingGenre", testMovieGenreToOneGenreUsingGenre)
	t.Run("MovieGenreToMovieUsingMovie", testMovieGenreToOneMovieUsingMovie)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("GenreToMovieGenres", testGenreToManyMovieGenres)
	t.Run("MovieToMovieGenres", testMovieToManyMovieGenres)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("MovieGenreToGenreUsingMovieGenres", testMovieGenreToOneSetOpGenreUsingGenre)
	t.Run("MovieGenreToMovieUsingMovieGenres", testMovieGenreToOneSetOpMovieUsingMovie)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("GenreToMovieGenres", testGenreToManyAddOpMovieGenres)
	t.Run("MovieToMovieGenres", testMovieToManyAddOpMovieGenres)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Genres", testGenresReload)
	t.Run("MovieGenres", testMovieGenresReload)
	t.Run("Movies", testMoviesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Genres", testGenresReloadAll)
	t.Run("MovieGenres", testMovieGenresReloadAll)
	t.Run("Movies", testMoviesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Genres", testGenresSelect)
	t.Run("MovieGenres", testMovieGenresSelect)
	t.Run("Movies", testMoviesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Genres", testGenresUpdate)
	t.Run("MovieGenres", testMovieGenresUpdate)
	t.Run("Movies", testMoviesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Genres", testGenresSliceUpdateAll)
	t.Run("MovieGenres", testMovieGenresSliceUpdateAll)
	t.Run("Movies", testMoviesSliceUpdateAll)
}
