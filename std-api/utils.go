package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func decodeBodyRequest(r *http.Request) (Postcard, error) {
	var updatedCard Postcard

	if r.Method == http.MethodPut {
		updatedCard.ID, _ = strconv.Atoi(r.PathValue("id"))
	}

	err := json.NewDecoder(r.Body).Decode(&updatedCard)

	if err != nil || updatedCard.ID == 0 {
		return Postcard{}, errors.New("error decoding postcard, or ID is empty")
	}

	return updatedCard, nil
}
