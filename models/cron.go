package models

import (
	"github.com/astaxie/beego/orm"
)

func init(){
	orm.RegisterModel(new(Cron))
}
type Cron struct {
	Id int `orm:"pk"`
	TaskName string `orm:"column(task_name)"`
	Method string `orm:"column(method)"`
	Rule string `orm:"column(rule)"`
	Status int `orm:"column(status)"`
	EntryId int `orm:"column(entry_id)"`
	Error string `orm:"column(error)"`
}

// 自定义表名
func (u *Cron) TableName() string {
	return "crontab"
}