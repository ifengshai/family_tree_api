package dto

import (
     

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysUserRelationGetPageReq struct {
	dto.Pagination     `search:"-"`
    Title string `form:"title"  search:"type:exact;column:title;table:sys_user_relation" comment:"关系名称"`
    SysUserRelationOrder
}

type SysUserRelationOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:sys_user_relation"`
    Title string `form:"titleOrder"  search:"type:order;column:title;table:sys_user_relation"`
    
}

func (m *SysUserRelationGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysUserRelationInsertReq struct {
    Id int `json:"-" comment:"主键"` // 主键
    Title string `json:"title" comment:"关系名称"`
    common.ControlBy
}

func (s *SysUserRelationInsertReq) Generate(model *models.SysUserRelation)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
}

func (s *SysUserRelationInsertReq) GetId() interface{} {
	return s.Id
}

type SysUserRelationUpdateReq struct {
    Id int `uri:"id" comment:"主键"` // 主键
    Title string `json:"title" comment:"关系名称"`
    common.ControlBy
}

func (s *SysUserRelationUpdateReq) Generate(model *models.SysUserRelation)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
}

func (s *SysUserRelationUpdateReq) GetId() interface{} {
	return s.Id
}

// SysUserRelationGetReq 功能获取请求参数
type SysUserRelationGetReq struct {
     Id int `uri:"id"`
}
func (s *SysUserRelationGetReq) GetId() interface{} {
	return s.Id
}

// SysUserRelationDeleteReq 功能删除请求参数
type SysUserRelationDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysUserRelationDeleteReq) GetId() interface{} {
	return s.Ids
}
