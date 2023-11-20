package helper

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

// IsValidURLs will check whether rawURL is valid or not. it receives rawURL string as parameter
func IsValidURLs(rawURL string) bool {
	// ParseRequestURI parses a raw url into a URL structure. It assumes that url was received in an HTTP request
	_, err := url.ParseRequestURI(rawURL)

	return err == nil
}

// SetFilenameFromURL will return file name same as url name with .html extension
func SetFilenameFromURL(url string) string {
	filename := strings.Split(url, "//")

	// skip first domain
	return fmt.Sprintf("%s.html", filename[1])
}

func SaveToHtmlFile(filename, content string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	_, errWrite := file.WriteString(content)
	if errWrite != nil {
		log.Fatal(errWrite)
	}

	errs := os.Chmod(filename, 0777)
	if errs != nil {
		return
	}
}
