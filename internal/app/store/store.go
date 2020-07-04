package store

import (
	"github.com/jmoiron/sqlx"
)

type store struct {
	db           *sqlx.DB
	ipRepository IpRepository
}

func New(db *sqlx.DB) *store {
	return &store{
		db: db,
	}
}

func (s *store) IpRepository() IpRepository {
	if s.ipRepository != nil {
		return s.ipRepository
	}
	s.ipRepository = &ipRepository{
		store: s,
	}
	return s.ipRepository
}
