package application

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/tokatu4561/tasks/pkg/domain"
)

type Task struct {
	ID        string `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TaskUsecase domain.TaskUseCaseInterface
type TaskRepository domain.TaskRepositoryInterface
type TaskController struct {
	taskUsecase TaskUsecase
}

func NewTaskController(taskUsecase domain.TaskUseCaseInterface) *TaskController {
	return &TaskController{
		taskUsecase: taskUsecase,
	}
}

func (t *TaskController) GetTasks(request events.APIGatewayProxyRequest) ([]*domain.Task, error) {
	tasks, err := t.taskUsecase.GetTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskController) GetTask(request events.APIGatewayProxyRequest) (*domain.Task, error) {
	id := request.PathParameters["id"]

	task, err := t.taskUsecase.GetTask(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskController) CreateTask(request events.APIGatewayProxyRequest) (*domain.Task, error) {
	type RequestPayload struct {
		Task Task `json:"task"`
	}
	var requestPayload RequestPayload
	t.readJson(request, &requestPayload)

	newId := uuid.New()

	task := domain.Task{
		ID:        newId.String(),
		UserID:    requestPayload.Task.UserID,
		Title:     requestPayload.Task.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newTask, err := t.taskUsecase.AddTask(&task)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (t *TaskController) UpdateTask(request events.APIGatewayProxyRequest) (*domain.Task, error) {
	type RequestPayload struct {
		Task Task
	}
	var requestPayload RequestPayload
	t.readJson(request, &requestPayload)

	task := domain.Task{
		ID:        requestPayload.Task.ID,
		UserID:    requestPayload.Task.UserID,
		Title:     requestPayload.Task.Title,
		UpdatedAt: time.Now(),
	}

	updatedTask, err := t.taskUsecase.UpdateTask(&task)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (t *TaskController) DeleteTask(request events.APIGatewayProxyRequest) error {
	type RequestPayload struct {
		Task Task `json:"task"`
	}
	var requestPayload RequestPayload
	t.readJson(request, &requestPayload)

	task := domain.Task{
		ID:     requestPayload.Task.ID,
		UserID: requestPayload.Task.UserID,
		Title:  requestPayload.Task.Title,
	}

	err := t.taskUsecase.DeleteTask(&task)
	if err != nil {
		return err
	}

	return nil
}

func (c *TaskController) readJson(req events.APIGatewayProxyRequest, data interface{}) error {
	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		return err
	}

	return nil
}

// func (c *TaskController) readJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
// 	maxBytes := 1048576

// 	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

// 	dec := json.NewDecoder(r.Body)
// 	err := dec.Decode(data)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (c *TaskController) writeJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
// 	out, err := json.MarshalIndent(data, "", "\t")
// 	if err != nil {
// 		return err
// 	}

// 	if len(headers) > 0 {
// 		for k, v := range headers[0] {
// 			w.Header()[k] = v
// 		}
// 	}

// 	w.Header().Set("Content-type", "application/json")
// 	w.WriteHeader(status)
// 	w.Write(out)

// 	return nil
// }

// func (t *TaskController) badRequest(w http.ResponseWriter, err error) error {
// 	var payload struct {
// 		Error   bool   `json:"error"`
// 		Message string `json:"message"`
// 	}

// 	payload.Error = true
// 	payload.Message = err.Error()

// 	out, err := json.MarshalIndent(payload, "", "\t")
// 	if err != nil {
// 		return err
// 	}

// 	w.Header().Set("Content-type", "application/json")
// 	w.Write(out)

// 	return nil
// }
