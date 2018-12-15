package filtering

import "strings"

type cityFilter struct {
	Field     string
	Operation string
	Value     string
}

type cityAnyFilter struct {
	Field     string
	Operation string
	Value     []string
}

func makeCityFilter(field, operation, value string) (Filter, error) {
	return cityFilter{
		Field:     field,
		Operation: operation,
		Value:     value,
	}, nil
}

func makeCityAnyFilter(field, operation, value string) (Filter, error) {
	values := strings.Split(value, ",")

	return cityAnyFilter{
		Field:     field,
		Operation: operation,
		Value:     values,
	}, nil
}
