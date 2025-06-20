package v1

import (
	"admin-api/internal/dao"
	"admin-api/internal/dao/model"
	"admin-api/internal/logic"
	"admin-api/pkg/lib/def"
	"admin-api/pkg/lib/text/sstr"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	Api
}

var User = &UserApi{}

func (a *UserApi) Self(c *gin.Context) {
	uid := a.GetUid(c)

	logic := logic.NewUser(c.Request.Context())
	ud, le := logic.Self(uid)
	if le != nil {
		a.LogicErrRes(c, le)
		return
	}
	a.OkRes(c, ud)
}

func (a *UserApi) Info(c *gin.Context) {
	v, ok := c.Get("uid")
	if !ok {
		v = nil
	}
	a.OkRes(c, v)
}

func (a *UserApi) ChangeSelf(c *gin.Context) {
	data := &struct {
		RealName string `json:"real_name" binding:"required,min=2"`
		Mobile   string `json:"mobile" binding:"required,min=1"`
		Email    string `json:"email" binding:"required,email"`
		Desc     string `json:"desc"`
	}{}
	if ok := a.ValidJson(c, &data); !ok {
		return
	}
	uid := a.GetUid(c)
	q := dao.GetQuery().User
	user, err := q.WithContext(c.Request.Context()).Where(q.UID.Eq(uid)).First()
	if err != nil || user.UID == 0 {
		a.ErrCodeRes(c, 10003)
		return
	}

	user.RealName = data.RealName
	user.Mobile = data.Mobile
	user.Email = data.Email
	user.Desc = data.Desc
	_, err = q.WithContext(c.Request.Context()).Where(q.UID.Eq(uid)).Updates(user)
	if err != nil {
		a.ErrRes(c, 1, "Update Info error", nil)
		return
	}

	a.OkRes(c, "ok")
}

func (a *UserApi) List(c *gin.Context) {
	query := &struct {
		Name     string `form:"name"`
		Size     int    `from:"size"`
		Page     int    `from:"page"`
		GroupIds string `form:"group_ids"`
	}{}

	if ok := a.ValidQuery(c, &query); !ok {
		return
	}

	if query.Size < 1 {
		query.Size = 10
	}

	if query.Page < 1 {
		query.Page = 1
	}

	conds := logic.UserSearchConds{
		Limit:    query.Size,
		Offset:   (query.Page - 1) * query.Size,
		NameLike: query.Name,
		GroupIds: []int32{},
	}

	if query.GroupIds != "" {
		ids := sstr.Split2Int32(query.GroupIds, ",")
		if len(ids) > 0 {
			conds.GroupIds = ids
		}
	}

	ud, le := logic.NewUser(c.Request.Context()).Search(conds)
	if le != nil {
		a.LogicErrRes(c, le)
		return
	}

	a.OkRes(c, ud)
}

func (a *UserApi) Add(c *gin.Context) {
	data := &struct {
		Name     string `json:"name" binding:"required,min=6"`
		Password string `json:"password" binding:"required,min=6"`
		RealName string `json:"real_name" binding:"required,min=2"`
		Mobile   string `json:"mobile" binding:"required"`
		Email    string `json:"email"`
		Desc     string `json:"desc"`
	}{}
	if ok := a.ValidJson(c, &data); !ok {
		return
	}

	user := model.User{
		Name:     data.Name,
		Password: data.Password,
		RealName: data.RealName,
		Mobile:   data.Mobile,
		Email:    data.Email,
		Desc:     data.Desc,
	}
	l := logic.NewUser(c.Request.Context())
	err := l.Add(&user)
	if err != nil {
		a.LogicErrRes(c, err)
		return
	}

	a.OkRes(c, def.KV{"uid": user.UID})
}

func (a *UserApi) Change(c *gin.Context) {
	data := &struct {
		UID      int32  `json:"uid" binding:"required"`
		Password string `json:"password" binding:"omitempty,min=6"`
		Name     string `json:"name" binding:"omitempty,min=6"`      // required
		RealName string `json:"real_name" binding:"omitempty,min=2"` // required
		Mobile   string `json:"mobile" binding:"omitempty,min=1"`    // required
		Email    string `json:"email"`
		Desc     string `json:"desc"`
	}{}
	if ok := a.ValidJson(c, &data); !ok {
		return
	}

	// change password
	if data.Password != "" {
		l := logic.NewUser(c.Request.Context())
		err := l.ChangePwd(data.UID, data.Password)
		if err != nil {
			a.LogicErrRes(c, err)
			return
		}
		a.OkRes(c, "ok")
		return
	}

	// change user info
	q := dao.GetQuery().User
	user, err := q.WithContext(c.Request.Context()).Where(q.UID.Eq(data.UID)).First()
	if err != nil || user == nil {
		a.ErrRes(c, 1, "user not found", nil)
		return
	}

	if data.Name != "" && data.Name != user.Name {
		cnt, err := q.WithContext(c.Request.Context()).Where(q.Name.Eq(data.Name)).Count()
		if err != nil || cnt > 0 {
			a.ErrRes(c, 10007, "", nil)
			return
		}
		user.Name = data.Name
	}
	if data.RealName != "" { // not empty
		user.RealName = data.RealName
	}
	if data.Mobile != "" { // not empty
		user.Mobile = data.Mobile
	}

	user.Email = data.Email
	user.Desc = data.Desc

	_, err = q.WithContext(c.Request.Context()).Where(q.UID.Eq(data.UID)).Updates(user)
	if err != nil {
		a.ErrRes(c, 1, "user update error", nil)
		return
	}

	a.OkRes(c, "ok")
}

func (a *UserApi) ChangeStatus(c *gin.Context) {
	data := &struct {
		UID    int32 `json:"uid" binding:"required"`
		Status int32 `json:"status" binding:"oneof=0 1"`
	}{}
	if ok := a.ValidJson(c, &data); !ok {
		return
	}

	q := dao.GetQuery().User
	_, err := q.WithContext(c.Request.Context()).Where(q.UID.Eq(data.UID)).Update(q.Status, data.Status)
	if err != nil {
		a.ErrRes(c, 1, "user not found", nil)
		return
	}

	a.OkRes(c, "ok")
}

func (a *UserApi) SetUserGroup(c *gin.Context) {
	data := &struct {
		ID   int32   `json:"id" binding:"required"`
		Gids []int32 `json:"group_ids" binding:"required"`
	}{}
	if ok := a.ValidJson(c, &data); !ok {
		return
	}
	l := logic.NewUser(c.Request.Context())
	err := l.SetGroup(data.ID, data.Gids)
	if err != nil {
		a.LogicErrRes(c, err)
		return
	}

	a.OkRes(c, "ok")
}
