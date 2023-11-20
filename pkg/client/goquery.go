package client

import (
	"log"
	"net/http"
	"sync"

	"github.com/Nansky/custom-fetch/helper"

	"github.com/PuerkitoBio/goquery"
	"github.com/yosssi/gohtml"
)

// GetPageTotalInfo get total page information on Images and Links
func GetPageTotalInfo(doc *goquery.Document) {
	log.Println("num_links", doc.Find("a").Length())
	log.Println("images:", doc.Find("img").Length())
}

func WriteToHtmlFile(c *Client, url string, wg *sync.WaitGroup, sig chan struct{}) {
	defer wg.Done()

	res, err := c.httpClient.Get(url)
	if err != nil {
		sig <- struct{}{}
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("HTTP request failed with status: %d %s", res.StatusCode, res.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	html, _ := doc.Selection.Html()
	formattedHtml := gohtml.Format(html)

	helper.SaveToHtmlFile(helper.SetFilenameFromURL(url), formattedHtml)
	sig <- struct{}{}
}
