package filtering

type emailFilter struct {
	Field     string
	Operation string
	Value     string
}

func makeEmailFilter(field, operation, value string) (Filter, error) {
	return emailFilter{
		Field:     field,
		Operation: operation,
		Value:     value,
	}, nil
}
