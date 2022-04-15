package movie

import (
	"encoding/json"
	"net/http"
	"strconv"
	"web-service/api/src/domain/movie/models"
	"web-service/api/src/interfaces"
	"web-service/api/src/middlewares"

	"github.com/go-chi/chi"
)

type MovieController struct {
	interfaces.IMovieRepository
}

func (controller *MovieController) GetAllMovies(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	// Recover all movies
	allMovies, err := controller.GetAllMoviesDB()
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Response OK
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(allMovies)
}

func (controller *MovieController) GetOneMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// Recover the params in url
	id_movie := chi.URLParam(req, "idmovie")
	ID, _ := strconv.Atoi(id_movie)

	// Get movie by ID
	movie, err := controller.GetMovieByIdDB(ID)
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Response OK
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(movie)
}

func (controller *MovieController) CreateMovie(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	// Recover the req.Body
	var movie models.ModelMovie
	if err := json.NewDecoder(req.Body).Decode(&movie); err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Verify if the movie exist
	isExist := controller.MovieExistDB(movie.Title)
	if isExist {
		err := middlewares.ServiceFonctionalError(middlewares.ErrMovieExist.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Save the movie in DB
	if err := controller.SaveMovieDB(movie); err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Response OK
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(movie)
}

func (controller *MovieController) UpdateMovie(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	// Recover the req.Body
	var movie models.ModelMovieRequest
	if err := json.NewDecoder(req.Body).Decode(&movie); err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Verify if the movie exist
	isExist := controller.MovieExistDB(movie.Title)
	if isExist {
		err := middlewares.ServiceFonctionalError(middlewares.ErrMovieExist.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Update the movie
	idMovieUpdate, err := controller.UpdateMovieDB(movie)
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Response OK
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(idMovieUpdate)
}

func (controller *MovieController) DeleteMovie(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	// Recover the req.body
	var id int
	if err := json.NewDecoder(req.Body).Decode(&id); err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Delete movie
	idMovie, err := controller.DeleteMovieDB(id)
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Response OK
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(idMovie)
}
