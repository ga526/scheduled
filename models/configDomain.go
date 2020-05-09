package models

import (
	"github.com/astaxie/beego/orm"
)

func init(){
	orm.RegisterModel(new(ConfigDomain))
}

type ConfigDomain struct{
	Id int `orm:"pk"`
	Url string `orm:"column(url)"`
	Status int `orm:"column(status)"`
	Type int `orm:"column(type)"`
	Weight int `orm:"column(weight)"`
	ErrNo int `orm:"column(err_no)"`
	UpdateTime string `orm:"column(update_time)"`
}

// 自定义表名
func (u *ConfigDomain) TableName() string {
	return "config_domain"
}