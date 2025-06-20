package main

import (
	"net/http"

	_ "github.com/YourPainkiller/WorkMateTest/cmd/server/docs"
	"github.com/YourPainkiller/WorkMateTest/internal/handler"
	"github.com/YourPainkiller/WorkMateTest/internal/service"
	"github.com/YourPainkiller/WorkMateTest/internal/storage"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title		I/O bound task API
// @version		1.0
// @description	Manipulation with I/O bound tasks via API
func main() {
	r := chi.NewRouter()

	store := storage.NewMemoryStore()
	service := service.NewTaskService(store)
	handler := handler.NewTaskHandler(service)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Post("/tasks", handler.CreateTask)
	r.Get("/tasks/{id}", handler.GetTask)
	r.Delete("/tasks/{id}", handler.DeleteTask)

	http.ListenAndServe(":8080", r)
}
