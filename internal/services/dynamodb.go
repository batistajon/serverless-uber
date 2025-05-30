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
	dtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

func (ds *DynamoDBService) AddItem(ctx context.Context, table types.Table) error {
	item, err := attributevalue.MarshalMap(table)
	if err != nil {
		log.Printf("Couldn't unmarshal. Here's why: %v\n", err)
	}

	_, err = ds.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table.GetName()), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return err
}

func (ds *DynamoDBService) GetItem(ctx context.Context, table types.Table) error {
	return nil
}

func (ds *DynamoDBService) GetUserCredentials(ctx context.Context, table types.Table) (types.Credentials, error) {
	ut, err := attributevalue.Marshal(table)
	if err != nil {
		log.Printf("Couldn't Marshal UserId Here's why: %v\n", err)
		return types.Credentials{}, err
	}

	res, err := ds.Client.GetItem(ctx, &dynamodb.GetItemInput{
		Key: map[string]dtypes.AttributeValue{"UserId": uid}, TableName: aws.String(table.GetName()),
	})
	if err != nil {
		log.Printf("Unable to get user. Here's why: %v\n", err)
		return types.Credentials{}, err
	}

	if err := attributevalue.UnmarshalMap(res.Item, &u); err != nil {
		log.Printf("Unable to Unmarshal response into User object for user %v Here's why: %v\n", uberUserId, err)
		return types.Credentials{}, err
	}
	return ut.GetCredentials(), err
}
