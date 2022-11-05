package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/line/line-bot-sdk-go/linebot"
)

// TODO: env管理する
const AWS_REGION = "ap-northeast-1"
const DYNAMO_ENDPOINT = "http://dynamodb:8000"

type Memo struct {
	MemoID    string `dynamo:"MemoID,hash"`
	Text      string `dynamo:"Text"`
	CreatedAt string `dynamo:"CreatedAt"`
}

type Line struct {
	ChannelSecret string
	ChannelToken  string
	Client        *linebot.Client
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, _ := setUpDB()

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", string("hello")),
		StatusCode: 200,
	}, nil
}

func GetAllMemo(db *dynamo.DB, memoID string) (*Memo, error) {
	table := db.Table("Momo")

	var memo Memo

	err := table.Get("UserID", memoID).All()
	if err != nil {
		return nil, err
	}

	return &memo, err
}

func GetMemo(db *dynamo.DB, memoID string) (*Memo, error) {
	table := db.Table("Momo")

	var memo Memo

	err := table.Get("UserID", memoID).One(&memo)
	if err != nil {
		return nil, err
	}

	return &memo, err
}

func SaveMemo(db *dynamo.DB, text string) error {
	table := db.Table("Momo")

	err := table.Put(&Memo{MemoID: "1234", Text: text, CreatedAt: "sss"}).Run()
	if err != nil {
		return err
	}

	return nil
}

func setUpDB() (*dynamo.DB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Endpoint:    aws.String(DYNAMO_ENDPOINT),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		return nil, err
	}

	db := dynamo.New(sess)

	return db, nil
}

func main() {
	lambda.Start(handler)
}
