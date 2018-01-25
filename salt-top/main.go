package main

import architect "github.com/cznewt/goarchitect"

func main() {
	cmd := "salt-top"
	architect.RunCmd(cmd, nil)
}
