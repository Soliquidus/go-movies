package models

import "time"

// Models is the wrapper for database

// Movie is the type for movies
type Movie struct {
	ID          int        `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Year        int        `json:"year" db:"year"`
	ReleaseDate time.Time  `json:"release_date" db:"release_date"`
	Runtime     int        `json:"runtime" db:"runtime"`
	Rating      int        `json:"rating" db:"rating"`
	MPAARating  string     `json:"mpaa_rating" db:"mpaa_rating"`
	CreatedAt   time.Time  `json:"-" db:"created_at"`
	UpdatedAt   time.Time  `json:"-" db:"updated_at"`
	MovieGenre  map[int]string `json:"genres" db:"-"`
}

// Genre is the type for genre
type Genre struct {
	ID        int       `json:"-" db:"id"`
	GenreName string    `json:"genre_name" db:"genre_name"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

// MovieGenre is the type for movie genre
type MovieGenre struct {
	ID        int       `json:"-" db:"id"`
	MovieID   int       `json:"-" db:"movie_id"`
	GenreID   int       `json:"-" db:"genre_id"`
	Genre     Genre     `json:"genre" db:"genre"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}
