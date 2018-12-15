package filtering

import (
	"strings"

	"github.com/bobrnor/highloadcup2018/pkg/account"
)

type phoneFilter struct {
	operation string
	value     string
}

func makePhoneFilter(operation, value string) (Filter, error) {
	return phoneFilter{
		operation: operation,
		value:     value,
	}, nil
}

func (f phoneFilter) Test(account account.Account) error {
	switch f.operation {
	case "code":
		if account.Phone != nil && strings.Contains(*account.Phone, "("+f.value+")") {
			return nil
		}
	case "null":
		if f.value == "0" && account.Phone != nil {
			return nil
		} else if f.value == "1" && account.Phone == nil {
			return nil
		}
	}

	return ErrTestFailed
}

func (f phoneFilter) Fill(src account.Account, dst map[string]interface{}) {
	if src.Phone != nil {
		dst["phone"] = src.Phone
	}
}
