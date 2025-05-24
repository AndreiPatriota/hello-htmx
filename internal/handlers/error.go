package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	renderPage("_404", w, nil)
}

func HandleNotFoundHtmx(w http.ResponseWriter, r *http.Request) {
	renderPage("_404", w, nil)
}

func HandleServerErrorsHtmx(w http.ResponseWriter, r *http.Request) {
	mensagem := chi.URLParam(r, "mensagem")
	if mensagem == "" {
		mensagem = "Erro interno do servidor"
	}

	data := struct {
		Mensagem string
	} {
		Mensagem: mensagem,
	}

	renderPage("_500s", w, data)
}
