package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunyandz/tiktok-demo-micro/cmd/http/handle"
	"github.com/yunyandz/tiktok-demo-micro/internal/jwtx"
	"go.uber.org/zap"
)

func JWTAuth(logger *zap.Logger, strict bool, parser *jwtx.Parser) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取token
		token := c.Query("token")
		if token == "" || len(token) == 0 {
			token = c.PostForm("token")
			if token == "" || len(token) == 0 {
				if strict {
					rsp := handle.Response{
						StatusCode: -1,
						StatusMsg:  "Invalid token",
					}
					logger.Sugar().Errorf("Invalid token: %v", rsp)
					c.JSON(http.StatusUnauthorized, rsp)
					c.Abort()
				} else {
					c.Next()
				}
				return
			}
		}

		// 解析token

		parsedToken, err := parser.ParseUserClaims(token)
		if err != nil {
			if strict {
				rsp := handle.Response{
					StatusCode: -1,
					StatusMsg:  "Parse token failed",
				}
				logger.Sugar().Errorf("Parse token failed: %v", rsp)
				c.JSON(http.StatusNonAuthoritativeInfo, rsp)
				c.Abort()
			} else {
				c.Next()
			}
			return
		}

		c.Set("claims", parsedToken)
		c.Next()
	}
}
