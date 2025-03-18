package flow

type FlowOutput struct {
	Data     map[string]interface{} `json:"data"`
	Function string                 `json:"function"`
}

type FlowInput struct {
	Args     map[string]interface{} `json:"args"`
	Children map[string]FlowOutput  `json:"children"`
}
