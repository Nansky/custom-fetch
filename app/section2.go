package app

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Nansky/custom-fetch/helper"
	"github.com/Nansky/custom-fetch/pkg/client"

	"github.com/PuerkitoBio/goquery"
)

// FetchPageMetadata will fetch and shows specific Metadata URL informations
func (fc FetchCli) FetchPageMetadata(ctx context.Context, url string) {
	// check if url is valid
	if isValid := helper.IsValidURLs(url); !isValid {
		log.Fatal("cannot fetch invalid URL metadata")
		return
	}

	// send GET request to valid url
	res, _ := fc.httpCli.Get(url)
	defer res.Body.Close()

	// check http status code
	if res.StatusCode != http.StatusOK {
		log.Fatalf("HTTP request failed with status: %d %s", res.StatusCode, res.Status)
		return
	}

	// read http response and store it in goquery-doc struct
	doc, errDoc := goquery.NewDocumentFromReader(res.Body)
	if errDoc != nil {
		// error when failed when store to doc
		log.Fatalf("%v", errDoc)
		return
	}

	// Print required Metadata
	log.Println("site:", strings.Split(url, "//")[1])
	client.GetPageTotalInfo(doc)
	log.Println("last_fetch:", time.Now().Format(time.RFC1123))
}
