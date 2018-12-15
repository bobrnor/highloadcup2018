package filtering

import (
	"strconv"

	"github.com/pkg/errors"
)

type birthFilter struct {
	Field     string
	Operation string
	Value     int64
}

func makeBirthFilter(field, operation, value string) (Filter, error) {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return birthFilter{
		Field:     field,
		Operation: operation,
		Value:     intValue,
	}, nil
}
