package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/tokatu4561/tasks/pkg/domain"
)

//FIXME:環境変数の管理
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
		// Endpoint:    aws.String(os.Getenv("DYNAMODB_ENDPOINT")), //FIXME: ローカル接続のためには必要？？s
		// Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
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

func (t *TaskRepositoryGateway) UpdateTask(task *domain.Task) (*domain.Task, error) {
	updatedTask, err := Update(t.databaseHandler, task)

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (t *TaskRepositoryGateway) DeleteTask(task *domain.Task) error {
	err := Delete(t.databaseHandler, task)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskRepositoryGateway) GetTask(id string) (*domain.Task, error) {
	task, err := Get(t.databaseHandler, id)
	if err != nil {
		return nil, err
	}

	return task, nil
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

	newTask := task

	err := table.Put(&newTask).Run()
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func Update(db *dynamo.DB, task *domain.Task) (*domain.Task, error) {
	table := db.Table("Task")

	var updatedTask *domain.Task

	err := table.Update("id", task.ID).Set("title", task.Title).Set("updatedAt", task.UpdatedAt).Value(&updatedTask)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func Delete(db *dynamo.DB, task *domain.Task) error {
	table := db.Table("Task")

	err := table.Delete("id", task.ID).Run()
	if err != nil {
		return err
	}

	return nil
}

func Get(db *dynamo.DB, id string) (*domain.Task, error) {
	table := db.Table("Task")

	var task *domain.Task

	err := table.Get("id", id).One(&task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func GetAll(db *dynamo.DB) ([]*domain.Task, error) {
	table := db.Table("Task")

	var tasks []*domain.Task

	err := table.Scan().All(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
