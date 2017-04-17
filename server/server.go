package server

import (
	"net/http"

	"github.com/blazed/shorten/storage"
	"github.com/goware/cors"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

type Config struct {
	AllowedOrigins []string
	Storage        storage.Storage
}

type Server struct {
	mux     *chi.Mux
	storage storage.Storage
}

func NewServer(c Config) (*Server, error) {
	s := &Server{
		storage: c.Storage,
	}

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   c.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	r.Use(cors.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.DefaultCompress)

	r.Get("/", s.handleIndex)
	r.Post("/", s.handleCreate)
	r.Get("/:urlSlug", s.handleURLSlug)

	s.mux = r

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
