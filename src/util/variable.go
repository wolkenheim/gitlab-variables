package util

import "encoding/json"

type ChangeVariable struct {
	Variable   Variable
	ChangeType ChangeType
}

type ChangeType int

const (
	CREATE ChangeType = iota + 1
	UPDATE
	DELETE
)

type Variable struct {
	Type      string `json:"variable_type,omitempty"`
	Key       string `json:"key,omitempty"`
	Value     string `json:"value,omitempty"`
	Protected bool   `json:"protected,omitempty"`
	Masked    bool   `json:"masked,omitempty"`
	Scope     string `json:"environment_scope,omitempty"`
}

func NewVariable(key string, value string) *Variable {
	return &Variable{"env_var", key, value, false, false, "*"}
}

func ParseVariableJson(variableList []byte) []Variable {
	var list []Variable
	json.Unmarshal(variableList, &list)

	var keys = make(map[string]string)
	for _, v := range list {
		keys[v.Key] = v.Value
	}
	return list
}
