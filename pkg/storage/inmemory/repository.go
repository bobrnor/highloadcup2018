package inmemory

import (
	"sort"

	"github.com/bobrnor/highloadcup2018/pkg/account"
	"github.com/bobrnor/highloadcup2018/pkg/filtering"
)

type sortableAccounts []account.Account

func (a sortableAccounts) Len() int           { return len(a) }
func (a sortableAccounts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortableAccounts) Less(i, j int) bool { return a[i].ID > a[j].ID }

type Storage struct {
	accounts sortableAccounts
}

func New(accounts account.Accounts) *Storage {
	s := Storage{
		accounts: accounts.Accounts,
	}

	sort.Sort(s.accounts)

	return &s
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
