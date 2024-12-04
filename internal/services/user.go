package services

import (
	"context"
	"fmt"

	"yug_server/internal/data/user/model"
	"yug_server/internal/dto"
	"yug_server/internal/repo"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type UserUseCase struct {
	repo   repo.UserRepo
	rds    *redis.Client
	logger *zap.Logger
}

func NewUserUseCase(repo repo.UserRepo, rds *redis.Client, logger *zap.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, rds: rds, logger: logger}
}

func (uc *UserUseCase) Register(ctx context.Context, dto *dto.RegisterRequest) error {
	err := uc.repo.Register(ctx, dto)
	if err != nil {
		uc.logger.Error("注册失败", zap.Error(err))
		return err
	}
	return nil
}

func (uc *UserUseCase) Login(ctx context.Context, username, password string) (*model.User, error) {

	// 查询用户
	userInfo, err := uc.repo.Login(ctx, username, password)
	if err != nil {
		uc.logger.Error("用户不存在", zap.Error(err))
		return nil, err
	}

	return userInfo, nil
}

// 按用户名查询用户
func (uc *UserUseCase) QueryUserByUsername(ctx context.Context, username string) (*model.User, error) {
	userInfo, err := uc.repo.QueryUserByUsername(ctx, username)
	if err != nil {
		uc.logger.Error("按用户名查询用户失败", zap.Error(err))
		return nil, err
	}
	return userInfo, nil
}

// 按邮箱查询用户
func (uc *UserUseCase) QueryUserByEmail(ctx context.Context, email string) (*model.User, error) {
	userInfo, err := uc.repo.QueryUserByEmail(ctx, email)
	if err != nil {
		uc.logger.Error("按邮箱查询用户失败", zap.Error(err))
		return nil, err
	}
	return userInfo, nil
}

// 按手机号查询用户
func (uc *UserUseCase) QueryUserByPhone(ctx context.Context, phone string) (*model.User, error) {
	userInfo, err := uc.repo.QueryUserByPhone(ctx, phone)
	if err != nil {
		uc.logger.Error("按手机号查询用户失败", zap.Error(err))
		return nil, err
	}
	return userInfo, nil
}

type QueryStrategy interface {
	Query(ctx context.Context) (*model.User, error)
}

type QueryByUsername struct {
	username string
	uc       *UserUseCase
}

func (q *QueryByUsername) Query(ctx context.Context) (*model.User, error) {
	return q.uc.QueryUserByUsername(ctx, q.username)
}

type QueryByEmail struct {
	email string
	uc    *UserUseCase
}

func (q *QueryByEmail) Query(ctx context.Context) (*model.User, error) {
	return q.uc.QueryUserByEmail(ctx, q.email)
}

type QueryByPhone struct {
	phone string
	uc    *UserUseCase
}

func (q *QueryByPhone) Query(ctx context.Context) (*model.User, error) {
	return q.uc.QueryUserByPhone(ctx, q.phone)
}

func (uc *UserUseCase) QueryUser(ctx context.Context, username, email, phone string) (*model.User, error) {
	// 定义一个映射，将查询类型与对应的策略绑定
	strategies := map[string]QueryStrategy{
		"username": &QueryByUsername{username: username, uc: uc},
		"email":    &QueryByEmail{email: email, uc: uc},
		"phone":    &QueryByPhone{phone: phone, uc: uc},
	}

	// 遍历映射，选择第一个非空的策略
	for key, strategy := range strategies {
		if key == "username" && username != "" {
			return strategy.Query(ctx)
		}
		if key == "email" && email != "" {
			return strategy.Query(ctx)
		}
		if key == "phone" && phone != "" {
			return strategy.Query(ctx)
		}
	}

	return nil, fmt.Errorf("请输入用户名、邮箱或手机号")
}
