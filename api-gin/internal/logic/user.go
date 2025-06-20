package logic

import (
	"admin-api/internal/constant"
	"admin-api/internal/dao"
	"admin-api/internal/dao/model"
	"admin-api/pkg/lib/def"
	"admin-api/pkg/lib/text/sstr"
	"context"
)

type UserLogic struct {
	ctx context.Context
}

var User = &UserLogic{context.Background()}

func NewUser(ctx context.Context) *UserLogic {
	return &UserLogic{ctx}
}

func (l *UserLogic) Self(uid int32) (ud def.KV, le error) {
	q := dao.GetQuery().User
	user, le := q.WithContext(l.ctx).Where(q.UID.Eq(uid)).First()
	if le != nil {
		return
	}

	ud, le = dao.Row2KV(user)
	if le != nil {
		return
	}

	ud.Del("password")
	ud.Del("salt")
	return
}

type UserSearchConds struct {
	Limit    int
	Offset   int
	NameLike string
	GroupIds []int32
}

func (l *UserLogic) Search(s UserSearchConds) (ud def.KV, le error) {
	ud = def.KV{
		"total": 0,
		"list":  []def.KV{},
	}

	groupLogic := NewGroup(l.ctx)
	u := dao.GetQuery().User
	uq := u.WithContext(l.ctx)

	if s.NameLike != "" {
		uq.Where(u.Name.Like("%" + s.NameLike + "%"))
	}

	if len(s.GroupIds) > 0 {
		uids := groupLogic.GetUidsByGids(s.GroupIds)
		if len(uids) == 0 {
			return
		}
		uq.Where(u.UID.In(uids...))
	}

	cnt, le := uq.Count()
	if le != nil {
		return
	}

	if cnt > 0 {
		ud["total"] = cnt
		rows, le := uq.Limit(s.Limit).Offset(s.Offset).Find()
		if le != nil {
			return ud, le
		}
		list := []def.KV{}
		for _, v := range rows {
			list = append(list, def.KV{
				"uid":       v.UID,
				"name":      v.Name,
				"real_name": v.RealName,
				"mobile":    v.Mobile,
				"email":     v.Email,
				"desc":      v.Desc,
				"status":    v.Status,
				"group":     []def.KV{},
			})
		}
		groupLogic.AddInfo(&list)
		ud["list"] = list
	}

	return
}

func (l *UserLogic) SetGroup(uid int32, gids []int32) error {
	if uid == 0 {
		return nil
	}
	gu := dao.GetQuery().GroupUser
	_, err := gu.WithContext(l.ctx).Where(gu.UID.Eq(uid)).Delete()
	if err != nil {
		return err
	}

	if len(gids) == 0 {
		return nil
	}
	for _, gid := range gids {
		gu.Create(&model.GroupUser{Gid: gid, UID: uid})
	}

	return nil
}

func (l *UserLogic) Add(user *model.User) error {
	q := dao.GetQuery().User
	cnt, err := q.WithContext(l.ctx).Where(q.Name.Eq(user.Name)).Count()
	if err != nil || cnt > 0 {
		return NewCodeErr(10007)
	}

	user.Salt = sstr.Random(6, sstr.DefStrSet)
	user.Password = Auth.EncodePwd(user.Password, user.Salt)
	user.Status = constant.UserStatusActive
	user.PwdWrong = 0
	if err = q.WithContext(l.ctx).Create(user); err != nil {
		return NewCodeErr(10006)
	}

	return nil
}

func (l *UserLogic) ChangePwd(uid int32, pwd string) error {
	if uid < 1 {
		return nil
	}
	q := dao.GetQuery().User
	user, err := q.WithContext(l.ctx).Where(q.UID.Eq(uid)).First()
	if err != nil || user == nil {
		return NewCodeErr(10003)
	}

	user.Salt = sstr.Random(6, sstr.DefStrSet)
	user.Password = Auth.EncodePwd(pwd, user.Salt)
	_, err = q.WithContext(l.ctx).Where(q.UID.Eq(uid)).Updates(user)
	if err != nil {
		return NewMsgErr(1, "update password error")
	}

	return nil
}
