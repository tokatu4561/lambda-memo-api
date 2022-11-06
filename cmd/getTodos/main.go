package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tokatu4561/tasks/pkg/application"
	"github.com/tokatu4561/tasks/pkg/infrastructure/dynamo"
	"github.com/tokatu4561/tasks/pkg/usecases"
)

type Memo struct {
	MemoID    string `dynamo:"MemoID,hash"`
	Text      string `dynamo:"Text"`
	CreatedAt string `dynamo:"CreatedAt"`
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

	ctl.GetTasks()
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", string("hello")),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
