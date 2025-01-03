package handler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jumpei00/aws-bedrock-app/lambda/handler/common"
	"github.com/jumpei00/aws-bedrock-app/lambda/models"
)

type GetConversationHandler struct {
	dynamoHandler *common.DynamoHandler
}

func NewGetConversationHandler(tableName string) (*GetConversationHandler, error) {
	dynamoHandler, err := common.NewDynamoHandler(tableName)
	if err != nil {
		return nil, err
	}
	return &GetConversationHandler{dynamoHandler: dynamoHandler}, nil
}

func (h *GetConversationHandler) Handle(ctx context.Context, req common.PromptFlowEvent) ([]models.Conversation, error) {
	// 1番目のnodeInputsにsessionKeyが存在する前提
	sessionKey := req.Node.NodeInputs[0].Value

	expr, err := expression.NewBuilder().WithKeyCondition(expression.Key("sessionKey").Equal(expression.Value(sessionKey))).Build()
	if err != nil {
		return nil, err
	}

	result, err := h.dynamoHandler.Client().Query(ctx, &dynamodb.QueryInput{
		TableName:                 h.dynamoHandler.TableName(),
		IndexName:                 aws.String("SessionIndex"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if err != nil {
		return nil, err
	}

	var conversations []models.Conversation
	if err := attributevalue.UnmarshalListOfMaps(result.Items, &conversations); err != nil {
		return nil, err
	}

	return conversations, nil
}
