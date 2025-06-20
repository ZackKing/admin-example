package logic

import (
	"admin-api/internal/dao"
	"admin-api/internal/dao/model"
	"admin-api/pkg/lib/def"
	"context"
)

type GroupLogic struct {
	ctx context.Context
}

var Group = &GroupLogic{context.Background()}

func NewGroup(ctx context.Context) *GroupLogic {
	return &GroupLogic{ctx}
}

func (l *GroupLogic) Search(s SearchConds) (gl []def.KV, le error) {

	gl = make([]def.KV, 0)

	g := dao.GetQuery().Group
	gq := g.WithContext(l.ctx)

	if s.Limit > 0 {
		gq = gq.Limit(s.Limit)
	}

	if s.Offset > 0 {
		gq = gq.Offset(s.Offset)
	}

	rows, le := gq.Find()
	if le != nil {
		return
	}

	if len(rows) > 0 {
		for _, row := range rows {
			gl = append(gl, def.KV{
				"id":           row.ID,
				"name":         row.Name,
				"status":       row.Status,
				"remark":       row.Remark,
				"created_time": row.CreatedTime,
				"mids":         l.GetMidsByGids([]int32{row.ID}),
				"uids":         l.GetUidsByGids([]int32{row.ID}),
			})
		}
	}

	return
}

func (l *GroupLogic) GetUidsByGids(gids []int32) []int32 {
	if len(gids) == 0 {
		return nil
	}

	ug := dao.GetQuery().GroupUser
	rows, err := ug.WithContext(l.ctx).Where(ug.Gid.In(gids...)).Select(ug.UID).Find()
	if err != nil {
		return nil
	}

	ids := make([]int32, 0)
	for _, v := range rows {
		ids = append(ids, v.UID)
	}

	return ids
}

func (l *GroupLogic) GetMidsByGids(gids []int32) []int32 {
	if len(gids) == 0 {
		return nil
	}

	um := dao.GetQuery().GroupMenu
	rows, err := um.WithContext(l.ctx).Where(um.Gid.In(gids...)).Select(um.Mid).Find()
	if err != nil {
		return nil
	}

	ids := make([]int32, 0)
	for _, v := range rows {
		ids = append(ids, v.Mid)
	}

	return ids
}

func (l *GroupLogic) AddInfo(ul *[]def.KV) {
	if len(*ul) == 0 {
		return
	}

	uids := make([]int32, 0)
	for _, v := range *ul {
		if uid, ok := v["uid"]; ok {
			uids = append(uids, uid.(int32))
		}
	}

	if len(uids) == 0 {
		return
	}

	gu := dao.GetQuery().GroupUser
	g := dao.GetQuery().Group

	rows := []struct {
		UID    int32
		Gid    int32
		Name   string
		Remark string
	}{}

	e := gu.WithContext(l.ctx).LeftJoin(g, g.ID.EqCol(gu.Gid)).
		Select(gu.UID, gu.Gid, g.Name, g.Remark).
		Where(gu.UID.In(uids...)).
		Scan(&rows)
	if e != nil {
		return
	}

	ugm := make(map[int32][]def.KV)

	for _, v := range rows {
		if _, ok := ugm[v.UID]; !ok {
			ugm[v.UID] = []def.KV{}
		}
		ugm[v.UID] = append(ugm[v.UID], def.KV{
			"id":     v.Gid,
			"name":   v.Name,
			"remark": v.Remark,
		})
	}

	for _, v := range *ul {
		if groups, ok := ugm[v["uid"].(int32)]; ok {
			v["group"] = groups
		} else {
			v["group"] = []def.KV{}
		}
	}

}

func (l *GroupLogic) SetMenu(id int32, mids []int32) error {
	if id == 0 {
		return nil
	}
	gm := dao.GetQuery().GroupMenu
	_, err := gm.WithContext(l.ctx).Where(gm.Gid.Eq(id)).Delete()
	if err != nil {
		return err
	}

	if len(mids) == 0 {
		return nil
	}
	for _, mid := range mids {
		gm.Create(&model.GroupMenu{Gid: id, Mid: mid})
	}

	return nil
}

func (l *GroupLogic) SetUser(id int32, uids []int32) error {
	if id == 0 {
		return nil
	}
	gu := dao.GetQuery().GroupUser
	_, err := gu.WithContext(l.ctx).Where(gu.Gid.Eq(id)).Delete()
	if err != nil {
		return err
	}

	if len(uids) == 0 {
		return nil
	}
	for _, uid := range uids {
		gu.Create(&model.GroupUser{Gid: id, UID: uid})
	}

	return nil
}
