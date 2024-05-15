// controllers/browse.go

package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// understand cli requests
func HandleBrowseRequest(c *gin.Context) {
    // parse JSON request body
    var criteria map[string]interface{}
    if err := c.BindJSON(&criteria); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    // extract criterias
    contains := criteria["contains"].(string)
    notContains := criteria["notContains"].(string)
    date := criteria["date"].(string)
    count := criteria["count"].(int)
    trainable := criteria["trainable"].(*bool)

    // process browse criteria

    // only for testing
    fmt.Printf("[request] contains=%s, !contains=%s, date=%s, num=%d, train=",
		contains, notContains, date, count)
    fmt.Println(trainable)

    // send response
    c.JSON(http.StatusOK, gin.H{"message": "request received"})
}

