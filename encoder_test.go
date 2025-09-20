package go_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("customerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Vito",
		MiddleName: "Jihad",
		LastName:   "Abdillah",
	}

	encoder.Encode(customer)
}
