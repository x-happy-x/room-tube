package model

type Movie struct {
	ID     int64  `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Genre  string `json:"genre" db:"genre"`
	Year   int    `json:"year" db:"year"`
	Length int    `json:"length" db:"length"`
}
