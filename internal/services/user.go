package services

import (
	"gin_template/internal/data/user"
	"gin_template/pkg"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, username, password, email string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		pkg.Error("生成密码哈希失败", err)
		return err
	}
	user := &user.User{Username: username, Password: string(hashedPassword), Email: email}

	err = user.Add(db)
	if err != nil {
		pkg.Error("添加用户失败", err)
		return err
	}
	return nil
}

func Login(db *gorm.DB, username, password string) (*user.User, error) {

	// 查询用户
	user, err := user.GetByUsername(db, username)
	if err != nil {
		pkg.Error("用户不存在", err)
		return nil, err
	}

	return user, nil
}
