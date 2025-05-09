package handlers

import (
	"errors"
	"fmt"
	"html/template"
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

func renderFragment(fragmentName string, w http.ResponseWriter, data any) {
	templ := template.Must(template.ParseFiles(fmt.Sprintf("web/views/fragments/%s.html", fragmentName)))

	templ.Execute(w, data)
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

	// renderFragment("tarefa", w, novaTarefa)
	templ := template.Must(template.ParseFiles("web/views/fragments/tarefa.html"))

	templ.Execute(w, novaTarefa)
}

func GetTarefas(w http.ResponseWriter, r *http.Request) {
	var tarefas []models.Tarefa

	models.DB.Find(&tarefas)

	data := struct {
		Tarefas []models.Tarefa
	} {
		Tarefas: tarefas,
	}

	renderFragment("tarefas", w, data)
}

func PatchTarefasId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var tarefa models.Tarefa
	result := models.DB.First(&tarefa, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	tarefa.Concluida = !tarefa.Concluida
	models.DB.Save(&tarefa)

	renderFragment("tarefa", w, tarefa)
}