package utils

import (
	"fmt"
	"os/exec"

	"github.com/tiagoposse/connect/internal/types"
)

type IpTablesOperationType string
const (
	IpTablesAdd IpTablesOperationType = "add"
	IpTablesDrop IpTablesOperationType = "drop"
)

type IpTablesOperation struct {
	Operation IpTablesOperationType
	Origin types.Cidr
	Destination types.Cidr
	Action types.RuleType
}

func ExecRulesWithRollback(rules []IpTablesOperation) (retErr error) {
	rollbackFrom := -1
	for index, rule := range rules {
		if err := ExecIPTables(rule); err != nil {
			rollbackFrom = index
			retErr = err
			break
		}
	}

	if rollbackFrom > -1 {
		for i :=0; i <= rollbackFrom; i++ {
			rule := rules[i]
			if rule.Operation == IpTablesAdd {
				rule.Operation = IpTablesDrop
			} else {
				rule.Operation = IpTablesAdd
			}

			ExecIPTables(rule)
		}
	}

	return
}

func ExecIPTables(rule IpTablesOperation) error {
	var trafficAction, tablesAction string
	if rule.Action == types.RuleAllow {
		trafficAction = "ACCEPT"
	} else {
		trafficAction = "DROP"
	}

	if rule.Operation == IpTablesAdd {
		tablesAction = "A"
	} else {
		tablesAction = "D"
	}

	fmt.Println([]string{"iptables", "-" + tablesAction, "FORWARD", "-s", rule.Origin.String(), "-d", rule.Destination.String(), "-j", trafficAction})
	cmd := exec.Command("iptables", "-" + tablesAction, "FORWARD", "-s", rule.Origin.String(), "-d", rule.Destination.String(), "-j", trafficAction)
	return cmd.Run()
}
