package service

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPasswordEncrypted(t *testing.T) {
	password := []byte("123456#")

	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	assert.NoError(t, err)
	err = bcrypt.CompareHashAndPassword(encrypted, password)
	assert.NoError(t, err)
}
