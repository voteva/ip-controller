package ip_controller

import (
	"context"
	"fmt"
	keys "github.com/voteva/ip-controller/internal/app/constants"
	"github.com/voteva/ip-controller/internal/app/store"
	"github.com/voteva/ip-controller/internal/app/store/db"
	"github.com/voteva/ip-controller/pkg/model"
	"time"
)

type ipService struct {
	repository store.IpRepository
}

func New() *ipService {
	return &ipService{
		repository: store.New(db.DB).IpRepository(),
	}
}

func (s *ipService) HandleSetAccessTime(ctx context.Context) {
	newCtx := context.WithValue(ctx, keys.Datetime, time.Now())
	s.repository.Update(newCtx)
}

func (s *ipService) HandleGetFirstAccessTime(ctx context.Context) *model.AccessTimeResponse {
	s.HandleSetAccessTime(ctx)
	ipEntity, _ := s.repository.FindByIp(ctx)
	result := &model.AccessTimeResponse{
		Datetime: ipEntity.DatetimeFirst.String(),
	}
	return result
}

func (s *ipService) HandleCheckAccessTime(ctx context.Context, request model.IpRequest) (*model.AccessTimeResponse, error) {
	newCtx := context.WithValue(ctx, keys.IP, request.IP)

	ipEntity, err := s.repository.FindByIp(newCtx)
	if err != nil {
		return nil, fmt.Errorf("no info about this ip = %s", request.IP)
	}

	result := &model.AccessTimeResponse{
		Datetime: ipEntity.DatetimeLast.String(),
	}

	s.HandleSetAccessTime(newCtx)

	return result, nil
}
