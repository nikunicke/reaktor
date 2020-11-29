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

type warehouseHandler struct {
	router           chi.Router
	baseURL          url.URL
	warehouseService *warehouse.Warehouse
}

func newWarehouseHandler() *warehouseHandler {
	h := &warehouseHandler{router: chi.NewRouter()}
	h.router.Use(middleware.SetHeader("Etag", cacheKey))
	h.router.Use(middleware.SetHeader("Cache-Control", cacheAge))
	h.router.Get("/{category}", h.getWarehouse)
	h.router.Get("/hard-refresh", h.hardRefresh)
	return h
}

func (h *warehouseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *warehouseHandler) getWarehouse(w http.ResponseWriter, r *http.Request) {
	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, cacheKey) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	category := chi.URLParam(r, "category")
	if err := warehouse.IsValidProductCategory(category); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	products, err := h.warehouseService.GetProducts(category)
	if err != nil {
		http.Error(w, err.Error(), 422)
	}
	render.JSON(w, r, products)
}

func (h *warehouseHandler) hardRefresh(w http.ResponseWriter, r *http.Request) {
	if err := h.warehouseService.Update(); err != nil {
		http.Error(w, "Error: "+err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Data has been updated. Please refresh the page")
}
