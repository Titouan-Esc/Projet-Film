package container

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	databases "web-service/api/database"
	movies "web-service/api/src/domain/movie"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type IServiceContainer interface {
	InjectMovieController() movies.MovieController
}

type kernel struct{}

func (k *kernel) InjectMovieController() movies.MovieController {
	godotenv.Load()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable search_path=public", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	sqlConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error InjectMovieController { %v }", err.Error())
	}

	postgresHandler := &databases.PostgresHandler{}
	postgresHandler.Conn = sqlConn

	movieRepository := &movies.MovieRepository{postgresHandler}
	movieController := movies.MovieController{movieRepository}

	return movieController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
