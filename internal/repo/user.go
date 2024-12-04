package repo

import (
	"context"
	"yug_server/internal/data/user/model"
	"yug_server/internal/dto"
)

type UserRepo interface {
	Register(ctx context.Context, dto *dto.RegisterRequest) error
	Login(ctx context.Context, username, password string) (*model.User, error)
	// QueryUser(ctx context.Context, username, email, phone string) (*model.User, error)
	QueryUserByUsername(ctx context.Context, username string) (*model.User, error)
	QueryUserByEmail(ctx context.Context, email string) (*model.User, error)
	QueryUserByPhone(ctx context.Context, phone string) (*model.User, error)
}
