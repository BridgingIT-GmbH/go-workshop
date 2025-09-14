package model

type Movie struct {
	Id                string   `json:"id"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	Actors            []string `json:"actors"`
	DurationInMinutes int      `json:"durationInMinutes"`
}
