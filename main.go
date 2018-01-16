package main

import (
	"fmt"
	"os"
)

func readConfig() (string, string) {
	url := os.Getenv("ARCHITECT_INVENTORY_API_URL")
	if url == "" {
		url = "https://localhost:8181"
	}

	inv := os.Getenv("ARCHITECT_INVENTORY_NAME")
	if inv == "" {
		inv = "default"
	}

	return url, inv
}

func main() {
	var url, inv string

	url, inv = readConfig()

	fmt.Println(url, inv)
}
