package main

import architect "github.com/tomkukral/architect_client"

func main() {
	cmd := "salt-top"

	c := architect.Client{}
	c.Configure(nil)
	c.Output(cmd)
}
