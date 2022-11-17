package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jeheskielSunloy77/go-ecommerce-scraper/scrapers"
)

func Json(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query().Get("q")

	splitedQuery := strings.Split(query, " ")
	searchQuery := strings.Join(splitedQuery, "+")
	products := scrapers.Scraper(searchQuery)

	json.NewEncoder(w).Encode(products)
}
