package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tokatu4561/tasks/pkg/application"
	"github.com/tokatu4561/tasks/pkg/domain"
	"github.com/tokatu4561/tasks/pkg/infrastructure/dynamo"
	"github.com/tokatu4561/tasks/pkg/usecases"
)

type Response struct {
	Tasks []*domain.Task `json: "tasks"`
}

func NewTaskController() *application.TaskController {
	db := dynamo.NewDynamoDatabaseHandler()
	taskRepositoryInterface := dynamo.NewTaskRepository(db)
	taskUseCaseInterface := usecases.NewTaskUsecase(taskRepositoryInterface)
	taskController := application.NewTaskController(taskUseCaseInterface)
	return taskController
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctl := NewTaskController()

	tasks, err := ctl.GetTasks(request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	response := Response{
		Tasks: tasks,
	}

	jsonByte, _ := json.MarshalIndent(response, "", "\t")

	return events.APIGatewayProxyResponse{
		Body:       string(jsonByte),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
