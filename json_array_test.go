package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Vito",
		MiddleName: "Jihad",
		LastName:   "Abdillah",
		Age:        23,
		Married:    false,
		Hobbies:    []string{"Gaming", "Hiking", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Vito","MiddleName":"Jihad","LastName":"Abdillah","Age":23,"Married":false,"Hobbies":["Gaming","Hiking","Coding"]}`
	jsonBytes := []byte(jsonString)

	Customer := &Customer{}
	err := json.Unmarshal(jsonBytes, Customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(Customer)
	fmt.Println(Customer.FirstName)
	fmt.Println(Customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Vito",
		MiddleName: "Jihad",
		LastName:   "Abdillah",
		Age:        23,
		Married:    false,
		Hobbies:    []string{"Gaming", "Hiking", "Coding"},
		Addresses: []Address{
			{
				Street:     "Jl Layang",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
			{
				Street:     "Jl Tidak Melayang",
				Country:    "Amerika",
				PostalCode: "8888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Vito","MiddleName":"Jihad","LastName":"Abdillah","Age":23,"Married":false,"Hobbies":["Gaming","Hiking","Coding"],"Addresses":[{"Street":"Jl Layang","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jl Tidak Melayang","Country":"Amerika","PostalCode":"8888"}]}`
	jsonBytes := []byte(jsonString)

	Customer := &Customer{}
	err := json.Unmarshal(jsonBytes, Customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(Customer)
	fmt.Println(Customer.FirstName)
	fmt.Println(Customer.Addresses)
}
