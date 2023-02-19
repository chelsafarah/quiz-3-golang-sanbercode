package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth() gin.Accounts {
	return gin.Accounts{
		"admin":  "password",
		"editor": "secret",
	}
}
