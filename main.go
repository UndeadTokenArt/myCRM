package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/undeadtokenart/myCRM/customerDB"
)

func main() {

	// Initialize Gin
	router := gin.Default()
	// Serve static files from the "static" directory
	router.Static("/static", "./static")
	// define the location of tmpl files
	router.LoadHTMLGlob("templates/*")

	// Initialize the database
	var err error
	customerDB.MyDataBase, err = customerDB.GetDB("sqlite3", "crm.db")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer customerDB.MyDataBase.Close()

	var customerClient customerDB.Customer

	// Auto-migrate the database to create the Customer table
	customerDB.MyDataBase.AutoMigrate(customerClient)

	// Define your routes and handlers here
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.tmpl", gin.H{
			"Message": "Welcome to the CRM Page",
		})
	})

	// For all intensive purposes, I was using this to test the gin.H passing of data
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.tmpl", gin.H{
			"Message": "Welcome",
		})
	})

	// Saving these for once the form is ready
	router.POST("/customers", customerDB.CreateCustomer)
	router.GET("/customers/:id", customerDB.GetCustomer)

	router.GET("/newCustomerForm", func(c *gin.Context) {
		c.HTML(http.StatusOK, "NewCustomerForm.tmpl", gin.H{
			"Message": "Please tell me about the new Client",
		})
	})
	// Start the Gin server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	router.Run(port)
}
