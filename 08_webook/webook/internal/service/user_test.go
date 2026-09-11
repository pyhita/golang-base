package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestEncryptPass(t *testing.T) {

	input := "123"
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(input))

	assert.NoError(t, err, hash)
}
