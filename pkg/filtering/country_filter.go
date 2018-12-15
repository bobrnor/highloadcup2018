package filtering

import (
	"strings"

	"github.com/bobrnor/highloadcup2018/pkg/account"
)

type countryFilter struct {
	operation string
	value     string
}

func makeCountryFilter(operation, value string) (Filter, error) {
	return countryFilter{
		operation: operation,
		value:     value,
	}, nil
}

func (f countryFilter) Test(account account.Account) error {
	switch f.operation {
	case "eq":
		if account.Country != nil && strings.EqualFold(*account.Country, f.value) {
			return nil
		}
	case "null":
		if f.value == "0" && account.Country != nil {
			return nil
		} else if f.value == "1" && account.Country == nil {
			return nil
		}
	}

	return ErrTestFailed
}
