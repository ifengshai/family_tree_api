package models

import (
	"go-admin/common/models"
)

type SysUserRelationData struct {
	models.Model

	UserId          int             `json:"user_id" gorm:"type:int;comment:用户表id"`
	UserRelationId  int             `json:"user_relation_id" gorm:"type:int;comment:关系表id"`
	RelationUserId  int             `json:"relation_user_id" gorm:"type:int;comment:关联用户的id"`
	SysUserRelation SysUserRelation `gorm:"foreignKey:Id;references:UserRelationId""`
	models.ModelTime
}

func (SysUserRelationData) TableName() string {
	return "sys_user_relation_data"
}

func (e *SysUserRelationData) GetId() interface{} {
	return e.Id
}
