package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Vito","MiddleName":"Jihad","LastName":"Abdillah","Age":23,"Married":false}`
	jsonBytes := []byte(jsonString)

	Customer := &Customer{}

	err := json.Unmarshal(jsonBytes, Customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(Customer)
	fmt.Println(Customer.FirstName)
	fmt.Println(Customer.MiddleName)
	fmt.Println(Customer.LastName)
}
