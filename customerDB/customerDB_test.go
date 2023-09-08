package customerDB_test

import (
	"fmt"
	"customerDB"
)

// Testing data
myC := Customer {
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