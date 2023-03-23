package models

import (
	"time"
)

type MenuModel struct {
	Conn  string
	Table string
}

var MenuMdl = &MenuModel{"default", "menu"}

type MenuRow struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`   //  菜单名称
	Uri         string    `db:"uri"`    //  uri地址
	Level       int64     `db:"level"`  //  节点等级
	Pid         int64     `db:"pid"`    //  父节点菜单id
	Icon        string    `db:"icon"`   //  菜单图标
	Status      int64     `db:"status"` //  状态（0-禁用，1-正常）
	Sort        int64     `db:"sort"`   //  排序
	Remark      string    `db:"remark"`
	CreatedTime time.Time `db:"created_time"`
	UpdatedTime time.Time `db:"updated_time"`
}

var MENU_STATUS_MAP = map[string]int{
	"valid":   1,
	"invalid": 0,
}
