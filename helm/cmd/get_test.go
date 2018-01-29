package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFilename(t *testing.T) {
	td := []struct {
		name     string
		filename string
		req      string
	}{
		{"server.domain", "", "/tmp/ha-server.domain"},
		{"server.domain", "myfilename.rar", "myfilename.rar"},
	}

	for _, ts := range td {
		valuesFilename = ts.filename

		fn := getFilename(ts.name)

		assert.Equal(t, ts.req, fn, "Generated wrong filename")
	}

}
