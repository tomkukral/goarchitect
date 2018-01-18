package main

import architect "github.com/tomkukral/architect_client"

func main() {
	cmd := "ansible-inventory"

	c := architect.Client{}
	c.Configure(nil)
	c.Output(cmd)
}
