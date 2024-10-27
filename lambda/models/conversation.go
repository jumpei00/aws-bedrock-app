package models

type Conversation struct {
	ID          string `json:"id" dynamodbav:"id"`
	SessionKey  string `json:"session_key" dynamodbav:"sessionKey"`
	UserMessage string `json:"user_message" dynamodbav:"userMessage"`
	AIResponse  string `json:"ai_response" dynamodbav:"aiResponse"`
	CreatedAt   string `json:"created_at" dynamodbav:"createdAt"`
}
