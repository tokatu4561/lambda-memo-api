package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tokatu4561/tasks/pkg/application/di"
	"github.com/tokatu4561/tasks/pkg/domain"
)

type Response struct {
	Tasks []*domain.Task `json:"tasks"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctl := di.NewTaskController()

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
