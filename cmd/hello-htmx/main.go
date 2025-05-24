package main

import (
	"fmt"
	"net/http"

	"github.com/AndreiPatriota/hello-htmx/internal/models"
	"github.com/AndreiPatriota/hello-htmx/internal/routes"
)

func main() {
	models.InitDb()
	r := routes.RegisterRoutes()

	port := 80
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
	
	fmt.Printf("Servidor em http://localhost:%d\n", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}