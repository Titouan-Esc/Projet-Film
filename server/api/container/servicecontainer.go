package container

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	databases "web-service/api/database"
	movies "web-service/api/src/domain/movie"

	"github.com/joho/godotenv"
)

type IServiceContainer interface {
	InjectMovieController() movies.MovieController
}

type kernel struct{}

func (k *kernel) InjectMovieController() movies.MovieController {
	godotenv.Load()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s"+"password=%s dbname=%s sslmode=disable search_path=public", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_User"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	sqlConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error in connexion with the DB { %v }", sqlConn)
	}

	postgresHandler := &databases.PostgresHandler{}
	postgresHandler.Conn = sqlConn

	movieRepository := &movies.MovieRepository{postgresHandler}
	movieController := movies.MovieController{movieRepository}

	return movieController
}
