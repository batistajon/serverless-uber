package handler

import (
	"context"
	"encoding/json"
	"goLambda/internal/types"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleEndRideLocal(ctx context.Context, event json.RawMessage) (events.LambdaFunctionURLResponse, error) {
	var receipt types.Receipt
	if err := json.Unmarshal(event, &receipt); err != nil {
		log.Printf("core: Failed to unmarshal event: %v", err)
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       "Failed to unmarshal payload",
		}, err
	}

	// TODO
	// open connection with dynamobd
	// defer dynamodb connection
	// get Uber credentials from user
	// save receipts information into dynamo receipts table
	// save into a google sheets spreadsheet

	log.Println("core: Lambda process has worked so far")
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       "core: Handler has worked so far",
	}, nil
}
