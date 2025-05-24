package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)


func renderPage(pageName string, w http.ResponseWriter, data any) {
	templ := template.Must(template.ParseFiles("web/views/_layout.html", fmt.Sprintf("web/views/%s.html", pageName)))

	templ.Execute(w, data)
}

func sendFragments(w http.ResponseWriter, data any, fragmentNames ...string) {
	fragmentPaths := make([]string, 0, len(fragmentNames))
	for _, fragmentName := range fragmentNames {
		name := fmt.Sprintf("web/views/fragments/%s.html", fragmentName)
		fragmentPaths = append(fragmentPaths, name)
	}

	templ, err := template.ParseFiles(fragmentPaths...) 
	if err != nil {
		log.Println("Erro ao carregar templates:", fragmentPaths)
		http.Error(w, "Erro ao carregar templates", http.StatusInternalServerError)
		return
	}
	templ.ExecuteTemplate(w, fragmentNames[0], data)
}