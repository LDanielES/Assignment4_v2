package Htmlparser

import (
	"bytes"
	"net/http"
	"text/template"
)

// Parse HTML file
var Parser = template.Must(template.ParseFiles("index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := Parser.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}
