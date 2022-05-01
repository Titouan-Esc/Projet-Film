package movie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"web-service/api/src/domain/movie/models"
	"web-service/api/src/interfaces"
	"web-service/api/src/middlewares"
	utils "web-service/api/src/utils"
)

type MovieController struct {
	interfaces.IMovieRepository
}

// ? Fonction pour la récupération des films puis les insérer dans la Base De Donnée
func (controller *MovieController) GetFilmInMovieDB(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Créer un struct pour le req.body
	var Body struct {
		Page int `json:"page"`
	}

	// Décoder le req.Body dans le pointer du struct Body
	if err := json.NewDecoder(req.Body).Decode(&Body); err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Appel de ma fonction pour récupérer les films sur l'api
	response, err := utils.ConsumApi(fmt.Sprintf(`http://api.themoviedb.org/3/discover/movie?page=%d&`, Body.Page))
	if err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Unmarshal les données récupérer de l'api dans un json formaté
	var genre models.DiscoverModel
	if err := json.Unmarshal([]byte(response), &genre); err != nil {
		err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(err)
		return
	}

	// Créer un array du model d'un film
	allMovies := make([]*models.ResultsModel, 0)

	// Créer un boucle pour récupérer les films un à un
	for _, value := range genre.Results {

		// Parse overview et title car pour l'insertion en db les 's pose un gros problème et sur postgreSQL pour remédier à ce problème il faut doubler '
		overview := strings.Replace(value.Overview, "'", "''", -1)
		title := strings.Replace(value.Title, "'", "''", -1)

		// Vérifier si les films exist ou non, pour éviter d'ajout plusieurs fois le même film
		isExist, _ := controller.MovieExist(value.ID)
		if isExist {
			err := middlewares.ServiceFonctionalError(middlewares.ErrMovieExist.Error(), http.StatusInternalServerError)
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err)
			return
		}

		// Insérer les films dans la base de donnée
		movies, _ := controller.SaveMovieDB(value, overview, title)
		if err != nil {
			err := middlewares.ServiceFonctionalError(err.Error(), http.StatusInternalServerError)
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(err)
			return
		}

		// Ajout à l'array précédement créé
		allMovies = append(allMovies, &movies)
	}

	// Renvoyer un response 200 puis l'array de nos films inséré en db
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
}
