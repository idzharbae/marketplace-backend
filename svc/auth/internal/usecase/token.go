package usecase

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/idzharbae/marketplace-backend/svc/auth/constant"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/config"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/entity"
	"github.com/idzharbae/marketplace-backend/svc/auth/internal/request"
	"time"
)

type Token struct {
	secret []byte
}

func NewToken(cfg config.Config) *Token {
	return &Token{secret: []byte(cfg.Secret)}
}

func (t *Token) Get(req entity.User) (entity.AuthToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        req.ID,
		"role":      req.Type,
		"full_name": req.Name,
		"user_name": req.UserName,
		"email":     req.Email,
		"phone":     req.Phone,
		"expr":      time.Now().Unix() + constant.TokenExpiringEpoch,
	})
	signedToken, err := token.SignedString(t.secret)
	if err != nil {
		return entity.AuthToken{}, err
	}
	return entity.AuthToken{
		Token: signedToken,
	}, nil
}

func (t *Token) Refresh(req request.RefreshToken) (entity.AuthToken, error) {
	token, err := jwt.Parse(req.CurrentToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return t.secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		token, err := t.Get(entity.User{
			ID:       int64(claims["id"].(float64)),
			Name:     claims["full_name"].(string),
			UserName: claims["user_name"].(string),
			Email:    claims["email"].(string),
			Phone:    claims["phone"].(string),
			Type:     int32(claims["role"].(float64)),
		})
		if err != nil {
			return entity.AuthToken{}, err
		}
		return token, nil
	}
	return entity.AuthToken{}, err
}

func (t *Token) Validate(req entity.AuthToken) error {
	_, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return t.secret, nil
	})
	return err
}
