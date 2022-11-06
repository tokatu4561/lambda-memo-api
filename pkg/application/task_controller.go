package application

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tokatu4561/tasks/pkg/domain"
)

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

func (t *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.taskUsecase.GetTasks()

	if err != nil {
		t.badRequest(w, err)
	}

	t.writeJson(w, http.StatusOK, tasks)
}

func (t *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.Form.Get("user_id"))
	task := &domain.Task{
		UserID: userID,
		Title:  r.Form.Get("title"),
	}

	t.taskUsecase.AddTask(task)

	t.writeJson(w, http.StatusOK, nil)
}

func (c *TaskController) readJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func (c *TaskController) writeJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(out)

	return nil
}

func (t *TaskController) badRequest(w http.ResponseWriter, err error) error {
	var payload struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	payload.Error = true
	payload.Message = err.Error()

	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(out)

	return nil
}
