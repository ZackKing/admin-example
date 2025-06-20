package logic

import (
	"admin-api/internal/constant"
	"admin-api/internal/dao"
	"admin-api/internal/dao/model"
	"admin-api/pkg/lib/crypto/md5"
	"admin-api/pkg/lib/def"
	"admin-api/pkg/lib/text/sstr"
	"context"
	"fmt"
	"time"
)

type AuthLogic struct {
	ctx context.Context
}

var Auth = &AuthLogic{context.Background()}

func NewAuth(ctx context.Context) *AuthLogic {
	return &AuthLogic{ctx}
}

func (l *AuthLogic) Login(name string, pwd string, ext def.KV) (data def.KV, le error) {
	q := dao.GetQuery().User

	user, err := q.WithContext(l.ctx).
		Select(q.UID, q.Name, q.Password, q.Salt, q.PwdWrong, q.Status).
		Where(q.Name.Eq(name)).
		First()

	if err != nil || user == nil || user.UID == 0 {
		le = NewCodeErr(10003)
		return
	}

	if user.Status == constant.UserStatusInactive {
		if user.PwdWrong >= constant.PwdWrongLimit {
			le = NewCodeErr(10009)
		} else {
			le = NewCodeErr(10005)
		}
		return
	}

	if l.EncodePwd(pwd, user.Salt) != user.Password {
		retry := l.checkBan(user)
		if retry > 0 {
			le = NewCodeErr(10008, retry)
		} else {
			le = NewCodeErr(10009)
		}
		return
	}

	token, err := Jwt.GenJwtToken(user.UID)

	if err != nil {
		le = NewMsgErr(-1, fmt.Sprintf("Jwt gen error: %s", err.Error()))
		return
	}

	data = def.KV{
		"uid":   user.UID,
		"token": token,
	}

	// after login
	q.WithContext(l.ctx).
		Where(q.UID.Eq(user.UID)).
		Select(q.LoginTime, q.PwdWrong).
		Updates(model.User{
			LoginTime: int32(time.Now().Unix()),
			PwdWrong:  0, // empty value update, field should in select
		})

	return
}

func (l *AuthLogic) ChangePwd(uid int32, pwd string, old string) error {
	q := dao.GetQuery().User
	user, err := q.WithContext(l.ctx).Where(q.UID.Eq(uid)).Select(q.UID, q.Password, q.Salt).First()

	if err != nil || user == nil || user.UID == 0 {
		return NewCodeErr(10003)
	}

	if l.EncodePwd(old, user.Salt) != user.Password {
		return NewCodeErr(10004)
	}

	salt := sstr.Random(6, "")

	_, err = q.WithContext(l.ctx).
		Where(q.UID.Eq(uid)).
		Updates(model.User{
			Salt:     salt,
			Password: l.EncodePwd(pwd, salt),
		})
	if err != nil {
		return NewMsgErr(1, "update password error")
	}

	return nil
}

func (l *AuthLogic) MenuTree(uid int32) (mt []def.KV, e error) {
	gq := dao.GetQuery().GroupUser

	gl, e := gq.WithContext(l.ctx).Where(gq.UID.Eq(uid)).Select(gq.Gid).Find()
	if e != nil {
		return
	}

	var mids = []int32{}
	if !l.hasAdminGroup(gl) { // get from group menu table
		gm := dao.GetQuery().GroupMenu
		gids := []int32{}
		for _, g := range gl {
			gids = append(gids, g.Gid)
		}

		gml, _ := gm.WithContext(l.ctx).Where(gm.Gid.In(gids...)).Find()
		for _, v := range gml {
			mids = append(mids, v.Mid)
		}
		if len(mids) == 0 {
			return
		}
	}
	mt = NewMenu(l.ctx).Tree(mids)
	return
}

func (*AuthLogic) EncodePwd(pwd string, salt string) string {
	return md5.Encode(salt + md5.Encode(pwd))
}

func (l *AuthLogic) checkBan(user *model.User) int32 {
	user.PwdWrong++

	if user.PwdWrong >= constant.PwdWrongLimit {
		user.Status = constant.UserStatusInactive
	}
	q := dao.GetQuery().User
	q.WithContext(l.ctx).Select(q.PwdWrong, q.Status).Updates(user)

	return constant.PwdWrongLimit - user.PwdWrong
}

func (l *AuthLogic) hasAdminGroup(gl []*model.GroupUser) bool {
	for _, g := range gl {
		if g.Gid == constant.GroupAdminID {
			return true
		}
	}
	return false
}
