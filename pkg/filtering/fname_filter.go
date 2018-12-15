package filtering

import (
	"strings"

	"github.com/bobrnor/highloadcup2018/pkg/account"
)

type fnameFilter struct {
	operation string
	value     string
}

func makeFnameFilter(operation, value string) (Filter, error) {
	return fnameFilter{
		operation: operation,
		value:     value,
	}, nil
}

func (f fnameFilter) Test(account account.Account) error {
	switch f.operation {
	case "eq":
		if account.Fname != nil && strings.EqualFold(*account.Fname, f.value) {
			return nil
		}
	case "null":
		if f.value == "0" && account.Fname != nil {
			return nil
		} else if f.value == "1" && account.Fname == nil {
			return nil
		}
	}

	return ErrTestFailed
}

func (f fnameFilter) Fill(src account.Account, dst map[string]interface{}) {
	dst["fname"] = src.Fname
}

type fnameAnyFilter struct {
	operation string
	value     []string
}

func makeFnameAnyFilter(operation, value string) (Filter, error) {
	values := strings.Split(value, ",")

	return fnameAnyFilter{
		operation: operation,
		value:     values,
	}, nil
}

func (f fnameAnyFilter) Test(account account.Account) error {
	if account.Fname == nil {
		return ErrTestFailed
	}

	for _, v := range f.value {
		if strings.EqualFold(*account.Fname, v) {
			return nil
		}
	}

	return ErrTestFailed
}

func (f fnameAnyFilter) Fill(src account.Account, dst map[string]interface{}) {
	if src.Fname != nil {
		dst["fname"] = src.Fname
	}
}
