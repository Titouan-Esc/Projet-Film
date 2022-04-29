package movie

import (
	"encoding/json"
	"net/http"
	"web-service/api/src/domain/movie/models"
	"web-service/api/src/interfaces"
	"web-service/api/src/middlewares"
	utils "web-service/api/src/utils"

	"github.com/joho/godotenv"
)

type MovieController struct {
	interfaces.IMovieRepository
}

func (controller *MovieController) GetFilmInMovieDB(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	godotenv.Load()
	response, err := utils.ConsumApi("http://api.themoviedb.org/3/discover/movie?page=1&")
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	var genre models.DiscoverModel
	if err := json.Unmarshal([]byte(response), &genre); err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	for _, value := range genre.Results {
		isExist, _ := controller.MovieExist(value.ID)
		if isExist {
			err := middlewares.ServiceFonctionalError(middlewares.ErrMovieExist.Error(), http.StatusInternalServerError)
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err)
			return
		}

		idMovie, err := controller.SaveMovieDB(value)
		if err != nil {
			err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err)
			return
		}

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(idMovie)
	}

	res.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(res).Encode(genre)
}
