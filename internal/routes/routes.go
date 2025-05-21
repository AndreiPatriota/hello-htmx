package routes

import (
	"net/http"

	"github.com/AndreiPatriota/hello-htmx/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/public"))))

	// Routes
	r.Get("/", handlers.GetHomePage)
	r.Get("/sobre", handlers.GetSobrePage)

	r.Route("/tarefas", func(r chi.Router) {
		r.Get("/", handlers.GetTarefas)
		r.Post("/", handlers.PostTarefas)
	})
	r.Route("/tarefas/{id}", func(r chi.Router) {
		r.Get("/edita", handlers.GetTarefasIdEdita)
		r.Get("/", handlers.GetTarefasId)
		r.Put("/", handlers.PutTarefasId)
		r.Patch("/", handlers.PatchTarefasId)
		r.Delete("/", handlers.DeleteTarefasId)
	})
	
	return r
}