package handlers

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	renderPage("_404", w)
}

func NotFoundHtmx(w http.ResponseWriter, r *http.Request) {
	renderPage("_404", w)
}
