package filtering

import (
	"strings"

	"github.com/bobrnor/highloadcup2018/pkg/account"
)

type cityFilter struct {
	operation string
	value     string
}

func makeCityFilter(operation, value string) (Filter, error) {
	return cityFilter{
		operation: operation,
		value:     value,
	}, nil
}

func (f cityFilter) Test(account account.Account) error {
	switch f.operation {
	case "eq":
		if account.City != nil && strings.EqualFold(*account.City, f.value) {
			return nil
		}
	case "null":
		if f.value == "0" && account.City != nil {
			return nil
		} else if f.value == "1" && account.City == nil {
			return nil
		}
	}

	return ErrTestFailed
}

func (f cityFilter) Fill(src account.Account, dst map[string]interface{}) {
	dst["city"] = src.City
}

type cityAnyFilter struct {
	operation string
	value     []string
}

func makeCityAnyFilter(operation, value string) (Filter, error) {
	values := strings.Split(value, ",")

	return cityAnyFilter{
		operation: operation,
		value:     values,
	}, nil
}

func (f cityAnyFilter) Test(account account.Account) error {
	if account.City == nil {
		return ErrTestFailed
	}

	switch f.operation {
	case "any":
		for _, city := range f.value {
			if strings.EqualFold(*account.City, city) {
				return nil
			}
		}
	}

	return ErrTestFailed
}

func (f cityAnyFilter) Fill(src account.Account, dst map[string]interface{}) {
	if src.City != nil {
		dst["city"] = src.City
	}
}
