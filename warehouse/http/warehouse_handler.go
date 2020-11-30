package http

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/nikunicke/reaktor/warehouse"
)

type warehouseHandler struct {
	router           chi.Router
	baseURL          url.URL
	warehouseService warehouse.WarehouseService
}

func newWarehouseHandler() *warehouseHandler {
	h := &warehouseHandler{router: chi.NewRouter()}
	h.router.Get("/hard-refresh", h.hardRefresh)
	return h
}

func (h *warehouseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *warehouseHandler) hardRefresh(w http.ResponseWriter, r *http.Request) {
	if err := h.warehouseService.Update(); err != nil {
		http.Error(w, "Error: "+err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Data has been updated. Please refresh the page")
}
