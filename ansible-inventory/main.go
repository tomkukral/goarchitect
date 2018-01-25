package main

import architect "github.com/cznewt/goarchitect"

func main() {
	cmd := "ansible-inventory"
	architect.RunCmd(cmd, nil)
}
