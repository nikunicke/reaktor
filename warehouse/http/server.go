package http

import (
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/nikunicke/reaktor/warehouse"
	"github.com/rs/cors"
)

type Server struct {
	ln net.Listener

	ProductService   warehouse.ProductService
	WarehouseService warehouse.WarehouseService
	Addr             string
	Recoverable      bool
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

func (s *Server) URL() url.URL {
	if s.ln == nil {
		return url.URL{}
	}
	return url.URL{
		Scheme: "http",
		Host:   s.ln.Addr().String(),
	}
}

func (s *Server) router() http.Handler {
	r := chi.NewRouter()
	fileServer := http.FileServer(http.Dir("./warehouse/build/"))
	r.Use(cors.Default().Handler)
	r.Route("/", func(r chi.Router) {
		r.Handle("/*", fileServer)
		r.Mount("/products", s.productHandler())
		r.Mount("/warehouse", s.warehouseHandler())
	})
	return r
}

func (s *Server) productHandler() *productHandler {
	h := newProductHandler()
	h.baseURL = s.URL()
	h.productService = s.ProductService
	return h
}

func (s *Server) warehouseHandler() *warehouseHandler {
	h := newWarehouseHandler()
	h.baseURL = s.URL()
	h.warehouseService = s.WarehouseService
	return h
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the index page")
}
