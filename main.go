package architect_client

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type osInterface interface {
	Getenv(name string) string
}

type RealOs struct{}

func (o RealOs) Getenv(name string) string {
	return os.Getenv(name)
}

func config(o osInterface) (string, string) {
	url := o.Getenv("ARCHITECT_INVENTORY_API_URL")
	if url == "" {
		url = "https://localhost:8181"
	}

	inv := o.Getenv("ARCHITECT_INVENTORY_NAME")
	if inv == "" {
		inv = "default"
	}

	return url, inv
}

func resourceNameArg() string {
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("Missing resource name parameter")
	}

	return flag.Arg(0)
}

func getUrl(apiUrl, inv, resourceName string) string {
	return fmt.Sprintf(
		"%s/inventory/v1/%s/%s/data.json?source=%s",
		apiUrl, inv, resourceName, "salt-pillar",
	)

	// resource names: ansible-inventory, salt-pillar, salt-top
}

type Client struct {
	apiURL      string
	inventory   string
	osInterface osInterface
}

func (c *Client) Configure(o osInterface) {
	if c.osInterface == nil {
		o = RealOs{}
	}

	apiURL, inv := config(o)

	c.apiURL = apiURL
	c.inventory = inv

}

func (c *Client) ReadResource(resourceName string) []byte {
	// calculcate URL
	url := getUrl(c.apiURL, c.inventory, resourceName)

	// prepare request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	// send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return body

}

func keyInMap(m map[string]*json.RawMessage, key string) bool {
	if _, ok := m[key]; ok {
		return true
	}

	return false
}

func (c *Client) Output(command string) {

	resourceName := resourceNameArg()
	data := c.ReadResource(resourceName)

	var jsonRoot map[string]*json.RawMessage
	err := json.Unmarshal(data, &jsonRoot)
	if err != nil {
		log.Fatal(err)
	}

	if !keyInMap(jsonRoot, resourceName) {
		log.Fatalf("Host %s is missing in server reponse", resourceName)
	}

	var hostRoot map[string]*json.RawMessage
	err = json.Unmarshal(*jsonRoot[resourceName], &hostRoot)
	if err != nil {
		log.Fatal(err)
	}

	if val, ok := hostRoot["parameters"]; ok {
		fmt.Printf("%s\n", *val)
	} else {
		fmt.Println("{}")
	}

}

func (c *Client) Resource(resourceName string) {
	r := c.ReadResource(resourceName)
	fmt.Printf("%s", r)
}
