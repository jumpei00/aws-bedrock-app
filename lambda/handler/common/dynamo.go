package common

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoHandler struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoHandler(tableName string) (*DynamoHandler, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	return &DynamoHandler{client: dynamodb.NewFromConfig(cfg), tableName: tableName}, nil
}

func (d *DynamoHandler) Client() *dynamodb.Client {
	return d.client
}

func (d *DynamoHandler) TableName() *string {
	return &d.tableName
}
