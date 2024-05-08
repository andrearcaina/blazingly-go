package main

import "errors"

var PostcardsDB = make(map[string]Postcard)

func init() {
	// some seed data
	PostcardsDB["1"] = Postcard{ID: "1", Name: "Welcome to the CN Tower", Where: "Toronto, Canada", To: "John Doe"}
	PostcardsDB["2"] = Postcard{ID: "2", Name: "Welcome to the Chichén Itzá", Where: "Yucatan, Mexico", To: "Jane Doe"}
	PostcardsDB["3"] = Postcard{ID: "3", Name: "Welcome to the Washington Monument", Where: "Washington, D.C., USA", To: "John Smith"}
}

func FindByID(id string) (Postcard, error) {
	if postcard, ok := PostcardsDB[id]; ok {
		return postcard, nil
	}

	return Postcard{}, errors.New("postcard not found")
}

func UpdateByID(id string, postcard Postcard) error {
	if _, ok := PostcardsDB[id]; ok {
		PostcardsDB[id] = postcard
		return nil
	}

	return errors.New("postcard not found")
}

func RemoveByID(id string) error {
	if _, ok := PostcardsDB[id]; ok {
		delete(PostcardsDB, id)
		return nil
	}

	return errors.New("postcard not found")
}
