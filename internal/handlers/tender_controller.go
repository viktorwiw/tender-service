package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tender-service/internal/services"
)

func CreateTenderHandler(c *gin.Context) {
	var input services.CreateTenderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tender, err := services.CreateTender(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tender)
}

func PublishTenderHandler(c *gin.Context) {
	tenderID := c.Param("tenderId")
	username := c.Request.Header.Get("Username")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	tender, err := services.PublishTender(tenderID, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tender)
}
