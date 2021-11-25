package transaction

import "time"

type Transaction struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
