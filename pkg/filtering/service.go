package filtering

type Repository interface {
	Fetch(filters []FieldFilter, limit int) ([]Account, error)
}

type Service interface {
	Fetch(filters []FieldFilter, limit int) ([]Account, error)
}

type service struct {
	repo Repository
}

func New(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Fetch(filters []FieldFilter, limit int) ([]Account, error) {
	return s.repo.Fetch(filters, limit)
}
