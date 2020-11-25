package http

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const baseURL = "https://bad-api-assignment.reaktor.com"

type Server struct {
	ln net.Listener
	// Add services

	Addr        string
	Recoverable bool
}

func NewServer() *Server {
	return &Server{Recoverable: true}
}

func (s *Server) Open() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln
	go http.Serve(s.ln, s.router())
	return nil
}

func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}

func (s *Server) router() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Default().Handler)
	r.Route("/", func(r chi.Router) {
		r.Get("/", handleIndex)
	})
	return r
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the index page")
}
