package handlers

import (
	"api_battler/models"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var people models.People

	err := json.NewDecoder(r.Body).Decode(&people)
	
	if err != nil {
		log.Printf("Erro ao fazer decode do json %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	id, err := models.Insert(people)
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserido com sucesso! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}