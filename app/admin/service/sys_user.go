package service

import (
	"encoding/json"
	"errors"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/golang-module/carbon/v2"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"gorm.io/gorm"
)

type SysUser struct {
	service.Service
}

// GetPage 获取SysUser列表
func (e *SysUser) GetPage(c *dto.SysUserGetPageReq, p *actions.DataPermission, list *[]models.SysUser, count *int64) error {
	var err error
	var data models.SysUser

	err = e.Orm.Debug().
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Get 获取SysUser对象
func (e *SysUser) Get(d *dto.SysUserById, p *actions.DataPermission, response *dto.SysUserGetResponse) error {
	var data models.SysUser
	var model *models.SysUser
	var relationData []*models.SysUserRelationData
	var relation []dto.InfoList

	err := e.Orm.Model(&data).Debug().
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(&model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	err = e.Orm.Model(&models.SysUserRelationData{}).Debug().
		Preload("SysUserRelation").
		Where("user_id = ?", d.GetId()).
		Find(&relationData).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	response.UserId = model.UserId
	response.Username = model.Username
	response.NickName = model.NickName
	response.Phone = model.Phone
	response.RoleId = model.RoleId
	response.Address = model.Address
	response.Sex = model.Sex
	response.Email = model.Email
	response.Birthday = carbon.CreateFromTimestamp(model.Birthday, "PRC").ToDateTimeString()
	response.Remark = model.Remark
	response.DigitalLifeFileList = model.DigitalLifeFileList
	response.Status = model.Status

	for _, v := range relationData {
		relation = append(relation, dto.InfoList{
			RelationOptions: v.UserRelationId,
			UserListOptions: v.RelationUserId,
		})
	}
	response.InfoList = relation
	return nil
}

// Insert 创建SysUser对象
func (e *SysUser) Insert(c *dto.SysUserInsertReq) error {
	var err error
	var data models.SysUser
	var relationData []models.SysUserRelationData
	var i int64
	err = e.Orm.Model(&data).Where("username = ?", c.Username).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("用户名已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}

	parentId := 0
	for _, v := range c.InfoList {
		if v.RelationOptions == 1 {
			parentId = v.RelationOptions
		}
	}
	if parentId == 0 {
		err := errors.New("未设置上一代关系！")
		return err
	}

	c.Generate(&data)

	//事务测试
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	//更新用户关联关系表
	err = e.Orm.Create(&data).Error
	c.GenerateRelationData(&relationData, data.UserId)
	for _, v := range relationData {
		err = e.Orm.Create(&v).Error
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}
	}

	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Update 修改SysUser对象
func (e *SysUser) Update(c *dto.SysUserUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).Where("user_id = ?", c.GetId()).First(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	//校验数据
	var data models.SysUser
	var i int64
	err = e.Orm.Model(&data).Where("username = ?", c.Username).Where("user_id <> ?", c.UserId).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("用户名已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}

	parentId := 0
	for _, v := range c.InfoList {
		if v.RelationOptions == 1 {
			parentId = v.RelationOptions
		}
	}
	if parentId == 0 {
		err := errors.New("未设置上一代关系！")
		return err
	}

	c.Generate(&data)
	update := e.Orm.Model(&model).Where("user_id = ?", &model.UserId).Omit("password", "salt").Updates(&data)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update userinfo error")
		log.Warnf("db update error")
		return err
	}

	var relationData []models.SysUserRelationData
	// 使用 Unscoped 方法永久删除记录
	e.Orm.Unscoped().Debug().Where("user_id = ?", &model.UserId).Delete(models.SysUserRelationData{})
	c.GenerateRelationData(&relationData, model.UserId)
	for _, v := range relationData {
		err = e.Orm.Create(&v).Error
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}
	}

	return nil
}

// UpdateAvatar 更新用户头像
func (e *SysUser) UpdateAvatar(c *dto.UpdateSysUserAvatarReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = e.Orm.Table(model.TableName()).Where("user_id =? ", c.UserId).Updates(c).Error
	if err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	return nil
}

// UpdateStatus 更新用户状态
func (e *SysUser) UpdateStatus(c *dto.UpdateSysUserStatusReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	err = e.Orm.Table(model.TableName()).Where("user_id =? ", c.UserId).Updates(c).Error
	if err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	return nil
}

// ResetPwd 重置用户密码
func (e *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := e.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("At Service ResetSysUserPwd error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	c.Generate(&model)
	err = e.Orm.Omit("username", "nick_name", "phone", "role_id", "avatar", "sex").Save(&model).Error
	if err != nil {
		e.Log.Errorf("At Service ResetSysUserPwd error: %s", err)
		return err
	}
	return nil
}

// Remove 删除SysUser
func (e *SysUser) Remove(c *dto.SysUserById, p *actions.DataPermission) error {
	var err error
	var data models.SysUser

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveSysUser : %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// UpdatePwd 修改SysUser对象密码
func (e *SysUser) UpdatePwd(id int, oldPassword, newPassword string, p *actions.DataPermission) error {
	var err error

	if newPassword == "" {
		return nil
	}
	c := &models.SysUser{}

	err = e.Orm.Model(c).
		Scopes(
			actions.Permission(c.TableName(), p),
		).Select("UserId", "Password", "Salt").
		First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("无权更新该数据")
		}
		e.Log.Errorf("db error: %s", err)
		return err
	}
	var ok bool
	ok, err = pkg.CompareHashAndPassword(c.Password, oldPassword)
	if err != nil {
		e.Log.Errorf("CompareHashAndPassword error, %s", err.Error())
		return err
	}
	if !ok {
		err = errors.New("incorrect Password")
		e.Log.Warnf("user[%d] %s", id, err.Error())
		return err
	}
	c.Password = newPassword
	db := e.Orm.Model(c).Where("user_id = ?", id).
		Select("Password", "Salt").
		Updates(c)
	if err = db.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("set password error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

func (e *SysUser) GetProfile(c *dto.SysUserById, user *models.SysUser, roles *[]models.SysRole, posts *[]models.SysPost) error {
	err := e.Orm.Preload("Dept").First(user, c.GetId()).Error
	if err != nil {
		return err
	}
	err = e.Orm.Find(roles, user.RoleId).Error
	if err != nil {
		return err
	}

	return nil
}

// GetAll 获取SysUser对象
func (e *SysUser) GetAll(response *[]*dto.SysUserGetAll) error {
	var data models.SysUser
	var user []models.SysUser
	var relationData []models.SysUserRelationData
	tempRelationData := make(map[int][]models.SysUserRelationData)
	//var relation []dto.InfoList

	//查询所有用户
	err := e.Orm.Model(&data).Debug().Where("status = ?", 2).Find(&user).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	//查询所有关系
	err = e.Orm.Model(&models.SysUserRelationData{}).Debug().
		Preload("SysUserRelation").
		Find(&relationData).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	for _, v := range relationData {
		tempRelationData[v.UserId] = append(tempRelationData[v.UserId], v)
	}

	for _, v := range user {
		var temp *dto.SysUserGetAll
		temp = new(dto.SysUserGetAll)
		var tempRelation []dto.InfoList

		temp.UserId = v.UserId
		temp.ParentId = v.ParentUserId
		temp.Username = v.Username
		temp.NickName = v.NickName
		temp.Phone = v.Phone
		temp.RoleId = v.RoleId
		temp.Address = v.Address
		temp.Sex = v.Sex
		temp.Email = v.Email
		temp.Birthday = carbon.CreateFromTimestamp(v.Birthday, "PRC").ToDateTimeString()
		temp.Remark = v.Remark

		err := json.Unmarshal([]byte(v.DigitalLifeFileList), &temp.DigitalLifeFileList)
		if err != nil {
			return err
		}
		temp.Status = v.Status

		for _, v := range tempRelationData[v.UserId] {
			tempRelation = append(tempRelation, dto.InfoList{
				RelationOptions: v.UserRelationId,
				UserListOptions: v.RelationUserId,
			})
		}
		temp.InfoList = tempRelation
		*response = append(*response, temp)
	}

	return nil
}
