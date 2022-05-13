package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	chi "web-service/api/router"
	"web-service/api/src/domain/movie/models"
	"web-service/api/src/utils"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func recoverData() {
	var db = &sql.DB{}
	godotenv.Load()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable search_path=public", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error InjectMovieController { %v }", err.Error())
	}

	// Appel de ma fonction pour récupérer les films sur l'api
	pageNb := 1
	response, err := utils.ConsumApi(fmt.Sprintf(`http://api.themoviedb.org/3/discover/movie?page=%d&api_key=`, pageNb))
	if err != nil {
		fmt.Println(err.Error())
	}

	// Unmarshal les données récupérer de l'api dans un json formaté
	var genre models.DiscoverModel
	if err := json.Unmarshal([]byte(response), &genre); err != nil {
		fmt.Println(err.Error())
	}

	// Créer un boucle pour récupérer les films un à un
	for _, value := range genre.Results {

		// Parse overview et title car pour l'insertion en db les 's pose un gros problème et sur postgreSQL pour remédier à ce problème il faut doubler '
		overview := strings.Replace(value.Overview, "'", "''", -1)
		title := strings.Replace(value.Title, "'", "''", -1)
		// Insérer les films dans la base de donnée
		// var errorFunc func(handler *Ouai, value models.ResultsModel, overview, title string) error = (*Ouai).SaveMovieDB
		// if errorFunc(handler, value, overview, title) != nil {
		// 	fmt.Println(err)
		// }
		row, err := db.Query(fmt.Sprintf(`INSERT INTO "movies" ("id", "poster_path", "overview", "release_date", "title", "popularity", "likes", "dislikes", "comments") VALUES (%d, '%s', '%s', '%s', '%s', %e, 0, 0, '')`, value.ID, value.PosterPath, overview, value.ReleaseDate, title, value.Popularity))
		if err != nil {
			fmt.Println(err.Error())
		}

		defer row.Close()
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4200"
	}

	fmt.Println("#######################")
	fmt.Println("Recover data")
	fmt.Println("#######################")
	recoverData()
	fmt.Printf("Listen to port %s\n", port)
	http.ListenAndServe(":"+port, chi.ChiRouter().InitRouter())
}
