package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/undeadtokenart/myCRM/customerDB"
)

var customerClient customerDB.Customer

var db *gorm.DB

func main() {
	// Initialize Gin
	router := gin.Default()
	// Serve static files from the "static" directory
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	// Initialize the database
	var err error
	db, err = customerDB.GetDB("sqlite3", "crm.db")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	var customerClient customerDB.Customer

	// Auto-migrate the database to create the Customer table
	db.AutoMigrate(customerClient)

	// Start the Gin server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	router.Run(port)

	// Define your routes and handlers here
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.tmpl", gin.H{
			"Message": "Welcome to the CRM Page",
		})
	})

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.POST("/customers", customerDB.CreateCustomer)
	router.GET("/customers/:id", customerDB.GetCustomer)
}
