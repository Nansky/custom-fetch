package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Nansky/custom-fetch/helper"
)

func MirrorHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "url" parameter from the query string
	urlParam := r.URL.Query().Get("url")

	// Check if the "url" parameter is provided
	if urlParam == "" || !helper.IsValidURLs(urlParam) {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	u, _ := url.Parse(urlParam)

	// Create a file server for the mirrored content
	http.ServeFile(w, r, helper.MirroredDir+fmt.Sprintf("%s.html", u.Host))
}
