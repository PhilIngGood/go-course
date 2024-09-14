package middlewares

import (
	"net/http"

	"course.go/restapi/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		// it abort the request/response and no other event handler will be managed after
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not allowed to perform this action"})
		return
	}

	userid, err := utils.ValidateJWT(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not allowed to perform this action"})
		return
	}

	// Set add data to the context
	c.Set("userId", userid)

	// Execute next request handler
	c.Next()
}
