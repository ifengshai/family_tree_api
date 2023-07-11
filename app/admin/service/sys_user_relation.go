package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SysUserRelation struct {
	service.Service
}

// GetPage 获取SysUserRelation列表
func (e *SysUserRelation) GetPage(c *dto.SysUserRelationGetPageReq, p *actions.DataPermission, list *[]models.SysUserRelation, count *int64) error {
	var err error
	var data models.SysUserRelation

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SysUserRelationService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SysUserRelation对象
func (e *SysUserRelation) Get(d *dto.SysUserRelationGetReq, p *actions.DataPermission, model *models.SysUserRelation) error {
	var data models.SysUserRelation

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysUserRelation error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysUserRelation对象
func (e *SysUserRelation) Insert(c *dto.SysUserRelationInsertReq) error {
	var err error
	var data models.SysUserRelation
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SysUserRelationService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SysUserRelation对象
func (e *SysUserRelation) Update(c *dto.SysUserRelationUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SysUserRelation{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("SysUserRelationService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysUserRelation
func (e *SysUserRelation) Remove(d *dto.SysUserRelationDeleteReq, p *actions.DataPermission) error {
	var data models.SysUserRelation

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSysUserRelation error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
