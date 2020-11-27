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

type availabilityHandler struct {
	router              chi.Router
	baseURL             url.URL
	availabilityService warehouse.AvailabilityService
}

func newAvailabilityHandler() *availabilityHandler {
	h := &availabilityHandler{router: chi.NewRouter()}
	h.router.Use(middleware.SetHeader("Etag", cacheKey))
	h.router.Use(middleware.SetHeader("Cache-Control", cacheAge))
	h.router.Get("/", h.availabilityIndex)
	h.router.Get("/{manufacturer}", h.getAllAvailabilityInManufacturer)
	return h
}

func (h *availabilityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *availabilityHandler) availabilityIndex(
	w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request with manufacturer")
}

func (h *availabilityHandler) getAllAvailabilityInManufacturer(
	w http.ResponseWriter, r *http.Request) {
	if match := r.Header.Get("If-None-Match"); match != "" {
		fmt.Println("match")
		if strings.Contains(match, cacheKey) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	manufacturer := chi.URLParam(r, "manufacturer")
	availabilities, err := h.availabilityService.GetAvailability(setManufacturer(manufacturer))
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	render.JSON(w, r, availabilities)
}

func setManufacturer(manufacturer string) string {
	return "availability/" + manufacturer
}
