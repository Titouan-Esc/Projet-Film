package router

import (
	"sync"
	"web-service/api/container"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	movieController := container.ServiceContainer().InjectMovieController()

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)

	// -----------------------------
	// Routes for movies
	// -----------------------------

	r.Post("/movies", movieController.GetFilmInMovieDB)
	r.Post("/movie/like", movieController.Like)
	r.Post("/movie/dislike", movieController.Dislike)
	r.Post("/movie/comment", movieController.AddComment)
	r.Put("/movie/comment", movieController.UpdateComment)
	r.Delete("/movie/comment", movieController.DeleteComment)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
