package dto

type CreateTaskResponse struct {
	ID string `json:"id"`
}

type GetTaskResponse struct {
	ID           string  `json:"id"`
	Status       string  `json:"status"`
	CreatedAt    string  `json:"created_at"`
	DurationSecs float64 `json:"duration_secs"`
	Error        string  `json:"error,omitempty"`
}

type Empty struct{}
