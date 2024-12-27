package main

import (
	"fmt"
	"log"
	"repair-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define the route for repair handler
	router.GET("/repairs", handlers.RepairHandler)

	port := 8082
	log.Printf("Repair Service running at http://localhost:%d\n", port)

	// Start the server
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
