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

const env_url string = "ARCHITECT_INVENTORY_API_URL"
const env_inv string = "ARCHITECT_INVENTORY_NAME"
const defaultUrl string = "https://localhost:8181"
const defaultInv string = "default"

type osInterface interface {
	Getenv(name string) string
	FlagParse()
	FlagArgs() []string
	FlagArg(pos int) string
	LogFatal(v ...interface{})
	HttpDo(req *http.Request) (*http.Response, error)
}

type RealOs struct {
}

func (o RealOs) Getenv(name string) string {
	return os.Getenv(name)
}
func (o RealOs) FlagParse() {
	flag.Parse()
}
func (o RealOs) FlagArgs() []string {
	return flag.Args()
}

func (o RealOs) FlagArg(pos int) string {
	return flag.Arg(pos)
}

func (o RealOs) LogFatal(v ...interface{}) {
	log.Fatal(v)
}

func (o RealOs) HttpDo(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	return client.Do(req)
}

func config(o osInterface) (string, string) {
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

func resourceNameArg(o osInterface) string {
	o.FlagParse()

	if len(o.FlagArgs()) < 1 {
		o.LogFatal("Missing resource name parameter")
	}

	return o.FlagArg(0)
}

func getUrl(apiUrl, inv, resourceName string) string {
	return fmt.Sprintf(
		"%s/inventory/v1/%s/%s/data.json?source=%s",
		apiUrl, inv, resourceName, "salt-pillar",
	)

	// resource names: ansible-inventory, salt-pillar, salt-top
}

func RunCmd(cmd string, o osInterface) {
	c := Client{osInterface: o}
	c.Configure()
	c.Output(cmd)
}

type Client struct {
	apiURL      string
	inventory   string
	osInterface osInterface
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

	fmt.Println(resp)

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

	resourceName := resourceNameArg(c.osInterface)
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
