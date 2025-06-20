package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"time"

	"github.com/YourPainkiller/WorkMateTest/internal/dto"
	"github.com/YourPainkiller/WorkMateTest/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// CreateTask godoc
//
//	@Summary		Create task
//	@Description	Create task
//	@Tags			tasks
//	@Produce		json
//	@Success		200	{object}	dto.CreateTaskResponse
//
// @Router			/tasks [post]
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	task := h.service.CreateTask(r.Context())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.CreateTaskResponse{ID: task.ID})
}

// GetTask godoc
//
//	@Summary		Get task
//	@Description	Get task by ID
//	@Tags			tasks
//	@Param			id	query string true "Task ID"
//	@Produce		json
//	@Success		200	{object}	dto.GetTaskResponse
//
// @Router			/tasks/{id} [get]
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task, err := h.service.GetTask(id)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoSuchTask):
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	response := dto.GetTaskResponse{
		ID:           task.ID,
		Status:       string(task.Status),
		CreatedAt:    task.CreatedAt.Format(time.RFC3339),
		DurationSecs: task.Duration().Seconds(),
	}

	if task.Error != nil {
		response.Error = task.Error.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteTask godoc
//
//	@Summary		Delete task
//	@Description	Delete task by ID
//	@Tags			tasks
//	@Param			id	query string true "Task ID"
//	@Produce		json
//	@Success		204	{object}	dto.Empty
//
// @Router			/tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := h.service.DeleteTask(id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoSuchTask):
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(dto.Empty{})
}
