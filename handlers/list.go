package handlers

import (
	"api_battler/models"
	"net/http"
	"log"
	"encoding/json"
)


func List(w http.ResponseWriter, r *http.Request){

	people, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)

}