package main

type Postcard struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Where string `json:"where"`
	To    string `json:"to"`
}
