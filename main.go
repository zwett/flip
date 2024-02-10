package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "flip",
	Short:         "Flutter versioning, simplified.",
	Long:          "Flip between Flutter versions with ease.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
