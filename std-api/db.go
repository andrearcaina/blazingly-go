package main

import "errors"

var PostcardsDB = []Postcard{
	{ID: "1", Name: "Welcome to the CN Tower", Where: "Toronto, Canada", To: "John Doe"},
	{ID: "2", Name: "Welcome to the Chichén Itzá", Where: "Yucatan, Mexico", To: "Jane Doe"},
	{ID: "3", Name: "Welcome to the Washington Monument", Where: "Washington, D.C., USA", To: "John Smith"},
}

func FindByID(id string) (Postcard, error) {
	for _, postcard := range PostcardsDB {
		if postcard.ID == id {
			return postcard, nil
		}
	}

	return Postcard{}, errors.New("postcard not found")
}

func UpdateByID(id string, updated Postcard) (Postcard, error) {
	for i, postcard := range PostcardsDB {
		if postcard.ID == id {
			if updated.Name != "" {
				PostcardsDB[i].Name = updated.Name
			} else if updated.Where != "" {
				PostcardsDB[i].Where = updated.Where
			} else if updated.To != "" {
				PostcardsDB[i].To = updated.To
			}

			return postcard, nil
		}
	}

	return Postcard{}, errors.New("postcard not found")
}

func RemoveByID(id string) error {
	for i, postcard := range PostcardsDB {
		if postcard.ID == id {
			PostcardsDB = append(PostcardsDB[:i], PostcardsDB[i+1:]...)
			return nil
		}
	}

	return errors.New("postcard not found")
}
