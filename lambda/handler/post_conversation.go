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

type PostConversationRequest struct {
	SessionKey  string `json:"session_key"`
	UserMessage string `json:"user_message"`
	AIResponse  string `json:"ai_response"`
}

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

func (h *PostConversationHandler) Handle(ctx context.Context, req PostConversationRequest) error {
	conversation := models.Conversation{
		ID:          uuid.New().String(),
		SessionKey:  req.SessionKey,
		UserMessage: req.UserMessage,
		AIResponse:  req.AIResponse,
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
