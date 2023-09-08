package customerDB

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Define a model struct for your CRM data (e.g., Customer)
type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Address   string
	Phone     string
	Buyer     bool
	Seller    bool
	Stage     string
	// Add more fields as needed
}

var db *gorm.DB

func GetStructFieldNames(s interface{}) []string {
	var fieldNames []string

	// Get the type of the struct
	st := reflect.TypeOf(s)

	// Make sure s is a struct
	if st.Kind() != reflect.Struct {
		return fieldNames
	}

	// Iterate through the fields and add their names to the slice
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames
}

func initDB() {
	// sOpen a connection to the SQLite database
	var err error
	db, err = gorm.Open("sqlite3", "crm.db") // Change "mydb.sqlite" to your desired database file name
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migrate the schema
	db.AutoMigrate(&Customer{})
}

func createCustomer(c *gin.Context) {
	// Stubbed code for creating a customer
	// Parse the request body to get customer data
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the customer to the database
	db.Create(&customer)

	// Return the created customer
	c.JSON(http.StatusOK, customer)
}

func getCustomer(c *gin.Context) {
	// Stubbed code for retrieving a customer by ID
	// Fetch the customer from the database by ID (replace ":id" with the actual ID)
	var customer Customer
	id := c.Param("id")
	db.First(&customer, id)

	// Check if the customer exists
	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	// Return the customer
	c.JSON(http.StatusOK, customer)
}

func updateCustomer(c *gin.Context) {
	// Stubbed code for updating a customer by ID
	// Fetch the customer from the database by ID (replace ":id" with the actual ID)
	var customer Customer
	id := c.Param("id")
	db.First(&customer, id)

	// Check if the customer exists
	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	// Stubbed code to update customer data (you can replace these fields as needed)
	var updatedCustomer Customer
	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the customer data in the database
	db.Model(&customer).Updates(updatedCustomer)

	// Return the updated customer
	c.JSON(http.StatusOK, customer)
}

func deleteCustomer(c *gin.Context) {
	// Stubbed code for deleting a customer by ID
	// Fetch the customer from the database by ID (replace ":id" with the actual ID)
	var customer Customer
	id := c.Param("id")
	db.First(&customer, id)

	// Check if the customer exists
	if customer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	// Delete the customer from the database
	db.Delete(&customer)

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
