package architect_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetUrl(t *testing.T) {
	apiUrl := "https://localhost"
	inv := "testing"
	resourceName := "burger"

	url := getUrl(apiUrl, inv, resourceName)
	req := fmt.Sprintf(
		"%s/inventory/v1/%s/%s/data.json?source=%s",
		apiUrl, inv, resourceName, "salt-pillar",
	)

	if req != url {
		t.Errorf("URL don't match. Required '%s', got '%s'", req, url)
	}

}

type TestOs struct {
	Args        []string
	EmptyConfig bool
	Status      string
	Body        string
}

func (o TestOs) Getenv(name string) string {
	var r string

	fmt.Println(o)

	if o.EmptyConfig {
		r = ""
	} else {
		r = fmt.Sprintf("valueof:%s", name)
	}

	return r
}

func (o TestOs) FlagParse() {

}
func (o TestOs) FlagArgs() []string {
	return o.Args
}

func (o TestOs) FlagArg(pos int) string {
	return o.Args[pos]
}

func (o TestOs) LogFatal(v ...interface{}) {
	o.Status = "log.Fatal"
}

func (o TestOs) HttpDo(req *http.Request) (*http.Response, error) {

	header := make(http.Header, 0)
	header.Add("Content-Type", "application/json")

	t := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(o.Body)),
		ContentLength: int64(len(o.Body)),
		Request:       req,
		Header:        header,
	}

	return t, nil
}

var testInterface = TestOs{
	Args: []string{"arg1", "arg2", "arg3"},
}

func TestConfig(t *testing.T) {

	url, inv := config(testInterface)

	fmt.Println(url, inv)

	if url != "valueof:ARCHITECT_INVENTORY_API_URL" {
		t.Errorf("Failed reading URL from env variable, got: %s", url)
	}

	if inv != "valueof:ARCHITECT_INVENTORY_NAME" {
		t.Errorf("Failed reading inventory from env variable, got: %s", url)
	}

}

func TestConfigEmpty(t *testing.T) {
	ti := testInterface
	ti.EmptyConfig = true

	url, inv := config(ti)

	if url != defaultUrl {
		t.Errorf("Failed reading URL from env variable, got: %s", url)
	}

	if inv != defaultInv {
		t.Errorf("Failed reading inventory from env variable, got: %s", url)
	}
}

func TestResourceNameArg(t *testing.T) {
	r := resourceNameArg(testInterface)

	if r != testInterface.Args[0] {
		t.Errorf("Failed reading args, got: %s, expected: %s", r, testInterface.Args[0])
	}

}

func TestClient(t *testing.T) {
	c := Client{osInterface: testInterface}

	c.Configure()

	url, inv := config(testInterface)

	if c.apiURL != url {
		t.Errorf("apiURL is wrong, got: %s, expected: %s", c.apiURL, url)
	}
	if c.inventory != inv {
		t.Errorf("inventory is wrong, got: %s, expected: %s", c.inventory, inv)
	}

}

func TestEnsureInterface(t *testing.T) {
	c := Client{}

	if c.osInterface != nil {
		t.Error("Client interface isn't nil")
	}

	c.ensureInterface()

	df := RealOs{}
	if c.osInterface != df {
		t.Error("Default interface wasn't assigned fro client")
	}
}

func TestKeyInMap(t *testing.T) {
	rm := json.RawMessage(`{"precomputed": true}`)
	key := "exists"

	td := make(map[string]*json.RawMessage)
	td[key] = &rm

	if !keyInMap(td, key) {
		t.Errorf("Key %s reported as non-existing", key)
	}

	key = "lemur"
	if keyInMap(td, key) {
		t.Errorf("Key %s reported as existing", key)
	}

}

func TestReadResource(t *testing.T) {
	rn := "server.domain"
	c := Client{
		osInterface: TestOs{
			Args:        []string{rn},
			EmptyConfig: true,
		},
	}

	c.Configure()
	res := c.ReadResource(rn)
	fmt.Println(res)
}

func TestOutput(t *testing.T) {
	sn := "server.name"
	tpl := `
	{
		"%s": {
			"parameters": {}
		}
	}
	`

	rb := fmt.Sprintf(tpl, sn)

	ti := TestOs{
		EmptyConfig: true,
		Args:        []string{sn},
		Body:        rb,
	}

	c := Client{osInterface: ti}
	c.Configure()

	fmt.Println(c)

	RunCmd("salt-pillar", ti)
}
