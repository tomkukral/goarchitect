package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cznewt/goarchitect"
	"github.com/spf13/cobra"
)

var osi goarchitect.OsInterface = goarchitect.RealOs{}

func init() {
	rootCmd.AddCommand(getValuesCmd)
	getValuesCmd.Flags().StringVarP(&valuesFilename, "filename", "f", "", "filename for values file")
}

var valuesFilename string

var getValuesCmd = &cobra.Command{
	Use:   "get-values",
	Short: "Get values for given resourceName",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resourceName := args[0]
		filename := getFilename(resourceName)

		giveFile(resourceName, filename)

	},
}

func giveFile(resourceName string, filename string) {
	// load values
	cmdName := "helm-architect"
	values := goarchitect.RunCmd(cmdName, resourceName, osi)

	// save values to file
	cb := []byte(values)
	err := ioutil.WriteFile(filename, cb, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// return filename
	fmt.Printf("%s", filename)
}

func getFilename(resourceName string) string {
	if valuesFilename != "" {
		return valuesFilename
	} else {
		return fmt.Sprintf("/tmp/ha-%s", resourceName)
	}
}
