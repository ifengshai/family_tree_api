package dto

import (
	"github.com/golang-module/carbon/v2"
	jsoniter "github.com/json-iterator/go"
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysUserGetPageReq struct {
	dto.Pagination `search:"-"`
	UserId         int    `form:"userId" search:"type:exact;column:user_id;table:sys_user" comment:"用户ID"`
	Username       string `form:"username" search:"type:contains;column:username;table:sys_user" comment:"用户名"`
	NickName       string `form:"nickName" search:"type:contains;column:nick_name;table:sys_user" comment:"昵称"`
	Phone          string `form:"phone" search:"type:contains;column:phone;table:sys_user" comment:"手机号"`
	RoleId         string `form:"roleId" search:"type:exact;column:role_id;table:sys_user" comment:"角色ID"`
	Sex            string `form:"sex" search:"type:exact;column:sex;table:sys_user" comment:"性别"`
	Email          string `form:"email" search:"type:contains;column:email;table:sys_user" comment:"邮箱"`
	PostId         string `form:"postId" search:"type:exact;column:post_id;table:sys_user" comment:"岗位"`
	Status         string `form:"status" search:"type:exact;column:status;table:sys_user" comment:"状态"`
	DeptJoin       `search:"type:left;on:dept_id:dept_id;table:sys_user;join:sys_dept"`
	SysUserOrder
}

// SysUserInsertReq 插入用户信息传入数据
type SysUserInsertReq struct {
	UserId              int        `json:"userId" comment:"用户ID"` // 用户ID
	Username            string     `json:"username" comment:"用户名" vd:"len($)>0"`
	Password            string     `json:"password" comment:"密码"`
	NickName            string     `json:"nickName" comment:"昵称" vd:"len($)>0"`
	Sex                 string     `json:"sex" comment:"性别"`
	Birthday            string     `json:"birthday" comment:"出生日期"`
	Address             string     `json:"address" comment:"常驻地址"`
	Phone               string     `json:"phone" comment:"手机号"`
	Email               string     `json:"email" comment:"邮箱"`
	Status              string     `json:"status" comment:"状态" vd:"len($)>0" default:"1"`
	Remark              string     `json:"remark" comment:"备注"`
	InfoList            []InfoList `json:"infoList" comment:"用户关系"`
	DigitalLifeFileList []string   `json:"digital_life_file_list" comment:"数字生命"`
	RoleId              int        `json:"roleId" comment:"角色ID"`
	common.ControlBy
}

// SysUserUpdateReq 更新用户参数
type SysUserUpdateReq struct {
	UserId              int        `json:"userId" comment:"用户ID"` // 用户ID
	Username            string     `json:"username" comment:"用户名" vd:"len($)>0"`
	NickName            string     `json:"nickName" comment:"昵称" vd:"len($)>0"`
	Phone               string     `json:"phone" comment:"手机号" vd:"len($)>0"`
	Avatar              string     `json:"avatar" comment:"头像"`
	Sex                 string     `json:"sex" comment:"性别"`
	Birthday            string     `json:"birthday" comment:"出生日期"`
	Address             string     `json:"address" comment:"常驻地址"`
	Email               string     `json:"email" comment:"邮箱" vd:"len($)>0,email"`
	PostId              int        `json:"postId" comment:"岗位"`
	Status              string     `json:"status" comment:"状态" default:"1"`
	Remark              string     `json:"remark" comment:"备注"`
	InfoList            []InfoList `json:"infoList" comment:"用户关系"`
	DigitalLifeFileList []string   `json:"digital_life_file_list" comment:"数字生命"`
	RoleId              int        `json:"roleId" comment:"角色ID"`
	common.ControlBy
}

// SysUserGetResponse 获取用户信息返回结构
type SysUserGetResponse struct {
	Avatar              string     `json:"avatar" comment:"头像"`
	UserId              int        `json:"userId" comment:"用户ID"`
	Username            string     `json:"username" comment:"用户名"`
	NickName            string     `json:"nickName" comment:"昵称"`
	Phone               string     `json:"phone" comment:"手机号"`
	RoleId              int        `json:"roleId" comment:"角色ID"`
	Sex                 string     `json:"sex" comment:"性别"`
	Email               string     `json:"email" comment:"邮箱"`
	Birthday            string     `json:"birthday" comment:"出生日期"`
	Address             string     `json:"address" gorm:"comment:常驻地址"`
	Remark              string     `json:"remark" gorm:"comment:备注"`
	DigitalLifeFileList string     `json:"digital_life"  gorm:"comment:数字生命;column:digital_life"`
	Status              string     `json:"status" comment:"状态"`
	InfoList            []InfoList `json:"infoList" comment:"用户关系"`
}

// SysUserGetAll 获取用户信息返回结构
type SysUserGetAll struct {
	Avatar              string           `json:"avatar" comment:"头像"`
	UserId              int              `json:"userId" comment:"用户ID"`
	ParentId            int              `json:"parentId" comment:"上级ID"`
	Username            string           `json:"username" comment:"用户名"`
	NickName            string           `json:"nickName" comment:"昵称"`
	Phone               string           `json:"phone" comment:"手机号"`
	RoleId              int              `json:"roleId" comment:"角色ID"`
	Sex                 string           `json:"sex" comment:"性别"`
	Email               string           `json:"email" comment:"邮箱"`
	Birthday            string           `json:"birthday" comment:"出生日期"`
	Address             string           `json:"address" gorm:"comment:常驻地址"`
	Remark              string           `json:"remark" gorm:"comment:备注"`
	DigitalLifeFileList []string         `json:"digital_life"  gorm:"comment:数字生命;column:digital_life"`
	Status              string           `json:"status" comment:"状态"`
	InfoList            []InfoList       `json:"infoList" comment:"用户关系"`
	Children            []*SysUserGetAll `json:"children" comment:"下级"`
}

type SysUserOrder struct {
	UserIdOrder    string `search:"type:order;column:user_id;table:sys_user" form:"userIdOrder"`
	UsernameOrder  string `search:"type:order;column:username;table:sys_user" form:"usernameOrder"`
	StatusOrder    string `search:"type:order;column:status;table:sys_user" form:"statusOrder"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_user" form:"createdAtOrder"`
}

type DeptJoin struct {
	DeptId string `search:"type:contains;column:dept_path;table:sys_dept" form:"deptId"`
}

func (m *SysUserGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ResetSysUserPwdReq struct {
	UserId   int    `json:"userId" comment:"用户ID" vd:"$>0"` // 用户ID
	Password string `json:"password" comment:"密码" vd:"len($)>0"`
	common.ControlBy
}

func (s *ResetSysUserPwdReq) GetId() interface{} {
	return s.UserId
}

func (s *ResetSysUserPwdReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Password = s.Password
}

type UpdateSysUserAvatarReq struct {
	UserId int    `json:"userId" comment:"用户ID" vd:"len($)>0"` // 用户ID
	Avatar string `json:"avatar" comment:"头像" vd:"len($)>0"`
	common.ControlBy
}

func (s *UpdateSysUserAvatarReq) GetId() interface{} {
	return s.UserId
}

func (s *UpdateSysUserAvatarReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Avatar = s.Avatar
}

type UpdateSysUserStatusReq struct {
	UserId int    `json:"userId" comment:"用户ID" vd:"$>0"` // 用户ID
	Status string `json:"status" comment:"状态" vd:"len($)>0"`
	common.ControlBy
}

func (s *UpdateSysUserStatusReq) GetId() interface{} {
	return s.UserId
}

func (s *UpdateSysUserStatusReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Status = s.Status
}

type InfoList struct {
	RelationOptions int `json:"relationOptions" comment:"关系id"`
	UserListOptions int `json:"userListOptions" comment:"用户id"`
}

func (s *SysUserInsertReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Username = s.Username
	model.Password = s.Password
	model.NickName = s.NickName
	model.Phone = s.Phone
	model.RoleId = s.RoleId
	model.Sex = s.Sex
	model.Email = s.Email
	model.Remark = s.Remark
	model.Status = s.Status
	model.Address = s.Address
	model.CreateBy = s.CreateBy
	model.Birthday = carbon.Parse(s.Birthday, "PRC").Timestamp()

	parentId := 0
	for _, v := range s.InfoList {
		if v.RelationOptions == 1 {
			parentId = v.UserListOptions
		}
	}
	model.ParentUserId = parentId

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	str, e := json.Marshal(s.DigitalLifeFileList)
	if e == nil {
		model.DigitalLifeFileList = string(str)
	}
}

func (s *SysUserInsertReq) GenerateRelationData(model *[]models.SysUserRelationData, UserId int) {
	for _, v := range s.InfoList {
		*model = append(*model, models.SysUserRelationData{
			UserId:         UserId,
			UserRelationId: v.RelationOptions, //关系id
			RelationUserId: v.UserListOptions, //关联的用户id
		})
	}
}

func (s *SysUserUpdateReq) GenerateRelationData(model *[]models.SysUserRelationData, UserId int) {
	for _, v := range s.InfoList {
		*model = append(*model, models.SysUserRelationData{
			UserId:         UserId,
			UserRelationId: v.RelationOptions, //关系id
			RelationUserId: v.UserListOptions, //关联的用户id
		})
	}
}

func (s *SysUserInsertReq) GetId() interface{} {
	return s.UserId
}

func (s *SysUserUpdateReq) Generate(model *models.SysUser) {
	if s.UserId != 0 {
		model.UserId = s.UserId
	}
	model.Username = s.Username
	model.NickName = s.NickName
	model.Phone = s.Phone
	model.RoleId = s.RoleId
	model.Sex = s.Sex
	model.Email = s.Email
	model.Remark = s.Remark
	model.Status = s.Status
	model.Address = s.Address
	model.CreateBy = s.CreateBy
	model.Birthday = carbon.Parse(s.Birthday, "PRC").Timestamp()

	parentId := 0
	for _, v := range s.InfoList {
		if v.RelationOptions == 1 {
			parentId = v.UserListOptions
		}
	}
	model.ParentUserId = parentId

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	str, e := json.Marshal(s.DigitalLifeFileList)
	if e == nil {
		model.DigitalLifeFileList = string(str)
	}
}

func (s *SysUserUpdateReq) GetId() interface{} {
	return s.UserId
}

type SysUserById struct {
	dto.ObjectById
	common.ControlBy
}

func (s *SysUserById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func (s *SysUserById) GenerateM() (common.ActiveRecord, error) {
	return &models.SysUser{}, nil
}

// PassWord 密码
type PassWord struct {
	NewPassword string `json:"newPassword" vd:"len($)>0"`
	OldPassword string `json:"oldPassword" vd:"len($)>0"`
}
