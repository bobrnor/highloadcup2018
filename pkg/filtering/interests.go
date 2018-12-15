package filtering

import (
	"strings"

	"github.com/bobrnor/highloadcup2018/pkg/account"
)

type interestsFilter struct {
	operation string
	value     []string
}

func makeInterestsFilter(operation, value string) (Filter, error) {
	values := strings.Split(value, ",")

	return interestsFilter{
		operation: operation,
		value:     values,
	}, nil
}

func (f interestsFilter) Test(account account.Account) error {
	if len(account.Interests) == 0 {
		return ErrTestFailed
	}

	switch f.operation {
	case "contains":
		for _, v := range f.value {
			found := false
			for _, i := range account.Interests {
				if strings.EqualFold(v, i) {
					found = true
					break
				}
			}

			if !found {
				return ErrTestFailed
			}
		}

		return nil
	case "any":
		for _, v := range f.value {
			for _, i := range account.Interests {
				if strings.EqualFold(v, i) {
					return nil
				}
			}
		}
	}

	return ErrTestFailed
}

func (f interestsFilter) Fill(src account.Account, dst map[string]interface{}) {
	dst["interests"] = src.Interests
}
