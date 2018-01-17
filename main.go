package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func config() (string, string) {
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

func args() string {
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("Missing resource name parameter")
	}

	return flag.Arg(0)
}

func getUrl(apiUrl, inv, resourceName string) string {
	return "https://api.chucknorris.io/jokes/random"

	return fmt.Sprintf(
		"%s/inventory/v1/%s/%s/data.json?source=%s",
		apiUrl, inv, resourceName, "salt-pillar",
	)

	// resource names: ansible-inventory, salt-pillar, salt-top
}

func main() {
	var apiUrl, inv, resourceName string

	apiUrl, inv = config()
	resourceName = args()
	url := getUrl(apiUrl, inv, resourceName)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", body)

}
