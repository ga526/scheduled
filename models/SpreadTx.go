package models

import (
	"github.com/astaxie/beego/orm"
)

func init(){
	orm.RegisterModel(new(SpreadTx))
}

type SpreadTx struct{
	TxId int `orm:"pk;column(tx_id)"`
	TxNo string `orm:"column(tx_no)"`
	TxUid int `orm:"column(tx_uid)"`
	TxAmount int `orm:"column(tx_amount)"`
	TxFlag int `orm:"column(tx_flag)"`
	Meno float64 `orm:"column(meno)"`
	Addtime float64 `orm:"column(addtime)"`
	PayStatus int `orm:"column(pay_status)"`
	PayOrder string `orm:"column(pay_order)"`
	PayTime string `orm:"column(pay_time)"`
	TypeId int `orm:"column(type_id)"`
	LastUpdTime string `orm:"column(last_upd_time)"`
	ZfbNo string `orm:"column(zfb_no)"`
	ZfbRealName string `orm:"column(zfb_real_name)"`
	BankNo string `orm:"column(bank_no)"`
	BankName string `orm:"column(bank_name)"`
	BankBranchName string `orm:"column(bank_branch_name)"`
	BankHolderName string `orm:"column(bank_holder_name)"`
	Operator string `orm:"column(operator)"`
	OrderNum string `orm:"column(order_num)"`
	OutOrderNum string `orm:"column(out_order_num)"`
	Platform string `orm:"column(platform)"`
	ChannelId string `orm:"column(channel_id)"`
	WithdrawIp string `orm:"column(withdraw_ip)"`
}

// 自定义表名
func (u *SpreadTx) TableName() string {
	return "spread_tx"
}