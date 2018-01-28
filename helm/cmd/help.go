package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(helpCmd)
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print help",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		help := `Plugins for using Architect API in helm
This is a place for @cznewt to create nice readme/help`
		fmt.Println(help)
	},
}
