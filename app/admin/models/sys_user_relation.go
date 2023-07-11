package models

import (
     

	"go-admin/common/models"

)

type SysUserRelation struct {
    models.Model
    
    Title string `json:"title" gorm:"type:varchar(255);comment:关系名称"` 
    models.ModelTime
    models.ControlBy
}

func (SysUserRelation) TableName() string {
    return "sys_user_relation"
}

func (e *SysUserRelation) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysUserRelation) GetId() interface{} {
	return e.Id
}