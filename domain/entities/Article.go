package domain_entities

import "time"

type Article struct {
	Uuid      string    `json:"uuid"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
