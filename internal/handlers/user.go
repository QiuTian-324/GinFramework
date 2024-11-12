package handlers

import (
	"gin_template/internal/dto"
	"gin_template/internal/libs"
	"gin_template/internal/services"
	"gin_template/pkg"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {

	req := new(dto.RegisterRequest)
	db := ctx.MustGet("db").(*gorm.DB)
	if err := ctx.ShouldBindJSON(req); err != nil {
		libs.BadRequestResponse(ctx, "请求数据无效")
		pkg.Error("请求数据无效", err)
		return
	}

	if err := services.RegisterUser(db, req.Username, req.Password, req.Email); err != nil {
		libs.InternalServerErrorResponse(ctx, "注册失败")
		pkg.Error("注册失败", err)
		return
	}

	libs.SuccessResponse(ctx, "注册成功", nil)
}

func Login(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	req := new(dto.LoginRequest)
	res := new(dto.LoginResponse)

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		libs.BadRequestResponse(ctx, "请求数据无效")
		pkg.Error("请求数据无效", err)
		return
	}

	if libs.ValidateEmpty(req.Username) {
		libs.BadRequestResponse(ctx, "用户名不能为空")
		pkg.Error("用户名不能为空", nil)
		return
	}

	if libs.ValidateEmpty(req.Password) {
		libs.BadRequestResponse(ctx, "密码不能为空")
		pkg.Error("密码不能为空", nil)
		return
	}

	// 查询用户
	userInfo, err := services.Login(db, req.Username, req.Password)
	if err != nil {
		libs.InternalServerErrorResponse(ctx, "用户不存在")
		pkg.Error("用户不存在", err)
		return
	}

	// 生成token
	token, err := libs.GenToken(userInfo.ID, userInfo.Username)
	if err != nil {
		libs.InternalServerErrorResponse(ctx, "登陆失败")
		pkg.Error("生成token失败", err)
		return
	}

	// 生成token
	res.Token = token
	libs.SuccessResponse(ctx, "登录成功", res)
}
