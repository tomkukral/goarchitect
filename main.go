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
	var apiUrl, inv string

	apiUrl, inv = readConfig()

	url := fmt.Sprintf(
		"%s/inventory/v1/%s/data.json?source=%s",
		apiUrl, inv, "hovno",
	)

	fmt.Println(url)
}
