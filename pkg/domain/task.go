package domain

import "time"

type Task struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskRepositoryInterface interface {
	AddTask(task *Task) (*Task, error)
	GetTasks() ([]*Task, error)
}

type TaskUseCaseInterface interface {
	AddTask(t *Task) (*Task, error)
	GetTasks() ([]*Task, error)
}
