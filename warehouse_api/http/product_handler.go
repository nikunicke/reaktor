package http

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/nikunicke/reaktor/warehouse_api"
)

type productHandler struct {
	router         chi.Router
	baseURL        url.URL
	productService warehouse_api.ProductService
}

func newProductHandler() *productHandler {
	h := &productHandler{router: chi.NewRouter()}
	h.router.Get("/", h.productIndex)
	h.router.Get("/{category}", h.getAllProductsInCategory)
	return h
}

func (h *productHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *productHandler) productIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request with category")
}

func (h *productHandler) getAllProductsInCategory(
	w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	t := setCategory(category)
	products, err := h.productService.GetProducts(t)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	render.JSON(w, r, products)
}

func setCategory(ctg string) string {
	return "products/" + ctg
}
