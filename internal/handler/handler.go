package handler

import (
	"context"
	"encoding/json"
	"goLambda/internal/services"
	"goLambda/internal/types"
	"log"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleEndRideLocal(ctx context.Context, event json.RawMessage) error {
	var receipt types.Receipt
	if err := json.Unmarshal(event, &receipt); err != nil {
		log.Printf("core: Failed to unmarshal event: %v", err)
		return err
	}

	rt := types.ReceiptDTable{
		Name:        "Receipts",
		ReceiptId:   receipt.EventID,
		ReceiptData: receipt,
	}

	// TODO
	// open connection with dynamobd
	db, err := services.NewDynamoDBService(ctx)
	if err != nil {
		log.Printf("core: failed to authenticet with aws. Error: %v", err)
	}

	if err := db.Add(ctx, rt); err != nil {
		log.Printf("core: failed to add table. Error: %v", err)
		return err
	}

	// defer dynamodb connection
	// get Uber credentials from user
	// save receipts information into dynamo receipts table
	// save into a google sheets spreadsheet

	log.Println("core: Lambda process has worked so far")
	return nil
}
