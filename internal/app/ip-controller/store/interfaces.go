package store

import (
	"context"
	"github.com/voteva/ip-controller/internal/app/ip-controller/entity"
)

type Store interface {
	IpRepository() IpRepository
}

type IpRepository interface {
	Update(ctx context.Context)
	FindByIp(ctx context.Context) (*entity.IpAccessEntity, error)
}
