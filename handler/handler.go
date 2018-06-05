package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"search/storage"
	"strconv"
)

type handler struct {
	prefix  string
	storage storage.Service
}

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"response"`
}

var inputFields = map[string]bool{
	"searchTerm": true,
	"lat":        true,
	"lng":        true,
}

func New(prefix string, st storage.Service) http.Handler {
	mux := http.NewServeMux()
	h := handler{prefix, st}

	mux.HandleFunc("/search", responseHandler(h.search))

	return mux
}

func responseHandler(h func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r)
		if err != nil {
			data = err.Error()
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		err = json.NewEncoder(w).Encode(response{Data: data, Success: err == nil})
		if err != nil {
			log.Printf("could not encode response to output: %v", err)
		}
	}
}

func (h handler) search(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	if r.Method != http.MethodPost {
		return nil, http.StatusMethodNotAllowed, fmt.Errorf("method %s not allowed", r.Method)
	}

	var searchTerm string
	var lat, lng float64

	for key, val := range r.URL.Query() {

		if !inputFields[key] {
			return nil, http.StatusBadRequest, fmt.Errorf("invalid search term: %v", key)
		}

		switch key {
		case "searchTerm":
			searchTerm = val[0]
		case "lat", "lng":
			coor, err := strconv.ParseFloat(val[0], 64)
			if err != nil {
				return nil, http.StatusBadRequest, fmt.Errorf("invalid %v value: %v", key, val[0])
			}
			if key == "lat" {
				lat = coor
			} else {
				lng = coor
			}
		}
	}

	res, err := h.storage.Search(searchTerm, lat, lng)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("could not fetch result: %v", err)
	}

	return res, http.StatusOK, nil
}
