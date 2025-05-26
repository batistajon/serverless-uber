package main

import (
	"context"
	"encoding/json"
	"goLambda/internal/handler"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlerLambda)
}

func handlerLambda(ctx context.Context, event json.RawMessage) (events.LambdaFunctionURLResponse, error) {
	coreHandler := handler.NewHandler()

	res, err := coreHandler.HandleEndRideLocal(ctx, event)
	if err != nil {
		log.Printf("Main: Core handler failed: %v", err)
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       "Main: Failed to call core handler",
		}, err
	}

	log.Println("main: Core handler has worked!")
	return res, nil
}
