package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (

	// rootCmd store command state if we are using fetch command param
	rootCmd = &cobra.Command{
		Use:   "customfetch",
		Short: "customfetch is cli tool",
		Long: `customfetch is cli tool that can fetched static url data, shows url specific metadata, and ` +
			`can access its static content from generated fetched data`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

// Execute will execute command param
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
