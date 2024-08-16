package user

import "time"

// AuthUser 用户表
type AuthUser struct {
	ID       uint   `json:"id" gorm:"comment:主键id;column:id;primarykey"`
	Username string `json:"username" gorm:"comment:用户登录名;column:username;unique"`
	Password string `json:"password" gorm:"comment:用户登录密码;column:password"`
	// 昵称
	NickName string `json:"nick_name" gorm:"comment:用户昵称;column:nick_name"`
	// 用户类型
	AuthorityId int64 `json:"authority_id" gorm:"comment:密码权限id;column:authority_id"`
	// 路由权限id
	RouterId int64 `json:"router_id" gorm:"comment:路由权限id;column:router_id"`
	// 角色类型
	Authority string `json:"authority" gorm:"comment:用户角色类型;column:authority"`
	// 是否启用
	IsUse     bool      `json:"is_use" gorm:"default:1;comment:是否启用;column:is_use"`
	Remake    string    `json:"remake" gorm:"comment:备注;column:remake"`
	AllowIp   string    `json:"allow_ip" gorm:"comment:允许的ip地址来源;column:allow_ip"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 设置表名
func (user AuthUser) TableName() string {
	return "auth_user"
}
