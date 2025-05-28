package services

import (
	"context"
	"fmt"
	"goLambda/internal/types"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBService struct {
	Client *dynamodb.Client
}

func NewDynamoDBService(ctx context.Context) (*DynamoDBService, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %w", err)
	}

	return &DynamoDBService{
		Client: dynamodb.NewFromConfig(cfg),
	}, nil
}

func (ds DynamoDBService) Add(ctx context.Context, table types.Table) error {
	item, err := attributevalue.MarshalMap(table)
	if err != nil {
		panic(err)
	}

	_, err = ds.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table.GetName()), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return err
}
