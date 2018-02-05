package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/tomkukral/goarchitect"
	"github.com/stretchr/testify/assert"
)

func TestPrintParameters(t *testing.T) {
	host := "server.name"
	tpl := `{
		"%s": {
			"parameters": {
				"param1": "value1",
				"param2": "value2"
			}
		}
	}`

	ti := goarchitect.TestOs{
		EmptyConfig: true,
		Body:        fmt.Sprintf(tpl, host),
	}

	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	printParameters(host, ti)
	req := `{"param1":"value1","param2":"value2"}`

	printed := out.(*bytes.Buffer).String()

	assert.Equal(t, req, printed, "Got wrong output")
}
