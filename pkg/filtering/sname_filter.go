package filtering

import (
	"strings"

	"github.com/bobrnor/highloadcup2018/pkg/account"
)

type snameFilter struct {
	operation string
	value     string
}

func makeSnameFilter(operation, value string) (Filter, error) {
	return snameFilter{
		operation: operation,
		value:     value,
	}, nil
}

func (f snameFilter) Test(account account.Account) error {
	switch f.operation {
	case "eq":
		if account.Sname != nil && strings.EqualFold(*account.Sname, f.value) {
			return nil
		}
	case "starts":
		if account.Sname != nil && strings.HasPrefix(*account.Sname, f.value) {
			return nil
		}
	case "null":
		if f.value == "0" && account.Sname != nil {
			return nil
		} else if f.value == "1" && account.Sname == nil {
			return nil
		}
	}

	return ErrTestFailed
}

func (f snameFilter) Fill(src account.Account, dst map[string]interface{}) {
	if src.Sname != nil {
		dst["sname"] = src.Sname
	}
}
