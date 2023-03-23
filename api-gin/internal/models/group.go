package models

import "time"

type GroupModel struct {
	Conn  string
	Table string
}

var GroupMdl = &GroupModel{"default", "group"}

type GroupRow struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`   //  群组名称
	Status      int64     `db:"status"` //  状态（1-正常，0-禁用）
	Remark      string    `db:"remark"`
	CreatedTime time.Time `db:"created_time"`
	UpdatedTime time.Time `db:"updated_time"`
}
