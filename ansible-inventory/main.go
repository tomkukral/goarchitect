package main

import architect "github.com/cznewt/goarchitect"

func main() {
	cmd := "ansible-inventory"

	c := architect.Client{}
	c.Configure(nil)
	c.Output(cmd)
}
