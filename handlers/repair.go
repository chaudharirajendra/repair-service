package handlers

import (
	"log"
	"net/http"
	"repair-service/services"

	"github.com/gin-gonic/gin"
)

// RepairHandler handles GET requests for /repairs
func RepairHandler(c *gin.Context) {
	machineID := c.Query("machine_id")

	// Fetch filtered repairs from the service layer
	repairs, err := services.GetRepairsByMachineID(machineID)
	if err != nil {
		log.Printf("Error fetching repairs: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading repair data"})
		return
	}

	// Return the filtered repairs as JSON response
	c.JSON(http.StatusOK, repairs)
}
