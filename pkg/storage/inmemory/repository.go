package inmemory

import (
	"github.com/bobrnor/highloadcup2018/pkg/account"
	"github.com/bobrnor/highloadcup2018/pkg/filtering"
	"github.com/pkg/errors"
)

type Storage struct {
	accounts account.Accounts
}

func New(accounts account.Accounts) *Storage {
	return &Storage{
		accounts: accounts,
	}
}

func (s *Storage) Fetch(filters []filtering.Filter, limit int) ([]filtering.Account, error) {
	return nil, errors.New("not implemented yet")
}
