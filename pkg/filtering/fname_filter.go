package filtering

import "strings"

type fnameFilter struct {
	Field     string
	Operation string
	Value     string
}

type fnameAnyFilter struct {
	Field     string
	Operation string
	Value     []string
}

func makeFnameFilter(field, operation, value string) (Filter, error) {
	return fnameFilter{
		Field:     field,
		Operation: operation,
		Value:     value,
	}, nil
}

func makeFnameAnyFilter(field, operation, value string) (Filter, error) {
	values := strings.Split(value, ",")

	return fnameAnyFilter{
		Field:     field,
		Operation: operation,
		Value:     values,
	}, nil
}
