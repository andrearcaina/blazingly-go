package main

type Postcard struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Where string `json:"where"`
	To    string `json:"to"`
}
