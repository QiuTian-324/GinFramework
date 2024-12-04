package handlers

import (
	"fmt"
	"time"
	"yug_server/global"
	"yug_server/internal/dto"
	"yug_server/internal/libs"
	"yug_server/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type UserHandler struct {
	uc     *services.UserUseCase
	rds    *redis.Client
	logger *zap.Logger
}

func NewUserHandler(uc *services.UserUseCase, rds *redis.Client, logger *zap.Logger) *UserHandler {
	return &UserHandler{uc: uc, rds: rds, logger: logger}
}

func (h *UserHandler) Register(ctx *gin.Context) {

	req := new(dto.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		h.logger.Error("请求数据无效", zap.Error(err))
		libs.BadRequestResponse(ctx, "请求数据无效")
		return
	}

	if err := h.uc.Register(ctx, req); err != nil {
		h.logger.Error("注册失败", zap.Error(err))
		libs.InternalServerErrorResponse(ctx, "注册失败")
		return
	}

	libs.SuccessResponse(ctx, "注册成功", nil)
}

func (h *UserHandler) Login(ctx *gin.Context) {

	req := new(dto.LoginRequest)

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		h.logger.Error("请求数据无效", zap.Error(err))
		libs.BadRequestResponse(ctx, "请求数据无效")
		return
	}

	if libs.ValidateEmpty(req.Username) {
		h.logger.Error("用户名不能为空", zap.Error(err))
		libs.BadRequestResponse(ctx, "用户名不能为空")
		return
	}

	if libs.ValidateEmpty(req.Password) {
		h.logger.Error("密码不能为空", zap.Error(err))
		libs.BadRequestResponse(ctx, "密码不能为空")
		return
	}

	userInfo, err := h.uc.Login(ctx, req.Username, req.Password)
	if err != nil {
		h.logger.Error("用户不存在", zap.Error(err))
		libs.InternalServerErrorResponse(ctx, "用户不存在")
		return
	}

	token, err := libs.GenToken(uint64(userInfo.ID), userInfo.Username)
	if err != nil {
		h.logger.Error("生成token失败", zap.Error(err))
		libs.InternalServerErrorResponse(ctx, "登陆失败")
		return
	}

	redisKey := fmt.Sprintf("%s%d", global.RedisSessionKey, userInfo.ID)

	err = libs.RedisSet(ctx, redisKey, token, time.Hour*time.Duration(viper.GetInt64("redis.expires")))
	if err != nil {
		h.logger.Error("设置redis失败", zap.Error(err))
		libs.InternalServerErrorResponse(ctx, "登录失败")
		return
	}

	res := dto.LoginResponse{
		UserID:    cast.ToString(userInfo.ID),
		Username:  userInfo.Username,
		Nickname:  userInfo.Nickname,
		AvatarUrl: userInfo.AvatarUrl,
		Email:     userInfo.Email,
		Phone:     userInfo.Phone,
		Bio:       userInfo.Bio,
		Online:    userInfo.Online,
	}

	// 使用 AddExtra 添加 token 作为额外字段
	response := libs.NewResponse(libs.CodeSuccess, "登录成功", true, res, nil)
	libs.AddExtra(ctx, response, map[string]interface{}{"token": token})
}

// 查询好友

func (h *UserHandler) QueryUser(ctx *gin.Context) {
	username := ctx.Query("username")
	email := ctx.Query("email")
	phone := ctx.Query("phone")

	if libs.ValidateEmpty(username) && libs.ValidateEmpty(email) && libs.ValidateEmpty(phone) {
		h.logger.Error("请求参数不能同时为空")
		libs.BadRequestResponse(ctx, "请输入用户名、邮箱或手机号")
		return
	}

	friendInfo, err := h.uc.QueryUser(ctx, username, email, phone)
	if err != nil {
		h.logger.Error("查询用户失败", zap.Error(err))
		libs.InternalServerErrorResponse(ctx, "查询用户失败")
		return
	}

	res := dto.QueryUserResponse{
		UserID:    friendInfo.ID,
		Username:  friendInfo.Username,
		Nickname:  friendInfo.Nickname,
		Email:     friendInfo.Email,
		Phone:     friendInfo.Phone,
		AvatarUrl: friendInfo.AvatarUrl,
	}

	libs.SuccessResponse(ctx, "查询成功", res)
}

func (h *UserHandler) Logout(ctx *gin.Context) {
	userID := ctx.MustGet("id").(uint64)
	redisKey := fmt.Sprintf("%s%d", global.RedisSessionKey, userID)
	libs.RedisDelete(ctx, redisKey)
	libs.SuccessResponse(ctx, "退出成功", nil)
}
