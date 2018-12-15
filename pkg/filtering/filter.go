package filtering

import (
	"fmt"
	"strings"

	"github.com/bobrnor/highloadcup2018/pkg/account"

	"github.com/pkg/errors"
)

type filterMakeFunc func(string, string) (Filter, error)

var ErrTestFailed = fmt.Errorf("filter test failed")

type Filter interface {
	Test(account account.Account) error
	Fill(src account.Account, dst map[string]interface{})
}

type noFilter struct{}

func (f noFilter) Test(_ account.Account) error {
	return ErrTestFailed
}

func (f noFilter) Fill(src account.Account, dst map[string]interface{}) {}

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

	return f(operation, value)
}

func split(key string) (string, string, error) {
	parts := strings.Split(key, "_")
	if len(parts) != 2 {
		return "", "", errors.New("bad filter key")
	}

	return parts[0], parts[1], nil
}
