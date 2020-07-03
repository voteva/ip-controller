package store

import (
	"context"
	"database/sql"
	"errors"
	keys "git.ozon.dev/tvoteva/22_ide/homework/internal/app/constants"
	"git.ozon.dev/tvoteva/22_ide/homework/internal/app/entity"
)

type ipRepository struct {
	store *store
}

const (
	SQL_CREATE = `INSERT INTO public.ip_access (ip, datetime_first, datetime_last) 
					VALUES ($1, $2, $2)
					ON CONFLICT (ip) DO UPDATE 
					SET datetime_last = $2`

	SQL_SELECT = `SELECT ip, datetime_first, datetime_last 
					FROM public.ip_access 
					WHERE ip = $1`
)

func (r *ipRepository) Update(ctx context.Context) {
	r.store.db.QueryRowContext(
		ctx,
		SQL_CREATE,
		ctx.Value(keys.IP),
		ctx.Value(keys.Datetime),
	)
}

func (r *ipRepository) FindByIp(ctx context.Context) (*entity.IpAccessEntity, error) {
	ipEntity := &entity.IpAccessEntity{}

	if err := r.store.db.GetContext(ctx, ipEntity, SQL_SELECT, ctx.Value(keys.IP)); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return ipEntity, nil
}
