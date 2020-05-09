package models

import (
	"github.com/astaxie/beego/orm"
)

func init(){
	orm.RegisterModel(new(Exchange))
}

type Exchange struct{
	Id int `orm:"pk"`
	Uid int `orm:"column(uid)"`
	Type int `orm:"column(type)"`
	Mobile string `orm:"column(mobile)"`
	BankNum string `orm:"column(bank_num)"`
	Num int `orm:"column(num)"`
	Time int `orm:"column(time)"`
	Status int `orm:"column(status)"`
	WithdrawStatus int `orm:"column(withdraw_status)"`
	Fee float64 `orm:"column(fee)"`
	Result int `orm:"column(result)"`
	Remark string `orm:"column(remark)"`
	ChannelId string `orm:"column(channel_id)"`
	AuditTime string `orm:"column(audit_time)"`
	Operator string `orm:"column(operator)"`
	OrderNum string `orm:"column(order_num)"`
	OutOrderNum string `orm:"column(out_order_num)"`
	Platform string `orm:"column(platform)"`
	WithdrawIp string `orm:"column(withdraw_ip)"`
	IsClaim string `orm:"column(is_claim)"`
	ClaimId string `orm:"column(claim_id)"`
	Ip string `orm:"column(ip)"`
	Imei string `orm:"column(imei)"`
	LastGold string `orm:"column(last_gold)"`
}

// 自定义表名
func (u *Exchange) TableName() string {
	return "exchange"
}