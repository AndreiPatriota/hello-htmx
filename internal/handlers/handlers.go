package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AndreiPatriota/hello-htmx/internal/models"
	"github.com/go-chi/chi/v5"
)


func GetIndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/home")
	w.WriteHeader(http.StatusFound)
}

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	renderPage("home", w, nil)
}

func GetSobrePage(w http.ResponseWriter, r *http.Request) {
	renderPage("sobre", w, nil)
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

	if err := models.CreateTarefa(novaTarefa); err != nil {
		log.Println("Erro ao criar tarefa:", err)
		http.Error(w, "Erro ao criar tarefa", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	sendFragments(w, novaTarefa, "tarefas-id")
}

func GetTarefas(w http.ResponseWriter, r *http.Request) {
	tarefas, err := models.RetrieveTarefas()
	if err != nil {
		log.Println("Erro ao carregar tarefas:", err)
		http.Error(w, "Erro ao carregar tarefas", http.StatusInternalServerError)
		return
	}

	data := struct {
		Tarefas []models.Tarefa
	} {
		Tarefas: tarefas,
	}
	
	sendFragments(w, data, "tarefas", "tarefas-id")
}
func GetTarefasId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	tarefa, err := models.RetrieveTarefaById(id)
	if err != nil {
		log.Println("Erro ao carregar tarefa:", err)
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	sendFragments(w, tarefa, "tarefas-id")
}

func PatchTarefasId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	tarefa, err := models.RetrieveTarefaById(id)
	if err != nil {
		log.Println("Erro ao carregar tarefa:", err)
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	tarefa.Concluida = !tarefa.Concluida
	
	if err := models.UpdateTarefa(tarefa); err != nil {
		log.Println("Erro ao atualizar tarefa:", err)
		http.Error(w, "Erro ao atualizar tarefa", http.StatusInternalServerError)
		return
	}

	sendFragments(w, tarefa, "tarefas-id")
}

func DeleteTarefasId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := models.DeleteTarefa(id); err != nil {
		log.Println("Erro ao deletar tarefa:", err)
		http.Error(w, "Erro ao deletar tarefa", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetTarefasIdEdita(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	tarefa, err := models.RetrieveTarefaById(id)
	if err != nil {
		log.Println("Erro ao carregar tarefa:", err)
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	sendFragments(w, tarefa, "tarefas-id-edita")
}

func PutTarefasId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	tarefa, err := models.RetrieveTarefaById(id)
	if err != nil {
		log.Println("Erro ao carregar tarefa:", err)
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Não entendi os dados =/", http.StatusBadRequest)
		return
	}


	tarefa.Titulo = r.FormValue("titulo")
	tarefa.Descricao = r.FormValue("descricao")
	if err := models.UpdateTarefa(tarefa); err != nil {
		log.Println("Erro ao atualizar tarefa:", err)
		http.Error(w, "Erro ao atualizar tarefa", http.StatusInternalServerError)
		return
	}

	sendFragments(w, tarefa, "tarefas-id")
}