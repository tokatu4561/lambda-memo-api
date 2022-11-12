package domain

import "time"

// FIXEME: domainなのにdynamoの情報を書いてる
type Task struct {
	ID        string    `dynamo:"id" json:"id"`
	UserID    int       `dynamo:"userId" json:"user_id"`
	Title     string    `dynamo:"title" json:"title"`
	CreatedAt time.Time `dynamo:"createdAt" json:"created_at"`
	UpdatedAt time.Time `dynamo:"updatedAt" json:"updated_at"`
}

type TaskRepositoryInterface interface {
	AddTask(task *Task) (*Task, error)
	UpdateTask(task *Task) (*Task, error)
	DeleteTask(task *Task) error
	GetTask(id string) (*Task, error)
	GetTasks() ([]*Task, error)
}

type TaskUseCaseInterface interface {
	AddTask(t *Task) (*Task, error)
	UpdateTask(t *Task) (*Task, error)
	DeleteTask(t *Task) error
	GetTask(id string) (*Task, error)
	GetTasks() ([]*Task, error)
}
