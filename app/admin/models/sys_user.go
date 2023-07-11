package models

import (
	"go-admin/common/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	UserId              int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	ParentUserId        int    `json:"parentUserId" gorm:"size:20;comment:上一代id"`
	Username            string `json:"username" gorm:"size:64;comment:用户名"`
	Password            string `json:"-" gorm:"size:128;comment:密码"`
	NickName            string `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone               string `json:"phone" gorm:"size:11;comment:手机号"`
	RoleId              int    `json:"roleId" gorm:"size:20;comment:角色ID"`
	Salt                string `json:"-" gorm:"size:255;comment:加盐"`
	Avatar              string `json:"avatar" gorm:"size:255;comment:头像"`
	Sex                 string `json:"sex" gorm:"size:255;comment:性别"`
	Birthday            int64  `json:"birthday" gorm:"comment:出生日期，秒级时间戳"`
	Address             string `json:"address" gorm:"comment:常驻地址"`
	Email               string `json:"email" gorm:"size:128;comment:邮箱"`
	Remark              string `json:"remark" gorm:"comment:备注"`
	Status              string `json:"status" gorm:"size:4;comment:状态"`
	DigitalLifeFileList string `json:"digital_life"  gorm:"comment:数字生命;column:digital_life"`
	models.ControlBy
	models.ModelTime
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (e *SysUser) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysUser) GetId() interface{} {
	return e.UserId
}

// 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *SysUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if e.Password != "" {
		err = e.Encrypt()
	}
	return err
}

func (e *SysUser) AfterFind(_ *gorm.DB) error {
	return nil
}
