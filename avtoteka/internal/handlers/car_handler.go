package handlers

import (
	"avtoteka/avtoteka/internal/models"
	"avtoteka/avtoteka/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//обработчики запросов

type CarHandler struct {
	service *services.CarService //инъекция сервиса
}

func NewCarHandler(service *services.CarService) *CarHandler {
	return &CarHandler{service: service}
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cars models.Car
	if err := json.NewDecoder(r.Body).Decode(&cars); err != nil {
		http.Error(w, `{"error": "Invalid json"}`, http.StatusBadRequest)
		return
	}

	if cars.Brand == "" || cars.Model == "" {
		http.Error(w, `{"error": "Brand and Model are required"`, http.StatusBadRequest)
	}

	if err := h.service.CreateCar(r.Context(), &cars); err != nil {
		http.Error(w, `{"error": "Failed to create car"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cars)
}

func (h *CarHandler) GetCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error: "Invalid Car ID"}`, http.StatusBadRequest)
		return
	}

	car, err := h.service.GetCar(r.Context(), id)
	if err != nil {
		http.Error(w, `"{error: "DB error"}"`, http.StatusInternalServerError)
		return
	}
	if car == nil {
		http.Error(w, `{"error": "Car not found"}`, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}
