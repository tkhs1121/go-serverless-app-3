package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := parseRequest(request.Body)
	if err != nil {
		return errorResponse(err, 500)
	}
	if err := checkAmazonURL(req.URL); err != nil {
		return errorResponse(err, 500)
	}

	db, err := newDynamoDB()
	if err != nil {
		return errorResponse(err, 500)
	}
	if err := db.putAmazonURL(req.URL); err != nil {
		return errorResponse(err, 500)
	}

	return events.APIGatewayProxyResponse{
		Body:       "Success",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
