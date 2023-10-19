package types


type RuleType string

const (
	RuleAllow RuleType = "allow"
	RuleDeny RuleType = "deny"
)

type Rule struct {
	ID     string   `gorm:"primaryKey;autoIncrement:false"`
	Type   RuleType `json:"type"`
	Target string   `json:"target"`
}
