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

func GetDB(dbType, dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(dbType, dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

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

func CreateCustomer(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&customer)
	c.JSON(http.StatusOK, customer)
}

func GetCustomer(c *gin.Context) {
	var customer Customer
	id := c.Param("id")

	if err := db.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}
