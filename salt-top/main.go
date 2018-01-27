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
var host string
var osi goarchitect.OsInterface = goarchitect.RealOs{}

func main() {
	var args []string

	if host == "" {
		getopt.Parse()
		args = getopt.Args()

		if len(args) < 1 {
			log.Fatal("Missing host parameter")
		}

		host = args[0]
	}

	cmd := "salt-top"
	fmt.Fprintf(out, goarchitect.RunCmd(cmd, host, osi))
}
