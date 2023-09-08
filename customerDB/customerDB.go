package customerDB

import (
	"reflect"

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

func main() {
	// Testing data
	myC := Customer{
		FirstName: "Bill",
		LastName:  "Zee",
		Email:     "BillIsGreat@email.com",
		Address:   "1234 NE Cool St.",
		Phone:     "503-555-7678",
		Buyer:     false,
		Seller:    true,
		Stage:     "inContract",
	}

	// Testing Functions
	fieldNames := GetStructFieldNames(myC)

	// Print the field names
	for _, fieldName := range fieldNames {
		println(fieldName)
	}
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
