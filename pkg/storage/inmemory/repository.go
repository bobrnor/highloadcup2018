package inmemory

import (
	"github.com/bobrnor/highloadcup2018/pkg/account"
	"github.com/bobrnor/highloadcup2018/pkg/filtering"
)

type Storage struct {
	accounts []account.Account
}

func New(accounts account.Accounts) *Storage {
	return &Storage{
		accounts: accounts.Accounts,
	}
}

func (s *Storage) Fetch(filters []filtering.Filter, limit int) ([]account.Account, error) {
	var result []account.Account

	for _, acc := range s.accounts {
		passed := true
		for _, f := range filters {
			if err := f.Test(acc); err != nil {
				passed = false
				break
			}
		}

		if passed {
			result = append(result, acc)
		}

		if len(result) >= limit {
			break
		}
	}

	return result, nil
}
