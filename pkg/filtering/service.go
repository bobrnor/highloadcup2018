package filtering

import "github.com/bobrnor/highloadcup2018/pkg/account"

type Repository interface {
	Fetch(filters []Filter, limit int) ([]account.Account, error)
}

type Service interface {
	Fetch(filters []Filter, limit int) ([]account.Account, error)
}

type service struct {
	repo Repository
}

func New(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Fetch(filters []Filter, limit int) ([]account.Account, error) {
	return s.repo.Fetch(filters, limit)
}
