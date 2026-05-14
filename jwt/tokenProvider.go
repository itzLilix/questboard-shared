package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/itzLilix/questboard-shared/dtos"
)

type Parser interface {
	ParseToken(tokenString string) (*dtos.TokenClaims, error)
}

type Provider interface {
	Parser
	GenerateAccessToken(userID string, role dtos.Role) (string, error)
}

type tokenProvider struct {
	secretKey      []byte
	accessTokenTTL time.Duration
}

type claims struct {
	UserID string    `json:"user_id"`
	Role   dtos.Role `json:"role"`
	jwt.RegisteredClaims
}

func NewParser(secretKey []byte) Parser {
	return &tokenProvider{secretKey: secretKey}
}

func NewProvider(secretKey []byte, accessTokenTTL time.Duration) Provider {
	return &tokenProvider{
		secretKey:      secretKey,
		accessTokenTTL: accessTokenTTL,
	}
}

func (tp *tokenProvider) GenerateAccessToken(userID string, role dtos.Role) (string, error) {
	expirationTime := time.Now().Add(tp.accessTokenTTL)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	})

	return token.SignedString(tp.secretKey)
}

func (tp *tokenProvider) ParseToken(tokenString string) (*dtos.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (any, error) {
		return tp.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if c, ok := token.Claims.(*claims); ok && token.Valid {
		return &dtos.TokenClaims{
			UserID: c.UserID,
			Role:   c.Role,
		}, nil
	}

	return nil, fmt.Errorf("invalid claims")
}
