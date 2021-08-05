package entity

import "time"

type AdminUser struct {
	ID        string     `gorm:"column:id;primary_key;size:36;"`
	UserName  string     `gorm:"column:user_name;size:64;index;default:'';not null;"` // 用户名
	RealName  string     `gorm:"column:real_name;size:64;index;default:'';"` // 真实姓名
	Password  string     `gorm:"column:password;size:40;default:'';not null;"`        // 密码(sha1(md5(明文))加密)
	Status    int        `gorm:"column:status;index;default:0;not null;"`             // 状态(1:启用 2:停用)
	LoginTime time.Time  `gorm:"column:login_time;"`	// 登录时间
	LoginIp   string     `gorm:"column:login_ip;size:15"`
	Creator   string     `gorm:"column:creator;size:36;"`                             // 创建者
	CreatedAt time.Time  `gorm:"column:created_at;index;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;index;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;"`
}