package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type PostcardHandler struct{}

func (ph *PostcardHandler) ListPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(FindAll())
}

func (ph *PostcardHandler) ListPostByID(w http.ResponseWriter, r *http.Request) {
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

func (ph *PostcardHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	newCard, err := decodeBodyRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	_, err = FindByID(strconv.Itoa(newCard.ID))

	if err == nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Postcard with ID %s already exists\n", strconv.Itoa(newCard.ID))
		return
	}

	PostcardsDB[strconv.Itoa(newCard.ID)] = newCard

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCard)
}

func (ph *PostcardHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	updatedCard, err := decodeBodyRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	err = UpdateByID(id, updatedCard)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Postcard with ID %s not found\n", id)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCard)
}

func (ph *PostcardHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
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
