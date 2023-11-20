package cmd

import (
	"context"
	"log"
	"os"

	"github.com/Nansky/custom-fetch/app"
	"github.com/Nansky/custom-fetch/helper"
	"github.com/Nansky/custom-fetch/pkg/client"

	"github.com/spf13/cobra"
)

var (
	// store if there is metadata url
	metadataUrl string

	// rootCmd store command state if we are using fetch command param
	rootCmd = &cobra.Command{
		Use:   "fetch",
		Short: "fetch html data from urls",
		Long: `fetch html data from urls and then store the content in html file. this has metadata flag to get ` +
			`site information, number of links, total images and last fetched date`,
		Run: func(cmd *cobra.Command, args []string) {
			// function that executed if we run this command
			fetchUrls(args)
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

// fetchUrls called as default command of this program
// param (optional) :
//
//	--metadata, -m
func fetchUrls(args []string) {
	ctx := context.Background()
	httpClient := client.NewClient(helper.HttpTimeout)
	fetchCli := app.NewFetchImplementor(httpClient)

	// if metadata flag is set
	if metadataUrl != "" {
		fetchCli.FetchPageMetadata(ctx, metadataUrl)
		return
	}

	// execute fetch url
	fetchCli.FetchPages(ctx, args...)
}

// init to add flag metadata in command line
func init() {
	rootCmd.Flags().StringVarP(&metadataUrl, "metadata", "m", "", "use to fetch metadata of specific url")
}
