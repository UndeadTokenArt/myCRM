package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

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

	// Drop the existing tables (if they exist) DELETE FOR PRODUCTION
	customerDB.MyDataBase.DropTableIfExists(&customerDB.Customer{})

	// Auto-migrate the database to create the Customer table
	customerDB.MyDataBase.AutoMigrate(&customerDB.Customer{})

	// Define the routes and handlers
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.tmpl", gin.H{
			"Message": "Welcome to the CRM Page",
		})
	})

	// routes for displaying customer data. probably - Not fully implemented
	router.GET("/customers/:id", customerDB.GetCustomer)

	// render the newCustomerForm.tmpl - Note fields is passed but not used.
	router.GET("/newCustomerForm", func(c *gin.Context) {
		clientKeys := customerDB.Customer{}
		fields := getFieldNames(clientKeys)

		c.HTML(http.StatusOK, "NewCustomerForm.tmpl", fields)
	})

	router.POST("/newCustomerForm", func(c *gin.Context) {
		// Parse the form data into a Customer struct
		var clientKeys customerDB.Customer

		if err := c.ShouldBind(&clientKeys); err != nil {
			c.HTML(http.StatusBadRequest, "base.tmpl", gin.H{
				"Message": "Error: Unable to process the form data.",
			})
			fmt.Println("Buyer:", clientKeys.Buyer)
			fmt.Println("Seller:", clientKeys.Seller)
			return
		}

		// Create a new customer record in the database
		if err := customerDB.MyDataBase.Create(&clientKeys).Error; err != nil {
			fmt.Println("Error creating customer record:", err)
			c.HTML(http.StatusInternalServerError, "base.tmpl", gin.H{
				"Message": "Error: Unable to create a new customer record.",
			})
			fmt.Println("Buyer:", clientKeys.Buyer)
			fmt.Println("Seller:", clientKeys.Seller)
			return
		}

		c.HTML(http.StatusOK, "base.tmpl", gin.H{
			"Message": "Customer record created successfully!",
		})
		fmt.Println("Buyer:", clientKeys.Buyer)
		fmt.Println("Seller:", clientKeys.Seller)
	})

	// Start the Gin server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	router.Run(port)
}

func getFieldNames(s interface{}) []string {
	// This function retrieves the field names from a struct
	// and returns them as a slice of strings
	var fields []string
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fields = append(fields, field.Name)
	}
	return fields
}
