package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webback/requests"
	"webback/responses"
	"webback/services"
)

type User struct {
	authService services.Auth
	userService services.User
}

func NewUserController(authService services.Auth) *User {
	return &User{
		authService: authService,
	}
}

func (controller *User) Login(c *gin.Context) {
	var request requests.Login
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := controller.userService.Login(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	stringID := strconv.FormatUint(uint64(user.ID), 10)

	token, err := controller.authService.GenerateToken(stringID, user.Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &responses.Login{Token: token})
}

func (controller *User) Register(c *gin.Context) {
	var request requests.Register
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.userService.Register(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered"})
}
