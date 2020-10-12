package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

type body struct {
	Message string `json:"message"`
}

type message struct {
	Message string `json:"message"`
}

// Handle is
func Handle(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("api-trigger handler ran")

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	})
	connectionID, ok := request.PathParameters["connectionID"]
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "connectionID must be provided",
		}, nil
	}
	var b body
	err := json.Unmarshal([]byte(request.Body), &b)
	if err != nil {
		fmt.Println("body unmarshal error")
		fmt.Println(err)
	}
	m, err := json.Marshal(message{
		Message: b.Message,
	})
	if err != nil {
		fmt.Println("marshalling error")
		fmt.Println(err)
	}

	url := fmt.Sprintf("https://mudkgw5pwb.execute-api.eu-west-2.amazonaws.com/%s", request.RequestContext.Stage)
	apigw := apigatewaymanagementapi.New(sess, &aws.Config{
		Endpoint: aws.String(url),
	})
	_, err = apigw.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(connectionID),
		Data:         []byte(m),
	})
	if err != nil {
		fmt.Println("error")
		fmt.Println(err.Error())
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
