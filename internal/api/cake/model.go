package cake

import (
	"kueku/internal/usecase/model"
	"strconv"
)

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Rating      int    `json:"rating"`
}

type UpdateRequest struct {
	ID          string `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Rating      int    `json:"rating"`
}

func (c *CreateRequest) ToCommand() *model.CreateCakeCommand {
	return &model.CreateCakeCommand{
		Title:       c.Title,
		Description: c.Description,
		Rating:      c.Rating,
		Image:       c.Image,
	}
}

func (c *UpdateRequest) ToCommand() *model.UpdateCakeCommand {
	id, _ := strconv.Atoi(c.ID)
	return &model.UpdateCakeCommand{
		ID:          id,
		Title:       c.Title,
		Description: c.Description,
		Rating:      c.Rating,
		Image:       c.Image,
	}
}
