package models

import (
	"time"
)

type User struct {
	Id          int       `gorm:"primaryKey,column:id" json:"id"`
	Name        string    `gorm:"column:name,uniqueIndex:uni_name" json:"name"`
	Password    string    `gorm:"column:password" json:"password"`
	Salt        string    `gorm:"column:salt" json:"salt"`
	RealName    string    `gorm:"column:real_name" json:"real_name"`
	Mobile      string    `gorm:"column:mobile" json:"mobile"`
	Email       string    `gorm:"column:email" json:"email"`
	Desc        string    `gorm:"column:desc" json:"desc"`
	LoginTime   int       `gorm:"column:login_time" json:"login_time"`
	PwdWrong    uint8     `gorm:"column:pwd_wrong" json:"pwd_wrong"`
	Status      uint8     `gorm:"column:status" json:"status"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
}

func (*User) TableName() string {
	return "user"
}

type userMdl struct {
	baseMdl
	StatusMap map[string]int
}

var UserMdl *userMdl

func init() {
	UserMdl = &userMdl{}
	UserMdl.dbName = "default"
	UserMdl.StatusMap = map[string]int{
		"valid":   1,
		"invalid": 0,
	}
}

// Get user by name
func (m *userMdl) GetValidUserByName(name string) (u *User) {
	db := m.getDb()
	db.First(&u, "name = ? and status = ?", name, m.StatusMap["valid"])
	return u
}
