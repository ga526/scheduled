package models

import (
	"github.com/astaxie/beego/orm"
)

func init(){
	orm.RegisterModel(new(Withdraw))
}

type Withdraw struct{
	Id int `orm:"pk"`
	OrderNumber string `orm:"column(order_number)"`
	ExchangeId int `orm:"column(exchange_id)"`
	PlatformId int `orm:"column(platform_id)"`
	Uid int `orm:"column(uid)"`
	Amount float64 `orm:"column(amount);digits(10);decimals(2)"`
	RealAmount float64 `orm:"column(real_amount);digits(10);decimals(2)"`
	TransFee float64 `orm:"column(trans_fee)"`
	Status int `orm:"column(status)"`
	ExchangeType string `orm:"column(exchange_type)"`
	CustomerName string `orm:"column(customer_name)"`
	OutCardnumber string `orm:"column(out_cardnumber)"`
	VerifiedTime string `orm:"column(verified_time)"`
	CreatedTime string `orm:"column(created_time)"`
}

// 自定义表名
func (u *Withdraw) TableName() string {
	return "withdraw"
}