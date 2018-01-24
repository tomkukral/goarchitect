package main

import architect "github.com/cznewt/goarchitect"

func main() {
	cmd := "salt-top"

	c := architect.Client{}
	c.Configure(nil)
	c.Output(cmd)
}
