package http

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/nikunicke/reaktor/warehouse"
)

const cacheKey = string("productCache")
const cacheAge = string("max-age=300")

type productHandler struct {
	router         chi.Router
	baseURL        url.URL
	productService warehouse.ProductService
}

func newProductHandler() *productHandler {
	h := &productHandler{router: chi.NewRouter()}
	h.router.Use(middleware.SetHeader("Etag", cacheKey))
	h.router.Use(middleware.SetHeader("Cache-Control", cacheAge))
	h.router.Get("/", h.productIndex)
	h.router.Get("/{category}", h.getAllProductsInCategory)
	return h
}

func (h *productHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *productHandler) productIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request with category: [ jackets, shirts, accessories ]")
}

func (h *productHandler) getAllProductsInCategory(
	w http.ResponseWriter, r *http.Request) {
	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, cacheKey) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	category := chi.URLParam(r, "category")
	if err := warehouse.IsValidProductCategory(category); err != nil {
		http.Error(w, err.Error()+": Available categories: [ jackets, shirts, accessories ]", 404)
		return
	}
	products, err := h.productService.GetProducts(category)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	render.JSON(w, r, products)
}
