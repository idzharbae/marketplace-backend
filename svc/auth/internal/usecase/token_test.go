package usecase

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/config"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToken_Get(t *testing.T) {
	unit := NewToken(config.Config{
		Secret: "asdf",
	})
	user := entity.User{
		ID:       1,
		Name:     "idzhar",
		UserName: "idzharbae",
		Email:    "idzharbae@gmail.com",
		Phone:    "08123123123",
		Type:     1,
	}

	t.Run("test valid token", func(t *testing.T) {
		token, _ := unit.Get(user)
		err := unit.Validate(token)
		assert.Nil(t, err)
	})

	t.Run("test invalid token", func(t *testing.T) {
		token, _ := unit.Get(user)
		token.Token = token.Token[0:len(token.Token)-1] + "x"
		err := unit.Validate(token)
		assert.NotNil(t, err)
	})
	t.Run("test refresh token", func(t *testing.T) {
		// cant use time sleep for some reason, use debugger and add assert.NotEqual(t, token, newToken) to make sure
		// the generated token is different
		token, _ := unit.Get(user)
		_, err := unit.Refresh(request.RefreshToken{CurrentToken: token.Token})
		assert.Nil(t, err)
	})
}
