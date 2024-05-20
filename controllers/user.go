// controllers/user.go

package controllers

import (
	"net/http"
	"strconv"

	"webback/models"

	"github.com/gin-gonic/gin"
)

// CreateUser handles the creation of a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := models.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdUser)
}

// GetAllUsers handles fetching all users
func GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID handles fetching a user by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := models.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUserByAstronomerID handles fetching a user by Astronomer ID
func GetUserByAstronomerID(c *gin.Context) {
	astronomerID, err := strconv.Atoi(c.Param("astronomerID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid astronomer ID"})
		return
	}
	user, err := models.GetUserByAstronomerID(astronomerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser handles deleting a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := models.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// UpdateUser handles updating a user's details
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := models.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}
