package filtering

type premiumFilter struct {
	Field     string
	Operation string
	Value     string
}

func makePremiumFilter(field, operation, value string) (Filter, error) {
	return premiumFilter{
		Field:     field,
		Operation: operation,
		Value:     value,
	}, nil
}
