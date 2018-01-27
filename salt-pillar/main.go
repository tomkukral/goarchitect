package main

import (
	architect "github.com/cznewt/goarchitect"
)

func main() {
	cmd := "salt-pillar"
	architect.RunCmd(cmd, nil)
}
