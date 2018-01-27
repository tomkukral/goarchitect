package goarchitect

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUrl(t *testing.T) {
	apiUrl := "https://localhost"
	inv := "testing"
	resourceName := "burger"

	url := getUrl(apiUrl, inv, resourceName)
	req := fmt.Sprintf(
		"%s/inventory/v1/%s/%s/data.json?source=%s",
		apiUrl, inv, resourceName, "unknown",
	)

	assert.Equal(t, req, url, "Generated URL don't match")

	if req != url {
		t.Errorf("URL don't match. Required '%s', got '%s'", req, url)
	}

}

var testInterface = TestOs{
	Args: []string{"arg1", "arg2", "arg3"},
}

func TestConfig(t *testing.T) {

	url, inv := config(testInterface)

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
	c.ReadResource(rn)
	// TODO: check outptu
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

	printed := RunCmd("ansible-inventory", sn, ti)
	req := `{}`
	assert.Equal(t, req, printed, "Returned unexpected output")
}
