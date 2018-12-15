package filtering

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type likesFilter struct {
	Field     string
	Operation string
	Value     []int64
}

func makeLikesFilter(field, operation, value string) (Filter, error) {
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
		Field:     field,
		Operation: operation,
		Value:     intValues,
	}, nil
}
