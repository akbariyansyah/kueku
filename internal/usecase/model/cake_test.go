package model_test

import (
	"kueku/internal/usecase/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCommand_ToCake(t *testing.T) {
	mockCommand := &model.CreateCakeCommand{
		Title:       "Cheese Cake",
		Description: "Light sweet cake with cheese on top of it",
		Image:       "http://google.com/cheese.jpg",
		Rating:      5,
	}
	res := mockCommand.ToCake()

	assert.Equal(t, res.Title, mockCommand.Title)
	assert.Equal(t, res.Description, mockCommand.Description)
	assert.Equal(t, res.Image, mockCommand.Image)
	assert.Equal(t, res.Rating, mockCommand.Rating)
}

func TestUpdateCommand_ToCake(t *testing.T) {
	mockCommand := &model.UpdateCakeCommand{
		ID:          3,
		Title:       "Red Velvet Cake",
		Description: "Cake that branched off from the chocolate devil’s food cake using cocoa powder instead of melted chocolate bars.",
		Image:       "http://google.com/red-velvet.jpg",
		Rating:      3,
	}
	res := mockCommand.ToCake()

	assert.Equal(t, res.Title, mockCommand.Title)
	assert.Equal(t, res.ID, mockCommand.ID)
	assert.Equal(t, res.Description, mockCommand.Description)
	assert.Equal(t, res.Image, mockCommand.Image)
	assert.Equal(t, res.Rating, mockCommand.Rating)
}
