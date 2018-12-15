package filtering

import "github.com/pkg/errors"

type sexFilter struct {
	Field     string
	Operation string
	Value     string
}

func makeSexFilter(field, operation, value string) (Filter, error) {
	if value != "m" && value != "f" {
		return nil, errors.New("bad filter value")
	}

	return sexFilter{
		Field:     field,
		Operation: operation,
		Value:     value,
	}, nil
}
