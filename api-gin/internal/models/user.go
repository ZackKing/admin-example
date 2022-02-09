package models

import "fmt"

type UserRow struct {
	Uid         int    `db:"uid"`
	Name        string `db:"name"`
	Password    string `db:"password"`
	Salt        string `db:"salt"`
	RealName    string `db:"real_name"`
	Mobile      string `db:"mobile"`
	Email       string `db:"email"`
	Desc        string `db:"desc"`
	LoginTime   int    `db:"login_time"`
	PwdWrong    int    `db:"pwd_wrong"`
	Status      int    `db:"status"`
	CreatedTime string `db:"created_time"`
	UpdatedTime string `db:"updated_time"`
}

type UserModel struct {
	Conn  string
	Table string
}

var UserMdl = &UserModel{"default", "user"}

func (m *UserModel) GetUserByName(name string) (u UserRow) {
	sqlStr := fmt.Sprintf("select * from %s where name = ? limit 1", m.Table)
	err := GetConnect("default").Get(&u, sqlStr, name)
	if err != nil {
		fmt.Printf("GetUserByName error: %v", err)
	}
	return u
}
