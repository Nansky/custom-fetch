package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nansky/custom-fetch/helper"
	"github.com/Nansky/custom-fetch/internal/app/handler"
	"github.com/spf13/cobra"
)

var (
	mirrorCmd = &cobra.Command{
		Use:   "mirror",
		Short: "Access mirror site",
		Long:  "Access mirror site from fetched static data",
		Run: func(cmd *cobra.Command, args []string) {
			// function that executed if we run this command
			startMirror()
		},
	}
)

func init() {
	rootCmd.AddCommand(mirrorCmd)
}

func startMirror() {
	mux := http.NewServeMux()

	// Handle requests to the root URL ("/") with the MirrorHandler function
	mux.HandleFunc("/", handler.MirrorHandler)

	// http.HandleFunc("/", handler.MirrorHandler)
	// Set up and start the HTTP server on specified port

	log.Printf("Server started on :%d...\n", helper.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", helper.Port), mux)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
