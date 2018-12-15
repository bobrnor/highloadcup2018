package filtering

import (
	"strconv"
	"strings"

	"github.com/bobrnor/highloadcup2018/pkg/account"

	"github.com/pkg/errors"
)

type likesFilter struct {
	operation string
	value     []int64
}

func makeLikesFilter(operation, value string) (Filter, error) {
	var intValues []int64

	vv := strings.Split(value, ",")
	for _, v := range vv {
		intV, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		intValues = append(intValues, intV)
	}

	return likesFilter{
		operation: operation,
		value:     intValues,
	}, nil
}

func (f likesFilter) Test(account account.Account) error {
	if len(account.Likes) == 0 {
		return ErrTestFailed
	}

	for _, v := range f.value {
		found := false
		for _, l := range account.Likes {
			if v == l.ID {
				found = true
				break
			}
		}

		if !found {
			return ErrTestFailed
		}
	}

	return nil
}

func (f likesFilter) Fill(src account.Account, dst map[string]interface{}) {
	if len(src.Likes) > 0 {
		dst["likes"] = src.Likes
	}
}
