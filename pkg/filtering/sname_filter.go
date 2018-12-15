package filtering

type snameFilter struct {
	Field     string
	Operation string
	Value     string
}

func makeSnameFilter(field, operation, value string) (Filter, error) {
	return snameFilter{
		Field:     field,
		Operation: operation,
		Value:     value,
	}, nil
}
