package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tokatu4561/tasks/pkg/application/di"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Access-Control-Allow-Origin":     "*",
		"Access-Control-Allow-Methods":    "DELETE",
		"Access-Control-Allow-Credential": "true",
		"Access-Control-Allow-Headers":    "Authorization,X-XSRF-TOKEN,Content-Type,ContentType,x-amz-security-token,x-amz-date",
	}

	ctl := di.NewTaskController()

	err := ctl.DeleteTask(request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
			Headers:    headers,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    headers,
	}, nil
}

func main() {
	lambda.Start(handler)
}
