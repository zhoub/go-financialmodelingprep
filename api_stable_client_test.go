package gofmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSwagger(t *testing.T) {
	s, err := GetSwagger()
	assert.NotNil(t, s)
	assert.NoError(t, err)
}
