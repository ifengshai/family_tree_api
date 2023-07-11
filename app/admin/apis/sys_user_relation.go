package apis

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"html/template"
	"os"
	"strconv"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SysUserRelation struct {
	api.Api
}

// GetPage 获取用户关系表列表
// @Summary 获取用户关系表列表
// @Description 获取用户关系表列表
// @Tags 用户关系表
// @Param title query string false "关系名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysUserRelation}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user-relation [get]
// @Security Bearer
func (e SysUserRelation) GetPage(c *gin.Context) {
	req := dto.SysUserRelationGetPageReq{}
	s := service.SysUserRelation{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysUserRelation, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户关系表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户关系表
// @Summary 获取用户关系表
// @Description 获取用户关系表
// @Tags 用户关系表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysUserRelation} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user-relation/{id} [get]
// @Security Bearer
func (e SysUserRelation) Get(c *gin.Context) {
	req := dto.SysUserRelationGetReq{}
	s := service.SysUserRelation{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysUserRelation

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户关系表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建用户关系表
// @Summary 创建用户关系表
// @Description 创建用户关系表
// @Tags 用户关系表
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserRelationInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-user-relation [post]
// @Security Bearer
func (e SysUserRelation) Insert(c *gin.Context) {
	req := dto.SysUserRelationInsertReq{}
	s := service.SysUserRelation{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建用户关系表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户关系表
// @Summary 修改用户关系表
// @Description 修改用户关系表
// @Tags 用户关系表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysUserRelationUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-user-relation/{id} [put]
// @Security Bearer
func (e SysUserRelation) Update(c *gin.Context) {
	req := dto.SysUserRelationUpdateReq{}
	s := service.SysUserRelation{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改用户关系表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除用户关系表
// @Summary 删除用户关系表
// @Description 删除用户关系表
// @Tags 用户关系表
// @Param data body dto.SysUserRelationDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-user-relation [delete]
// @Security Bearer
func (e SysUserRelation) Delete(c *gin.Context) {
	s := service.SysUserRelation{}
	req := dto.SysUserRelationDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除用户关系表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// GetRelation 获取用户的关系
// @Summary 获取用户关系表
// @Description 获取用户关系表
// @Tags 用户关系表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysUserRelation} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user-relation/{id} [get]
// @Security Bearer
func (e SysUserRelation) GetRelation(c *gin.Context) {
	//声明关联
	s := service.SysUser{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	str, _ := os.Getwd()

	userAll := make([]*dto.SysUserGetAll, 0)
	//查询数据

	err = s.GetAll(&userAll)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	userAll = getTreeIterative(userAll, 0)

	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles(str + "/app/admin/views/relation.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	htmlChildren := echoHtml((*userAll[0]).Children)
	html := `
        <ul>
            <li>
                <a href="#">` + strconv.Itoa((*userAll[0]).UserId) + ":" + (*userAll[0]).Username + `</a>
                ` + htmlChildren + `
            </li>
        </ul>
`
	// 利用给定数据渲染模板，并将结果写入buf
	buf := new(bytes.Buffer) //实现了读写方法的可变大小的字节缓冲
	err = tmpl.Execute(buf,
		struct {
			Head template.HTML
			Body template.HTML
		}{
			Head: "<script async src=\"https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-5141030920527699\" crossorigin=\"anonymous\"></script>",
			Body: template.HTML(html),
		},
	)
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, buf.String())
}

func getTreeIterative(list []*dto.SysUserGetAll, parentId int) []*dto.SysUserGetAll {
	memo := make(map[int]*dto.SysUserGetAll, 0)
	for _, v := range list {
		if _, ok := memo[v.UserId]; ok {
			v.Children = memo[v.UserId].Children
			memo[v.UserId] = v
		} else {
			v.Children = make([]*dto.SysUserGetAll, 0)
			memo[v.UserId] = v
		}
		if _, ok := memo[v.ParentId]; ok {
			memo[v.ParentId].Children = append(memo[v.ParentId].Children, memo[v.UserId])
		} else {
			memo[v.ParentId] = &dto.SysUserGetAll{Children: []*dto.SysUserGetAll{memo[v.UserId]}}
		}
	}
	return memo[parentId].Children
}

func echoHtml(list []*dto.SysUserGetAll) string {
	tempString := ""
	if len(list) > 0 {
		//已输出的用户不重复输出
		walked := make(map[int]bool, 0)
		tempString = tempString + "<ul>"
		for _, v := range list {
			if _, ok := walked[v.UserId]; !ok {

				//同级共同下级
				children := make([]*dto.SysUserGetAll, 0)
				children = append(children, v.Children...)

				walked[v.UserId] = true

				//查寻当前用户的同级关系，如夫妻关系
				family := make(map[int]*dto.SysUserGetAll, 0)
				for _, vR := range v.InfoList {
					if vR.RelationOptions != 0 {
						family[vR.UserListOptions] = v
						walked[vR.UserListOptions] = true
					}
				}
				for _, vLL := range list {
					for _, vR := range vLL.InfoList {
						if vR.RelationOptions != 0 && vR.UserListOptions == v.UserId {
							family[vLL.UserId] = vLL
							walked[vLL.UserId] = true
						}
					}
				}

				tempString = tempString + "<li>"
				tempStringFamily := ""
				i := 0
				for _, f := range family {
					i += 1
					if len(family) == i {
						tempStringFamily += strconv.Itoa(f.UserId) + ":" + f.Username
					} else {
						tempStringFamily += strconv.Itoa(f.UserId) + ":" + f.Username + " "
					}
				}
				tempString = tempString + "<a href='#'>" + tempStringFamily + "</a>"

				for _, vL := range list {
					if vL.UserId != v.UserId {
						if _, o := family[vL.UserId]; o {
							children = append(children, vL.Children...)
						}
					}
				}
				if len(children) > 0 {
					tempString = tempString + echoHtml(children)
				}
				tempString = tempString + "</li>"
			}
		}
		tempString = tempString + "</ul>"
	}
	return tempString
}
