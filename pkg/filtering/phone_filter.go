package filtering

type phoneFilter struct {
	Field     string
	Operation string
	Value     string
}

func makePhoneFilter(field, operation, value string) (Filter, error) {
	return phoneFilter{
		Field:     field,
		Operation: operation,
		Value:     value,
	}, nil
}
