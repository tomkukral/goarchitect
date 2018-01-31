package cmd

import (
	"github.com/cznewt/goarchitect"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short:   "Architect helm plugin",
	Version: goarchitect.Version(),
}

func Execute() {
	rootCmd.Execute()
}
