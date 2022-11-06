package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tokatu4561/tasks/pkg/application/di"
	"github.com/tokatu4561/tasks/pkg/domain"
)

// TODO: env管理する
const AWS_REGION = "ap-northeast-1"
const DYNAMO_ENDPOINT = "http://dynamodb:8000"

type Response struct {
	Task domain.Task `json:"task"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctl := di.NewTaskController()

	task, err := ctl.CreateTask(request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	response := Response{
		Task: *task,
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
