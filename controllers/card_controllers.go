// controllers/card_controller.go
package controllers

import (
	"encoding/json"
	"net/http"
	"public-vault-ms/services"
)

// TokenizeCardHandler maneja la generación de un token desde un número de tarjeta
func TokenizeCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		CardNumber string `json:"card_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	token, err := services.TokenizeCard(request.CardNumber)
	if err != nil {
		http.Error(w, "Error al procesar la tarjeta", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DetokenizeCardHandler maneja la recuperación del número de tarjeta desde un token
func DetokenizeCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	cardNumber, err := services.DetokenizeCard(request.Token)
	if err != nil {
		http.Error(w, "Error al recuperar la tarjeta", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"card_number": cardNumber}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
