package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type snsMessage struct {
	ConnectionID string `json:"connectionId"`
	CallbackURL  string `json:"callbackUrl"`
}

// Handle is
func Handle(_ context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("connection handler ran")

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	})
	
	m, err := json.Marshal(snsMessage{
		ConnectionID: request.RequestContext.ConnectionID,
		CallbackURL:  fmt.Sprintf("https://%s/%s", request.RequestContext.DomainName, request.RequestContext.Stage),
	})
	if err != nil {
		fmt.Println("marshalling error")
		fmt.Println(err)
	}
	snsclient := sns.New(sess)
	_, err = snsclient.Publish(&sns.PublishInput{
		TopicArn: aws.String(os.Getenv("TOPIC_ARN")),
		Message:  aws.String(string(m)),
	})
	if err != nil {
		fmt.Println("sns publish error")
		fmt.Println(err)
	}
	resp := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}
	return resp, nil
}

func main() {
	fmt.Println("handler starting")
	lambda.Start(Handle)
}
