package model

type AdvanceFilterPayload struct {
	ModelType     string                 `json:"modelType"`
	Conditions    map[string]interface{} `json:"conditions"`
	StringPreLoad []string               `json:"stringPreLoad"`
	IsPreload     bool                   `json:"isPreload"`
	Page          int                    `json:"page"`
	PageSize      int                    `json:"pageSize"`
}
