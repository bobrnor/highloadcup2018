package filtering

import (
	"strings"

	"github.com/bobrnor/highloadcup2018/pkg/account"
)

type emailFilter struct {
	operation string
	value     string
}

func makeEmailFilter(operation, value string) (Filter, error) {
	return emailFilter{
		operation: operation,
		value:     value,
	}, nil
}

func (f emailFilter) Test(account account.Account) error {
	switch f.operation {
	case "domain":
		if strings.HasSuffix(account.Email, f.value) {
			return nil
		}
	case "lt":
		if strings.Compare(account.Email, f.value) < 0 {
			return nil
		}
	case "gt":
		if strings.Compare(account.Email, f.value) > 0 {
			return nil
		}
	}

	return ErrTestFailed
}

func (f emailFilter) Fill(src account.Account, dst map[string]interface{}) {
	dst["email"] = src.Email
}
