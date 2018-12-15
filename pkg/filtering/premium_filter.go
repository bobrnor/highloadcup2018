package filtering

import (
	"time"

	"github.com/bobrnor/highloadcup2018/pkg/account"
)

type premiumFilter struct {
	operation string
	value     string
}

func makePremiumFilter(operation, value string) (Filter, error) {
	return premiumFilter{
		operation: operation,
		value:     value,
	}, nil
}

func (f premiumFilter) Test(account account.Account) error {
	switch f.operation {
	case "now":
		if account.Premium != nil && account.Premium.Start >= time.Now().Unix() && account.Premium.Finish < time.Now().Unix() {
			return nil
		}
	case "null":
		if f.value == "0" && account.Premium != nil {
			return nil
		} else if f.value == "1" && account.Premium == nil {
			return nil
		}
	}

	return ErrTestFailed
}

func (f premiumFilter) Fill(src account.Account, dst map[string]interface{}) {
	dst["premium"] = src.Premium
}
