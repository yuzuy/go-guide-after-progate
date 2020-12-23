package todo

import "time"

type Task struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
}
