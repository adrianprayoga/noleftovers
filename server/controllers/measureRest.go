package controllers

import (
	"encoding/json"
	"github.com/adrianprayoga/noleftovers/server/models"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type MeasureResource struct {
	Service *models.MeasureService
}

func (rs MeasureResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List) // GET /posts - Read a list of measure.

	return r
}

func (rs MeasureResource) List(w http.ResponseWriter, r *http.Request) {
	measures, err := rs.Service.GetMeasures()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res, _ := json.Marshal(measures)
	_, err = w.Write(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
