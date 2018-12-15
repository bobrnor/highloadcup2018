package filtering

import (
	"strings"

	"github.com/bobrnor/highloadcup2018/pkg/account"
)

type statusFilter struct {
	operation string
	value     string
}

func makeStatusFilter(operation, value string) (Filter, error) {
	return statusFilter{
		operation: operation,
		value:     value,
	}, nil
}

func (f statusFilter) Test(account account.Account) error {
	switch f.operation {
	case "eq":
		if strings.EqualFold(account.Status, f.value) {
			return nil
		}
	case "neq":
		if !strings.EqualFold(account.Status, f.value) {
			return nil
		}
	}

	return ErrTestFailed
}

func (f statusFilter) Fill(src account.Account, dst map[string]interface{}) {
	dst["status"] = src.Status
}
