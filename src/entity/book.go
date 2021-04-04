package entity

// Book is an entity with title, writer and cost as its identifiers
type Book struct {
	Title  string  `json:"Title"`
	Writer string  `json:"Writer"`
	Cost   float64 `json:"Cost"`
}
