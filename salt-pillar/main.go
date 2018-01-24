package main

import (
	architect "github.com/cznewt/goarchitect"
)

func main() {
	cmd := "salt-pillar"

	c := architect.Client{}
	c.Configure(nil)
	c.Output(cmd)
}
