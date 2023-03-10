package controller

import (
	"jwt-project/dto"
	"jwt-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(c *gin.Context) {
	var dtoPerson dto.DtoSignUp
	c.BindJSON(&dtoPerson)

	insert, err := service.InsertInDatabase(c, dtoPerson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, insert)
}

func Login(c *gin.Context) {
	var dtoPerson dto.DtoLogIn
	c.BindJSON(&dtoPerson)

	foundPerson, err := service.FindInDatabase(c, dtoPerson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &foundPerson.ID)
}

func GetUser(c *gin.Context) {
	var dtoPerson dto.GetUser

	personId := c.Param("userid")

	person, err := service.GetFromDatabase(c, dtoPerson, personId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

func GetUsers(c *gin.Context) {
	var allUsers []primitive.M

	allUsersList, err := service.GetallFromDatabase(c, allUsers)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, allUsersList)
}
