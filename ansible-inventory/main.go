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

func main() {
	// define command options
	list := getopt.BoolLong("list", 'l', "List hosts and groups")
	host := getopt.StringLong("host", 'h', "", "Inventory hostname")

	getopt.Parse()

	// decide on command
	if *list {
		fmt.Fprintf(out, "{}")
	} else {
		printParameters(*host, goarchitect.RealOs{})
	}

}

func printParameters(host string, osInt goarchitect.OsInterface) {
	if host == "" {
		log.Fatal("Missing host parameter, please set --host")
	}

	cmd := "ansible-inventory"
	fmt.Fprintf(out, goarchitect.RunCmd(cmd, host, osInt))
}
