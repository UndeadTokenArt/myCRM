package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Define a model struct for your CRM data (e.g., Customer)
type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	// Add more fields as needed
}

var db *gorm.DB

func main() {
	// Initialize Gin
	router := gin.Default()
	// Serve static files from the "static" directory
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	// Initialize the database (SQLite in this example)
	var err error
	db, err = gorm.Open("sqlite3", "crm.db")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	// Auto-migrate the database to create the Customer table
	db.AutoMigrate(&Customer{})

	// Define your routes and handlers here
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.html", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// Example route for creating a customer
	router.POST("/customers", createCustomer)

	// Example route for fetching a customer by ID
	router.GET("/customers/:id", getCustomer)

	// Add more routes for updating, deleting, and listing customers

	// Start the Gin server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	router.Run(port)
}

func createCustomer(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&customer)
	c.JSON(http.StatusOK, customer)
}

func getCustomer(c *gin.Context) {
	var customer Customer
	id := c.Param("id")

	if err := db.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}
