package app

import (
	"context"
	"log"
	"sync"

	"github.com/Nansky/custom-fetch/helper"
	"github.com/Nansky/custom-fetch/pkg/client"
)

// FetchPages will fetch urls page and store to html file
func (fc FetchCli) FetchPages(ctx context.Context, urls ...string) {
	validUrls := make([]string, 0)

	// check if url is valid
	for _, url := range urls {
		if helper.IsValidURLs(url) {
			validUrls = append(validUrls, url)
		}
	}

	// if there are no valid urls to be fetched
	if len(validUrls) == 0 {
		log.Fatal("no valid url need to be fetched")
		return
	}

	// Processing fetch concurrently
	var wg sync.WaitGroup
	sig := make(chan struct{}, len(validUrls))

	for _, validUrl := range validUrls {
		wg.Add(1)
		go client.WriteToHtmlFile(fc.httpCli, validUrl, &wg, sig)
	}

	go func() {
		wg.Wait()
		close(sig)
	}()

	// get all signals after channel has closed
	for range sig {
		// waiting signal here. Do something if needed
		// log.Println("DONE")
	}
}
