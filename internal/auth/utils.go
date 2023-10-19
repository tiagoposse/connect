package auth

import (
	"fmt"

	"github.com/tiagoposse/connect/internal/utils"
)

type IgnoringCheck struct {
	ignoreUsers  []string
	ignoreGroups []string
	userFilter   []string
	groupFilter  []string
}

func (ic *IgnoringCheck) CheckIgnoreUser(email string, membership []string) (bool, error) {
	if m, err := utils.StringGlobSlice(ic.ignoreUsers, email); err != nil {
		return true, fmt.Errorf("globbing ignore users: %w", err)
	} else if m {
		return true, nil
	}

	if m, err := utils.StringGlobSlice(ic.userFilter, email); err != nil {
		return true, fmt.Errorf("globbing ignore users: %w", err)
	} else if len(ic.userFilter) > 0 && !m {
		return true, nil
	}

	for _, g := range membership {
		if m, err := utils.StringGlobSlice(ic.ignoreGroups, g); err != nil {
			return true, fmt.Errorf("globbing ignore groups: %w", err)
		} else if m {
			return true, nil
		}

		if m, err := utils.StringGlobSlice(ic.groupFilter, g); err != nil {
			return true, fmt.Errorf("globbing ignore groups: %w", err)
		} else if len(ic.groupFilter) > 0 && !m {
			return true, nil
		}
	}
	return false, nil
}
