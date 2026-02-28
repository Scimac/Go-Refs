package routes

import (
	users "events-booking/models"
	"events-booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func userSignup(c *gin.Context) {
	var user users.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse user info. Please try again.", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register the user. Please try again later.", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user.Email})
}

func userLogin(c *gin.Context) {
	var user users.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse user info. Please try again.", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials", "error": err})
		return
	}

	token, err := utils.GenerateJwtToken(user.Email, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create a token.", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Logged in Successfully", "user": user.Email, "token": token})
}

func getAllUsers(c *gin.Context) {
	users, err := users.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
