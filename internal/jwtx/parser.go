package jwtx

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"github.com/yunyandz/tiktok-demo-micro/internal/errorx"
	"github.com/yunyandz/tiktok-demo-micro/internal/logger"
)

type Parser struct {
	expirestime int64
	issuer      string
	signkey     string
}

func NewParser(cfg *viper.Viper) *Parser {
	return &Parser{
		expirestime: cfg.GetInt64("jwt.expirestime"),
		issuer:      cfg.GetString("jwt.issuer"),
		signkey:     cfg.GetString("jwt.signkey"),
	}
}

// ParseUserClaims parse jwt token and return user claims
func (p *Parser) ParseUserClaims(tokenString string) (*UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Suger().Error("ParseWithClaims but without SigningMethodHMAC method")
			return nil, errorx.ErrTokenMethod
		}
		return []byte(p.signkey), nil
	})

	if err != nil {
		//handle error in more detail
		// TODO:
		// It may be necessary to customize the Valid function which throw a custom error if an error occurs.
		// Allow for more detailed error handling.

		//if verr , ok := err.(*jwt.ValidationError); ok && errors.Is(verr.Inner, errorx.ExpiredTokenErr);

		return nil, err
		//return nil, errors.New(errorx.ParseJWTFailed)
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errorx.ErrInvalidToken
}
