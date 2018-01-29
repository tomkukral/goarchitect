package goarchitect

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const env_url string = "ARCHITECT_INVENTORY_API_URL"
const env_inv string = "ARCHITECT_INVENTORY_NAME"
const defaultUrl string = "http://localhost:8181"
const defaultInv string = "default"

var version = "master"

type OsInterface interface {
	Getenv(name string) string
	LogFatal(v ...interface{})
	HttpDo(req *http.Request) (*http.Response, error)
}

func config(o OsInterface) (string, string) {
	url := o.Getenv(env_url)
	if url == "" {
		url = defaultUrl
	}

	inv := o.Getenv(env_inv)
	if inv == "" {
		inv = defaultInv
	}

	return url, inv
}

func getUrl(apiUrl, inv, resourceName string) string {
	return fmt.Sprintf(
		"%s/inventory/v1/%s/%s/data.json?source=%s",
		apiUrl, inv, resourceName, "unknown",
	)

	// resource names: ansible-inventory, salt-pillar, salt-top
}

func RunCmd(cmd string, hostname string, o OsInterface) string {
	c := Client{hostname: hostname, osInterface: o}
	c.Configure()

	return c.Output(cmd, hostname)
}

type Client struct {
	hostname    string
	apiURL      string
	inventory   string
	osInterface OsInterface
}

func (c *Client) ensureInterface() {
	if c.osInterface == nil {
		c.osInterface = RealOs{}
	}

}

func (c *Client) Configure() {
	c.ensureInterface()

	apiURL, inv := config(c.osInterface)

	c.apiURL = apiURL
	c.inventory = inv
}

func (c *Client) ReadResource(resourceName string) []byte {
	// calculcate URL
	url := getUrl(c.apiURL, c.inventory, resourceName)

	// prepare request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	// send request
	resp, err := c.osInterface.HttpDo(req)
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

func (c *Client) Output(command string, resourceName string) string {

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

	if command == "ansible-inventory" || command == "salt-pillar" {

		if val, ok := hostRoot["parameters"]; ok {
			fr, err := json.Marshal(val)
			if err != nil {
				log.Fatal("Unable to format host parameters")
			}

			return string(fr)
		}
	}

	if command == "salt-top" {
		if val, ok := hostRoot["applications"]; ok {
			fr, err := json.Marshal(val)
			if err != nil {
				log.Fatal("Unable to format host parameters")
			}

			return fmt.Sprintf("{\"classes\": %s}", string(fr))
		}

	}

	return fmt.Sprintf("{}")
}
