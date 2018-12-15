package filtering

type statusFilter struct {
	Field     string
	Operation string
	Value     string
}

func makeStatusFilter(field, operation, value string) (Filter, error) {
	return statusFilter{
		Field:     field,
		Operation: operation,
		Value:     value,
	}, nil
}
