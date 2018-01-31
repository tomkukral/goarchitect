package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cznewt/goarchitect"
	"github.com/pborman/getopt"
)

var out io.Writer = os.Stdout
var osInt goarchitect.OsInterface = goarchitect.RealOs{}

func main() {
	// define command options
	list := getopt.BoolLong("list", 'l', "List hosts and groups")
	host := getopt.StringLong("host", 'h', "", "Inventory hostname")

	getopt.Parse()

	// decide on command
	if *list {
		cmd := "ansible-inventory-list"
		fmt.Fprintf(out, goarchitect.RunCmd(cmd, "", osInt))
	} else {
		printParameters(*host, osInt)
	}

}

func printParameters(host string, osInt goarchitect.OsInterface) {
	if host == "" {
		log.Fatal("Missing host parameter, please set --host")
	}

	cmd := "ansible-inventory-host"
	fmt.Fprintf(out, goarchitect.RunCmd(cmd, host, osInt))
}
