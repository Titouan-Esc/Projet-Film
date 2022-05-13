package movie

import (
	"encoding/json"
	"net/http"
	"strings"
	"web-service/api/src/domain/movie/models"
	"web-service/api/src/interfaces"
	"web-service/api/src/middlewares"
)

type MovieController struct {
	interfaces.IMovieRepository
}

func (controller *MovieController) RecoverMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")

	allMovies, err := controller.RecoverMoviesDB()
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(allMovies)
}

// ? Fonction pour le système de Like
func (controller *MovieController) Like(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var Body struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(req.Body).Decode(&Body); err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	allLikes, err := controller.RecoverLikesMovieDB(Body.ID, "likes")
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	allLikes += 1

	likes, err := controller.LikeMovieDB(Body.ID, allLikes, "likes")
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(likes)
}

// ? Fonction pour le système de Dislikes
func (controller *MovieController) Dislike(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var Body struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(req.Body).Decode(&Body); err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	allDislikes, err := controller.RecoverLikesMovieDB(Body.ID, "dislikes")
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	allDislikes += 1

	likes, err := controller.LikeMovieDB(Body.ID, allDislikes, "dislikes")
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(likes)
}

// ? Fonction pour ajouter un commentaire et l'update
func (controller *MovieController) AddComment(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var comments models.CommentRequest
	if err := json.NewDecoder(req.Body).Decode(&comments); err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	content := strings.Replace(comments.Comments, "'", "''", -1)

	comment, err := controller.AddCommentMovieDB(comments.ID, content)
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(comment)
}

// ? Fonction pour supprimer un commentaire
func (controller *MovieController) DeleteComment(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var Body struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(req.Body).Decode(&Body); err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	idMovie, err := controller.DeleteCommentMovieDB(Body.ID)
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(idMovie)
}
