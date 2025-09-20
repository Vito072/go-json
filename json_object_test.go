package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName, MiddleName, LastName string
	Age                             int
	Married                         bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Vito",
		MiddleName: "Jihad",
		LastName:   "Abdillah",
		Age:        23,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
