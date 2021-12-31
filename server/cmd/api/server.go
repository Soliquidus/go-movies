package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"server/models"
)

type Storage interface {
	Open() error
	Close() error

	GetMovie(ctx *gin.Context)
	GetMovies(ctx *gin.Context)
	CreateMovie(ctx *gin.Context)
}

var db *sqlx.DB
var err error

// Open the DB connection
func Open() {
	db, err = sqlx.Open("mysql", "user:password@tcp(host)/DBName?parseTime=true")
	if err != nil {
		log.Println(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Println(err.Error())
	}

	log.Println("Connected to DB")
}

// Close the DB connection
func Close() error {
	log.Println("Connection closed")
	return db.Close()
}

// GetMovie gets one movie with given id and its genres, if any
func GetMovie(ctx *gin.Context) {
	vars := ctx.Params.ByName("id")
	var movie = &models.Movie{}

	err = db.Get(movie, "SELECT * FROM movie WHERE id=?", vars)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot find given id"})
	}
	query := `select
				mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
				movie_genre mg
				left join genre g on (g.id = mg.genre_id)
			where
				mg.movie_id = ?
	`
	rows, _ := db.QueryContext(ctx, query, vars)
	defer rows.Close()

	genres := make(map[int]string)
	for rows.Next() {
		var mg models.MovieGenre
		err = rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return
		}
		genres[mg.ID] = mg.Genre.GenreName
	}

	movie.MovieGenre = genres

	ctx.JSON(http.StatusOK, gin.H{"movie": movie})
}

func GetMovies(ctx *gin.Context) {
	query := `SELECT * FROM movie`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	var movies []*models.Movie
	for rows.Next() {
		var movie models.Movie
		err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return
		}
		genreQuery := `select
				mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
				movie_genre mg
				left join genre g on (g.id = mg.genre_id)
			where
				mg.movie_id = ?
	`
		genreRows, _ := db.QueryContext(ctx, genreQuery, movie.ID)

		genres := make(map[int]string)
		for genreRows.Next() {
			var mg models.MovieGenre
			err = genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName,
			)
			if err != nil {
				return
			}
			genres[mg.ID] = mg.Genre.GenreName
		}
		genreRows.Close()
		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}
	ctx.JSON(http.StatusOK, gin.H{"movies": movies})
}

func deleteMOvie(ctx *gin.Context) {

}

func insertMovie(ctx *gin.Context) {

}

func updateMovie(ctx *gin.Context) {

}

func searchMovies(ctx *gin.Context) {

}
