package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostcardHandler struct{}

func (ph *PostcardHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PostcardsDB)
}

func (ph *PostcardHandler) GetID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	postcard, err := FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Postcard with ID %s not found\n", id)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postcard)
}

func (ph *PostcardHandler) Create(w http.ResponseWriter, r *http.Request) {
	newCard := decodeBodyRequest(w, r)
	_, err := FindByID(newCard.ID)

	if err == nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Postcard with ID %s already exists\n", newCard.ID)
		return
	}

	PostcardsDB = append(PostcardsDB, newCard)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCard)
}

func (ph *PostcardHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	updatedCard := decodeBodyRequest(w, r)
	postcard, err := UpdateByID(id, updatedCard)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Postcard with ID %s not found\n", id)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postcard)
}

func (ph *PostcardHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	_, err := FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Postcard with ID %s not found\n", id)
		return
	}

	err = RemoveByID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error removing postcard with ID %s\n", id)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Postcard " + id + " removed\n"))
}

func decodeBodyRequest(w http.ResponseWriter, r *http.Request) Postcard {
	var updatedCard Postcard

	err := json.NewDecoder(r.Body).Decode(&updatedCard)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding postcard: %v\n", err)
	}

	return updatedCard
}
