package filtering

type countryFilter struct {
	Field     string
	Operation string
	Value     string
}

func makeCountryFilter(field, operation, value string) (Filter, error) {
	return countryFilter{
		Field:     field,
		Operation: operation,
		Value:     value,
	}, nil
}
