package user

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Add 新增用户
func (u *User) Add(db *gorm.DB) error {
	return db.Create(u).Error
}

// GetAll 获取所有用户
func GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Where("is_deleted = ?", false).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID 根据 ID 获取用户
func GetByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.First(&user, "id = ? AND is_deleted = ?", id, false).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// Update 更新用户信息
func (u *User) Update(db *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return db.Save(u).Error
}

// GetByUsername 根据用户名获取用户
func GetByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	err := db.First(&user, "username = ? AND is_deleted = ?", username, false).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}
