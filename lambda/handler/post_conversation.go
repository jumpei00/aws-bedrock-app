package handler

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/jumpei00/aws-bedrock-app/lambda/handler/common"
	"github.com/jumpei00/aws-bedrock-app/lambda/models"
)

type PostConversationHandler struct {
	dynamoHandler *common.DynamoHandler
}

func NewPostConversationHandler(tableName string) (*PostConversationHandler, error) {
	dynamoHandler, err := common.NewDynamoHandler(tableName)
	if err != nil {
		return nil, err
	}
	return &PostConversationHandler{dynamoHandler: dynamoHandler}, nil
}

func (h *PostConversationHandler) Handle(ctx context.Context, req common.PromptFlowEvent) error {
	// 1番目のnodeInputsにsessionKeyが存在する前提
	// 2番目のnodeInputsにuserMessageが存在する前提
	// 3番目のnodeInputsにaiResponseが存在する前提

	nodeInputs := req.Node.NodeInputs
	conversation := models.Conversation{
		ID:          uuid.New().String(),
		SessionKey:  nodeInputs[0].Value,
		UserMessage: nodeInputs[1].Value,
		AIResponse:  nodeInputs[2].Value,
		CreatedAt:   time.Now().UTC().Format(time.RFC3339),
	}

	item, err := attributevalue.MarshalMap(conversation)
	if err != nil {
		return err
	}

	_, err = h.dynamoHandler.Client().PutItem(ctx, &dynamodb.PutItemInput{
		TableName: h.dynamoHandler.TableName(),
		Item:      item,
	})

	return err
}
