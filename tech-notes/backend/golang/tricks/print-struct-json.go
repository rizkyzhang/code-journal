package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	product := struct {
		Name        string
		Price       int
		IsAvailable bool
	}{
		Name:        "Phone",
		Price:       9000,
		IsAvailable: true,
	}

	productJson, err := json.Marshal(product)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(string(productJson))

	// Pretty print with indent
	productJsonPretty, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(string(productJsonPretty))
}
