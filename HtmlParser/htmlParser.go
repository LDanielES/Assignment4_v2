package Htmlparser

import (
	"bytes"
	"net/http"
	"text/template"
)

// Parse HTML file
var HtmlParser = template.Must(template.ParseFiles("/Frontend/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := HtmlParser.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}
