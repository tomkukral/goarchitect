package architect_client

import (
	"fmt"
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
