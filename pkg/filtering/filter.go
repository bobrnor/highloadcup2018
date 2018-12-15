package filtering

import (
	"strings"

	"github.com/pkg/errors"
)

type filterMakeFunc func(string, string, string) (Filter, error)

type Filter interface {
}

type noFilter struct{}

var filterMakeFuncs = map[string]map[string]filterMakeFunc{
	"sex": {
		"eq": makeSexFilter,
	},
	"email": {
		"domain": makeEmailFilter,
		"lt":     makeEmailFilter,
		"gt":     makeEmailFilter,
	},
	"status": {
		"eq":  makeStatusFilter,
		"neq": makeStatusFilter,
	},
	"fname": {
		"eq":   makeFnameFilter,
		"any":  makeFnameAnyFilter,
		"null": makeFnameFilter,
	},
	"sname": {
		"eq":     makeSnameFilter,
		"starts": makeSnameFilter,
		"null":   makeSnameFilter,
	},
	"phone": {
		"code": makePhoneFilter,
		"null": makePhoneFilter,
	},
	"country": {
		"eq":   makeCountryFilter,
		"null": makeCountryFilter,
	},
	"city": {
		"eq":   makeCityFilter,
		"any":  makeCityAnyFilter,
		"null": makeCityFilter,
	},
	"birth": {
		"lt":   makeBirthFilter,
		"gt":   makeBirthFilter,
		"year": makeBirthFilter,
	},
	"interests": {
		"contains": makeInterestsFilter,
		"any":      makeInterestsFilter,
	},
	"likes": {
		"contains": makeLikesFilter,
	},
	"premium": {
		"now":  makePremiumFilter,
		"null": makePremiumFilter,
	},
}

func Make(key, value string) (Filter, error) {
	field, operation, err := split(key)
	if err != nil {
		return noFilter{}, err
	}

	f, ok := filterMakeFuncs[field][operation]
	if !ok {
		return noFilter{}, errors.New("filter not found")
	}

	return f(field, operation, value)
}

func split(key string) (string, string, error) {
	parts := strings.Split(key, "_")
	if len(parts) != 2 {
		return "", "", errors.New("bad filter key")
	}

	return parts[0], parts[1], nil
}
