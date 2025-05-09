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
	fmt.Printf("Servidor em http://localhost:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		panic(err)
	}
}