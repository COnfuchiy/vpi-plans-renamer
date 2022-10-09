package components

import (
	"encoding/json"
	"os"
)

type Rules map[string]string

func NewRules(jsonPath string) (Rules, error) {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	var rules Rules
	err = json.Unmarshal(data, &rules)
	if err != nil {
		return nil, err
	}
	return rules, err
}
