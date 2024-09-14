package routes

import (
	"net/http"

	"course.go/restapi/models"
	"course.go/restapi/utils"
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

func login(c *gin.Context) {
	var user models.Users

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse user request data"})
		return
	}

	err = user.ValidatedCredentials()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(user.Email, user.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}
