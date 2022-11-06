package domain

import "time"

type Task struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `db:"title" json:"title"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type TaskRepositoryInterface interface {
	AddTask(task *Task) (*Task, error)
	GetTasks() ([]*Task, error)
}

type TaskUseCaseInterface interface {
	AddTask(t *Task) (*Task, error)
	GetTasks() ([]*Task, error)
}
