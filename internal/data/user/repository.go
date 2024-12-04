package user

import (
	"context"
	"errors"
	"yug_server/internal/data/user/model"
	"yug_server/internal/dto"
	"yug_server/internal/repo"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userRepo struct {
	db     *gorm.DB
	rds    *redis.Client
	logger *zap.Logger
}

func NewUserRepo(db *gorm.DB, rds *redis.Client, logger *zap.Logger) repo.UserRepo {
	return &userRepo{db: db, rds: rds, logger: logger}
}

func (u *userRepo) Register(ctx context.Context, dto *dto.RegisterRequest) error {
	user := &model.User{
		Username: dto.Username,
		Password: dto.Password,
		Email:    dto.Email,
		Phone:    dto.Phone,
	}

	// 检查用户是否存在
	if err := u.db.Where("username = ?", dto.Username).First(user).Error; err != nil {
		u.logger.Error("用户已存在", zap.Error(err))
		return err
	}

	err := u.db.Create(user).Error
	if err != nil {
		u.logger.Error("注册失败", zap.Error(err))
		return err
	}

	return nil
}

func (u *userRepo) Login(ctx context.Context, username, password string) (*model.User, error) {
	var user model.User
	err := u.db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Error("登录失败，用户不存在或密码错误", zap.Error(err))
		return nil, err
	}
	return &user, err
}

func (u *userRepo) QueryUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Error("查询用户失败，用户不存在", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) QueryUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Error("查询用户失败，用户不存在", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) QueryUserByPhone(ctx context.Context, phone string) (*model.User, error) {
	var user model.User
	err := u.db.Where("phone = ?", phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Error("查询用户失败，用户不存在", zap.Error(err))
		return nil, err
	}
	return &user, nil
}
