package types

import (
	"encoding/json"
	"fmt"
)


type RuleType string

const (
	RuleAllow RuleType = "allow"
	RuleDeny RuleType = "deny"
)

type Rule struct {
	Type   RuleType `json:"type"`
	Target Cidr   `json:"target"`
}

func (r RuleType) ToRawMessage() json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`"%s"`, r))
}

// func (r RuleType) Validate() error {
// 	if r.Target {
		
// 	}
// }
