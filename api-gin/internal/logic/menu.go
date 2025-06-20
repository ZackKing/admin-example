package logic

import (
	"admin-api/internal/dao"
	"admin-api/pkg/lib/def"
	"context"
)

type MenuLogic struct {
	ctx context.Context
}

var Menu = &MenuLogic{context.Background()}

func NewMenu(ctx context.Context) *MenuLogic {
	return &MenuLogic{ctx}
}

func (l *MenuLogic) Tree(mids []int32) []def.KV {
	q := dao.GetQuery().Menu
	mq := q.WithContext(l.ctx).Where(q.Status.Eq(1))
	if len(mids) > 0 {
		mq = mq.Where(q.ID.In(mids...))
	}

	menus, err := mq.Find()
	if err != nil {
		return []def.KV{}
	}

	ml := make([]def.KV, 0)
	for _, v := range menus {
		val, e := dao.Row2KV(v)
		if e != nil {
			return []def.KV{}
		}
		ml = append(ml, val)
	}

	return def.GenTree(ml, "id", "pid", "sub_menu")
}
