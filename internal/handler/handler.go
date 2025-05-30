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

	db, err := services.NewDynamoDBService(ctx)
	if err != nil {
		log.Printf("Core: Failed to authenticet with AWS DynamoDB. Error: %v", err)
		return err
	}

	ut := types.User{
		Name:   "Users",
		UserId: receipt.Meta.UserID,
	}
	cred, err := db.GetUserCredentials(ctx, ut)
	if err != nil {
		log.Printf("Core: User %s not found", receipt.Meta.UserID)
	}

	log.Printf(cred.GoogleApiKey)

	rt := types.ReceiptDTable{
		Name:        "Receipts",
		ReceiptId:   receipt.EventID,
		ReceiptData: receipt,
	}

	if err := db.AddItem(ctx, rt); err != nil {
		log.Printf("Core: Failed to add receipt %v in to the table Receipts. Error: %v", receipt.EventID, err)
		return err
	}

	// get ride details from uber

	// save into a google sheets spreadsheet

	log.Println("core: Lambda process has worked so far")
	return nil
}
