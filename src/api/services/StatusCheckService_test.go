package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusCheckService_GetStatusHealth_Success(t *testing.T) {
	//Arrange
	expected := "Pong"

	//Act
	actual, err := StatusCheckService{}.GetStatusHealth()

	//Assert
	assert.Equal(t, actual, expected)
	assert.Nil(t, err)
}
