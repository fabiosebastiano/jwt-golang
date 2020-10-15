package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	ISSUER     string        = "fabiosebastiano"
	SECRET     string        = "secret"
	EXPIRATION time.Duration = 1
)

var service JWTService = NewJWTService(ISSUER, SECRET, EXPIRATION)

func TestGenerateToken(t *testing.T) {
	token := service.GenerateToken("admin", true)
	assert.NotNil(t, token)
}

func TestValidateToken(t *testing.T) {
	token := service.GenerateToken("admin", true)

	jwt, err := service.ValidateToken(token)
	assert.Nil(t, err)

	err = jwt.Claims.Valid()
	assert.Nil(t, err)
}
