package v1

import (
	"admin-api/internal/dao"
	"admin-api/internal/dao/model"
	"admin-api/internal/logic"
	"admin-api/pkg/lib/def"

	"github.com/gin-gonic/gin"
)

type AuthApi struct {
	Api
}

var Auth = &AuthApi{}

func (a *AuthApi) Login(c *gin.Context) {
	data := &struct {
		Account  string `json:"account" binding:"required,min=5"`
		Password string `json:"password" binding:"required,min=6"`
	}{}
	if ok := a.ValidJson(c, &data); !ok {
		return
	}

	l := logic.NewAuth(c.Request.Context())

	ud, e := l.Login(data.Account, data.Password, def.KV{"ip": c.ClientIP()})
	if e != nil {
		a.LogicErrRes(c, e)
		return
	}
	a.OkRes(c, ud)
}

func (a *AuthApi) RenewToken(c *gin.Context) {
	uid := a.GetUid(c)
	token, ok := logic.Jwt.GenJwtToken(uid)
	if ok != nil {
		a.ErrRes(c, -1, "", nil)
	}
	a.OkRes(c, def.KV{"token": token})
}

func (a *AuthApi) JwtInfo(c *gin.Context) {
	uid := a.GetUid(c)
	a.OkRes(c, def.KV{"uid": uid})
}

func (a *AuthApi) MenuTree(c *gin.Context) {
	uid := a.GetUid(c)
	l := logic.NewAuth(c.Request.Context())

	mt, err := l.MenuTree(int32(uid))
	if err != nil {
		a.LogicErrRes(c, err)
		return
	}

	a.OkRes(c, mt)
}

func (a *AuthApi) ChangePwd(c *gin.Context) {
	data := &struct {
		Old      string `json:"old_password" binding:"required,min=6"`
		Password string `json:"password" binding:"required,min=6"`
	}{}
	if ok := a.ValidJson(c, &data); !ok {
		return
	}

	uid := int32(a.GetUid(c))
	err := logic.NewAuth(c.Request.Context()).ChangePwd(uid, data.Password, data.Old)
	if err != nil {
		a.LogicErrRes(c, err)
		return
	}
	a.OkRes(c, nil)
}

func (a *AuthApi) Group(c *gin.Context) {
	query := &struct {
		Size int `from:"size"`
		Page int `from:"page"`
	}{}
	if ok := a.ValidQuery(c, &query); !ok {
		return
	}

	conds := logic.SearchConds{
		Limit:  query.Size,
		Offset: 0,
	}

	if query.Size < 1 {
		conds.Limit = 10
	}

	if query.Page > 1 {
		conds.Offset = (query.Page - 1) * query.Size
	}

	gl, le := logic.NewGroup(c.Request.Context()).Search(conds)
	if le != nil {
		a.LogicErrRes(c, le)
		return
	}

	a.OkRes(c, gl)
}

func (a *AuthApi) AddGroup(c *gin.Context) {
	data := &struct {
		Name   string `json:"name" binding:"required,min=3"`
		Remark string `json:"remark"`
	}{}
	if ok := a.ValidJson(c, &data); !ok {
		return
	}

	group := model.Group{Name: data.Name, Remark: data.Remark}
	err := dao.GetQuery().Group.WithContext(c.Request.Context()).Create(&group)
	if err != nil {
		a.LogicErrRes(c, err)
		return
	}
	a.OkRes(c, group.ID)
}

func (a *AuthApi) ChangeGroup(c *gin.Context) {
	data := &struct {
		ID     int32  `json:"id" binding:"required"`
		Name   string `json:"name" binding:"required,min=3"`
		Remark string `json:"remark"`
		Status bool   `json:"status" binding:"required"`
	}{}
	if ok := a.ValidJson(c, &data); !ok {
		return
	}

	var status int32 = 0
	if data.Status {
		status = 1
	}

	g := dao.GetQuery().Group
	group, err := g.WithContext(c.Request.Context()).Where(g.ID.Eq(data.ID)).First()
	if err != nil || group.ID == 0 {
		a.ErrRes(c, 1, "group not found", nil)
		a.LogicErrRes(c, err)
		return
	}
	_, err = g.WithContext(c.Request.Context()).Where(g.ID.Eq(data.ID)).Updates(model.Group{
		Name:   data.Name,
		Remark: data.Remark,
		Status: status,
	})
	if err != nil {
		a.ErrRes(c, 1, "group update error", nil)
		return
	}

	a.OkRes(c, "ok")
}

func (a *AuthApi) SetGroupUser(c *gin.Context) {
	data := &struct {
		ID   int32   `json:"id" binding:"required"`
		Uids []int32 `json:"uids" binding:"required"`
	}{}
	if ok := a.ValidJson(c, &data); !ok {
		return
	}

	l := logic.NewGroup(c.Request.Context())
	if err := l.SetUser(data.ID, data.Uids); err != nil {
		a.LogicErrRes(c, err)
		return
	}
	a.OkRes(c, "ok")
}

func (a *AuthApi) SetGroupMenu(c *gin.Context) {
	data := &struct {
		ID   int32   `json:"id" binding:"required"`
		Mids []int32 `json:"menu_ids" binding:"required"`
	}{}
	if ok := a.ValidJson(c, &data); !ok {
		return
	}

	l := logic.NewGroup(c.Request.Context())
	if err := l.SetMenu(data.ID, data.Mids); err != nil {
		a.LogicErrRes(c, err)
		return
	}
	a.OkRes(c, "ok")
}

func (a *AuthApi) SetMenuGroup(c *gin.Context) {
	uid := a.GetUid(c)
	a.OkRes(c, def.KV{"uid": uid})
}

func (a *AuthApi) Menu(c *gin.Context) {
	mt := logic.Menu.Tree(nil)
	if mt == nil {
		a.ErrRes(c, 1, "menu tree error", nil)
		return
	}
	a.OkRes(c, mt)
}

func (a *AuthApi) MenuInfo(c *gin.Context) {
	a.OkRes(c, "ok")
}
