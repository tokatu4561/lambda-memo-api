package dynamo

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/tokatu4561/tasks/pkg/domain"
)

// TODO: env管理する
const AWS_REGION = "ap-northeast-1"
const DYNAMO_ENDPOINT = "http://dynamodb:8000"

type TaskRepositoryGateway struct {
	databaseHandler *dynamo.DB
}

type DatabaseHandler struct {
	Conn *dynamo.DB
}

func NewDynamoDatabaseHandler() *dynamo.DB {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Endpoint:    aws.String(DYNAMO_ENDPOINT),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})

	return dynamo.New(sess)
}

func NewTaskRepository(db *dynamo.DB) domain.TaskRepositoryInterface {
	return &TaskRepositoryGateway{
		databaseHandler: db,
	}
}

func (t *TaskRepositoryGateway) AddTask(task *domain.Task) (*domain.Task, error) {
	newTask, err := Insert(t.databaseHandler, task)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (t *TaskRepositoryGateway) GetTasks() ([]*domain.Task, error) {
	tasks, err := GetAll(t.databaseHandler)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func Insert(db *dynamo.DB, task *domain.Task) (*domain.Task, error) {
	table := db.Table("Task")

	err := table.Put(&domain.Task{ID: 1, UserID: 1, Title: task.Title, CreatedAt: time.Now(), UpdatedAt: time.Now()}).Run()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetALl returns all tasks in db
func GetAll(db *dynamo.DB) ([]*domain.Task, error) {
	table := db.Table("Momo")

	var task *domain.Task

	err := table.Get("ID").One(&task)
	if err != nil {
		return nil, err
	}

	return &task, err
}
