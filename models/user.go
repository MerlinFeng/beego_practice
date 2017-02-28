package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//用户表模型
type User struct {
	Id       int
	UserName string `orm:"unique;size(15)"`
	Password string `orm:"size(32)"`
	Email    string `orm:"size(50)"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/mygo?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User))

	// create table
	//	orm.RunSyncdb("default", false, true)
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
