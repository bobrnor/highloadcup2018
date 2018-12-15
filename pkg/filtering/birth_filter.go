package filtering

import (
	"strconv"
	"time"

	"github.com/bobrnor/highloadcup2018/pkg/account"
	"github.com/pkg/errors"
)

type birthFilter struct {
	operation string
	value     int64
}

func makeBirthFilter(operation, value string) (Filter, error) {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return birthFilter{
		operation: operation,
		value:     intValue,
	}, nil
}

func (f birthFilter) Test(account account.Account) error {
	switch f.operation {
	case "lt":
		if account.Birth < f.value {
			return nil
		}
	case "gt":
		if account.Birth > f.value {
			return nil
		}
	case "year":
		if time.Unix(account.Birth, 0).Year() == int(f.value) {
			return nil
		}
	}

	return ErrTestFailed
}

func (f birthFilter) Fill(src account.Account, dst map[string]interface{}) {
	dst["birth"] = src.Birth
}
