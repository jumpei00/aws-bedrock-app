package common

type PromptFlowEvent struct {
	MessageVersion string         `json:"messageVersion"`
	Flow           PromptFlow     `json:"flow"`
	Node           PromptFlowNode `json:"node"`
}

type PromptFlow struct {
	FlowArn      string `json:"flowArn"`
	FlowAliasArn string `json:"flowAliasArn"`
}

type PromptFlowNode struct {
	Name       string                `json:"name"`
	NodeInputs []PromptFlowNodeInput `json:"nodeInputs"`
}

type PromptFlowNodeInput struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Expression string `json:"expression"`
	Value      string `json:"value"`
}
