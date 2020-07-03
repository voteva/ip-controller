package store

import (
	"context"
	"git.ozon.dev/tvoteva/22_ide/homework/internal/app/entity"
)

type Store interface {
	IpRepository() IpRepository
}

type IpRepository interface {
	Update(ctx context.Context)
	FindByIp(ctx context.Context) (*entity.IpAccessEntity, error)
}
