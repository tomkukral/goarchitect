package main

import (
	architect "github.com/tomkukral/architect_client"
)

func main() {
	cmd := "salt-pillar"

	c := architect.Client{}
	c.Configure(nil)
	c.Output(cmd)
}
