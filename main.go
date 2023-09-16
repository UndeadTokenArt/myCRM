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

	// routes for displaying customer data. probalby
	router.POST("/customers", customerDB.CreateCustomer)
	router.GET("/customers/:id", customerDB.GetCustomer)

	// Generate a form with the Customer Struct's keys
	router.GET("/newCustomerForm", func(c *gin.Context) {
		clientKeys := customerDB.Customer{}
		fields := getFieldNames(clientKeys)

		c.HTML(http.StatusOK, "NewCustomerForm.tmpl", fields)
	})

	// router for binding newCustomer form data to Customer struct
	router.POST("/newCustomerForm", func(c *gin.Context) {
		var clientKeys customerDB.Customer
		c.Bind(&clientKeys)

		//Redirect after binding to DB
		c.Redirect(http.StatusSeeOther, "/success")
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
