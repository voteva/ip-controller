package ip_controller

import (
	"context"
	"git.ozon.dev/tvoteva/22_ide/homework/pkg/model"
)

type IpService interface {
	HandleSetAccessTime(ctx context.Context)
	HandleGetFirstAccessTime(ctx context.Context) *model.AccessTimeResponse
	HandleCheckAccessTime(ctx context.Context, request model.IpRequest) (*model.AccessTimeResponse, error)
}
