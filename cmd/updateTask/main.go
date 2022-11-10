package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tokatu4561/tasks/pkg/application/di"
	"github.com/tokatu4561/tasks/pkg/domain"
)

type Response struct {
	Task domain.Task `json:"task"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Access-Control-Allow-Origin":     "*",
		"Access-Control-Allow-Methods":    "PUT",
		"Access-Control-Allow-Credential": "true",
		"Access-Control-Allow-Headers":    "Authorization,X-XSRF-TOKEN,Content-Type,ContentType,x-amz-security-token,x-amz-date",
	}

	ctl := di.NewTaskController()

	task, err := ctl.UpdateTask(request)
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
		Headers:    headers,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
