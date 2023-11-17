package model

type BasicQueryPayload struct {
	Data      interface{}  `json:"data"`
	Option    OPTION_QUERY `json:"option"`
	ModelType string       `json:"modelType"`
}

type QueryableModel interface {
	TableName() string
}
