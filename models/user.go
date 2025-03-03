package models

import "time"

type (
	User struct {
		UserID    int       `json:"userId"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)
