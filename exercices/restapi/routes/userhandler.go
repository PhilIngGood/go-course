package routes

import (
	"net/http"

	"course.go/restapi/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.Users

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	if user.Save() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})

}
