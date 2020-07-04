package service

import (
	"context"
	"github.com/voteva/ip-controller/pkg/model"
)

type IpService interface {
	HandleSetAccessTime(ctx context.Context)
	HandleGetFirstAccessTime(ctx context.Context) *model.AccessTimeResponse
	HandleCheckAccessTime(ctx context.Context, request model.IpRequest) (*model.AccessTimeResponse, error)
}
