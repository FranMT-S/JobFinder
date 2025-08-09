package server

import (
	"fmt"
	"net/http"

	"github.com/FranMT-S/JobFinder/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type Server struct {
	port   string
	Router *chi.Mux
}

func NewServer(port string) *Server {
	router := chi.NewRouter()
	setupMiddleware(router)
	setupRoutes(router)
	return &Server{
		port:   port,
		Router: router,
	}
}

func setupMiddleware(router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.CleanPath)
	router.Use(middleware.StripSlashes)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{config.CLIENT_HOST},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Custom-Header"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// compression
	router.Use(middleware.Compress(5, "application/json", "application/xml", "text/plain"))
}

func setupRoutes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Api is running!"))
	})

	router.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("404 Not Found"))
	})

	// api routers
	router.Route("/api/"+config.API_VERSION, func(r chi.Router) {
		addScraperRoutes(r)
		addCategoriesRoutes(r)
		addSkillsRoutes(r)
	})
}

func (s *Server) Start() error {
	fmt.Println("Starting server on port", s.port)
	return http.ListenAndServe(":"+s.port, s.Router)
}
