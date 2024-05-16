// controllers/browse.go

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BrowseCriteria struct {
	Contains    string `json:"contains"`
	NotContains string `json:"notContains"`
	Date        string `json:"date"`
	Count       int    `json:"count"`
	Trainable   int    `json:"trainable"`
}

// HandleBrowseRequest understand cli requests
func HandleBrowseRequest(c *gin.Context) {
	// parse JSON request body
	var criteria = BrowseCriteria{}

	if err := c.BindJSON(&criteria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// process browse criteria

	// only for testing
	fmt.Printf("[request] contains=%v, !contains=%v, date=%v, num=%v, train=%v",
		criteria.Contains, criteria.NotContains, criteria.Date, criteria.Count, criteria.Trainable)

	// send response
	c.JSON(http.StatusOK, gin.H{"message": "request received"})
}
