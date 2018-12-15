package filtering

import "strings"

type interestsFilter struct {
	Field     string
	Operation string
	Value     []string
}

func makeInterestsFilter(field, operation, value string) (Filter, error) {
	values := strings.Split(value, ",")

	return interestsFilter{
		Field:     field,
		Operation: operation,
		Value:     values,
	}, nil
}
