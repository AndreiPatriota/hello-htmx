package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/AndreiPatriota/hello-htmx/internal/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)


func renderPage(pageName string, w http.ResponseWriter) {
	templ := template.Must(template.ParseFiles("web/views/_layout.html", fmt.Sprintf("web/views/%s.html", pageName)))

	templ.Execute(w, nil)
}

func sendMutipleFragment(w http.ResponseWriter, data any, fragmentNames ...string) {
	fragmentPaths := make([]string, len(fragmentNames))
	for _, fragmentName := range fragmentNames {
		fragmentPaths = append(fragmentPaths, fmt.Sprintf("web/views/fragments/%s.html", fragmentName))
	}

	templ := template.Must(template.ParseFiles(fragmentPaths...))
	templ.Execute(w, data)
}

func sendFragment(w http.ResponseWriter, data any, fragmentName string) {
	templ := template.Must(template.ParseFiles(fragmentName))

	templ.ExecuteTemplate(w, fragmentName, data)
}

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	renderPage("home", w)
}

func GetSobrePage(w http.ResponseWriter, r *http.Request) {
	renderPage("sobre", w)
}

func PostTarefas(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Não entendi os dados =/", http.StatusBadRequest)
		return
	}

	titulo := r.FormValue("titulo")
	descricao := r.FormValue("descricao")

	novaTarefa := &models.Tarefa{
		Titulo:   titulo,
		Descricao: descricao,
		Concluida: false,
	}
	models.DB.Create(novaTarefa)

	// sendFragment(w, novaTarefa, "tarefa")
	templ := template.Must(template.ParseFiles("web/views/fragments/tarefa.html"))
	templ.ExecuteTemplate(w, "tarefa", novaTarefa)
}

func GetTarefas(w http.ResponseWriter, r *http.Request) {
	var tarefas []models.Tarefa

	models.DB.Find(&tarefas)

	data := struct {
		Tarefas []models.Tarefa
	} {
		Tarefas: tarefas,
	}

	// templ := template.Must(template.ParseFiles("web/views/fragments/tarefas.html", "web/views/fragments/tarefa.html"))
	
	templ, err := template.ParseFiles("web/views/fragments/tarefas.html", "web/views/fragments/tarefa.html")
	if err != nil {
		log.Println("Erro ao carregar templates:", err)
		http.Error(w, "Erro ao carregar templates", http.StatusInternalServerError)
		return
	}
	templ.ExecuteTemplate(w, "tarefas", data)

	// sendMutipleFragment(w, data, "tarefas", "tarefa")
}

func PatchTarefasId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var tarefaTogada models.Tarefa
	result := models.DB.First(&tarefaTogada, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	tarefaTogada.Concluida = !tarefaTogada.Concluida
	models.DB.Save(&tarefaTogada)

	// sendFragment(w, tarefaTogada, "tarefa")
	templ := template.Must(template.ParseFiles("web/views/fragments/tarefa.html"))
	templ.ExecuteTemplate(w, "tarefa", tarefaTogada)
}

func DeleteTarefasId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var tarefa models.Tarefa
	result := models.DB.First(&tarefa, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	models.DB.Delete(&tarefa)

	w.WriteHeader(http.StatusNoContent)
}