package cake_test

import (
	"kueku/internal/api/cake"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequest_ToCommand(t *testing.T) {
	mockRequest := cake.CreateRequest{
		Title:       "Vanilla Cake",
		Description: "Vanilla comes from a tropical Mexican orchid",
		Image:       "https://www.instagram.com/p/Cdn2-cor4mN/?utm_source=ig_embed&ig_rid=67367a3c-7d76-467d-b8e7-6c269b239e55",
		Rating:      7,
	}

	res := mockRequest.ToCommand()

	assert.Equal(t, res.Title, mockRequest.Title)
	assert.Equal(t, res.Description, mockRequest.Description)
	assert.Equal(t, res.Image, mockRequest.Image)
	assert.Equal(t, res.Rating, mockRequest.Rating)

}

func TestUpdateRequest_ToCommand(t *testing.T) {
	mockRequest := cake.UpdateRequest{
		ID:          "10",
		Title:       "Ice Cream Cake",
		Description: "Instead of cake and ice cream, you can have them in one perfect dessert",
		Image:       "https://www.instagram.com/p/Cemd5f8tFaZ/?utm_source=ig_embed&ig_rid=85bbfeb5-34b0-4818-800a-abc3e8e9ce7e",
		Rating:      7,
	}

	res := mockRequest.ToCommand()

	assert.Equal(t, res.ID, 10)
	assert.Equal(t, res.Title, mockRequest.Title)
	assert.Equal(t, res.Description, mockRequest.Description)
	assert.Equal(t, res.Image, mockRequest.Image)
	assert.Equal(t, res.Rating, mockRequest.Rating)
}
