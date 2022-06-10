package jwtx

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"github.com/yunyandz/tiktok-demo-micro/internal/logger"
)

type Creater struct {
	expirestime int64
	issuer      string
	signkey     string
}

func NewCreater(cfg *viper.Viper) *Creater {
	return &Creater{
		expirestime: cfg.GetInt64("jwt.expirestime"),
		issuer:      cfg.GetString("jwt.issuer"),
		signkey:     cfg.GetString("jwt.signkey"),
	}
}

func (c *Creater) CreateUserClaims(uinfo UserInfo) (string, error) {
	// Create the Claims
	claims := UserClaims{
		uinfo,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * time.Duration(c.expirestime)).Unix(),
			Issuer:    c.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(c.signkey))
	if err != nil {
		logger.Suger().Errorw("CreateUserClaims SignedString failed", "err", err.Error())
		return "", err
	}
	return ss, err
}
