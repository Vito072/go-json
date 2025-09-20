package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"W0001","name":"Apple Watch","price":3000}`
	jsonBytes := []byte(jsonString)

	var result map[string]any
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]any{
		"id":    "P0001",
		"name":  "Apple Iphone 17",
		"price": "1799",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
