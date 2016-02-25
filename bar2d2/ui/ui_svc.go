package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var suffixMap = map[string]string{
	"html": "text/html",
	"htm":  "text/html",
	"js":   "application/javascript",
	"css":  "text/css",
	"eot":  "application/vnd.ms-fontobject",
	"woff": "application/font-woff",
	"tff":  "application/x-font-truetype",
	"svg":  "image/svg+xml",
	"otf":  "application/x-font-opentype",
}

func UIHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/ui/"):]
	filename := docPath + title
	fmt.Printf("Path: %s.\n", filename)
	body, _ := ioutil.ReadFile(filename)
	//fmt.Printf("Body: %s.\n", body)

	// Get the suffix
	asuffix := strings.Split(title, ".")
	suffix := asuffix[len(asuffix)-1]

	fmt.Printf("SUFFIX: %s\n", suffix)
	mimeType := suffixMap[suffix]
	fmt.Printf("MIME: %s\n", mimeType)

	w.Header().Set("Content-Type", mimeType)
	fmt.Fprintf(w, "%s", body)
}
