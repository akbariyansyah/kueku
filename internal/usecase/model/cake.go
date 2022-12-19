package model

import (
	"kueku/internal/domain/cake"
	"time"
)

type CreateCakeCommand struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Rating      int    `json:"rating"`
}

type UpdateCakeCommand struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Rating      int    `json:"rating"`
}

func (c *CreateCakeCommand) ToCake() *cake.Cake {
	return &cake.Cake{
		Title:       c.Title,
		Description: c.Description,
		Image:       c.Image,
		Rating:      c.Rating,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (c *UpdateCakeCommand) ToCake() *cake.Cake {
	return &cake.Cake{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		Image:       c.Image,
		Rating:      c.Rating,
		UpdatedAt:   time.Now(),
	}
}
