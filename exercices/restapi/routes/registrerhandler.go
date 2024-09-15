package routes

import (
	"net/http"
	"strconv"

	"course.go/restapi/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	var register models.Registration

	register.UserId = c.GetInt64("userid")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to fetch event"})
		return
	}

	register.EventId = eventId

	//err = c.ShouldBindJSON(&register)
	//
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
	//	return
	//}

	if register.Save() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not save registration, please try again later"})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registred successfully", "object": register})
}

func cancelRegistration(c *gin.Context) {
	regId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to pars request dataz"})
	}

	reg, err := models.GetRegistrationById(regId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to find this registration"})
	}

	reg.DeleteRegistration()

}

func getRegistration(c *gin.Context) {
	registerId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse ID to int"})
		return
	}

	registration, err := models.GetRegistrationById(registerId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch registration by id"})
		return
	}

	c.JSON(http.StatusOK, registration)
}
