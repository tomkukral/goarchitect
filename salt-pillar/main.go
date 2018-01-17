package main

import architect "github.com/tomkukral/architect_client"

func main() {
	res := "ansible-inventory"

	c := architect.Client{}
	c.Configure()
	c.Resource(res)
}
