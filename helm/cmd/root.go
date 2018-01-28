package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "helm-goarchitect",
	Short: "Get metrics from goarchitect",
	Long:  `Integrate Architect-API to helm and load values from API`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

		fmt.Println("rootCommand executed")
	},
}

func Execute() {
	rootCmd.Execute()
}
