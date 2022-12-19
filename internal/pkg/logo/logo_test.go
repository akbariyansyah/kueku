package logo_test

import (
	"kueku/internal/pkg/logo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogo(t *testing.T) {
	logo := logo.GetLogo()

	assert.IsType(t, "", logo)
}
