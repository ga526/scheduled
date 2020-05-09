package list

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"scheduled/common"
	"scheduled/models"
	"scheduled/singleton"
)

/**
 * 生成下个月 的月表
 * 每月执行一次
 * 原mysql定时任务  jlmjdl_fx.month
 */
type Daifu struct {
	r        common.R
	taskName string
	method   string
}


func (this Daifu) RunCicle() {
	db,_ := orm.GetDB("recharge")
	o,_ := orm.NewOrmWithDB("mysql","recharge",db)
	queryBuilder, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		panic("database recharge not be set !")
	}
	sql := queryBuilder.Select("*").From("withdraw").Where("status=0").String()
	var lists []models.Withdraw
	_, err = o.Raw(sql).QueryRows(&lists)
	if err != nil {
		singleton.LogSingleton().Error("代付检测执行失败,原因为:" + err.Error())
		return
	}
	for _, withdraw := range lists {
		this.checkDaifu(withdraw)
	}
}


func (this Daifu) checkDaifu(withdraw models.Withdraw){
	if withdraw.ExchangeType == "user" {
		db,_ := orm.GetDB("log_comm_jlmj")
		o,_ := orm.NewOrmWithDB("mysql","log_comm_jlmj",db)
		queryBuilder, _ := orm.NewQueryBuilder("mysql")
		sql := queryBuilder.Select("platform").From("exchange").Where(fmt.Sprintf("order_num=%s",withdraw.OrderNumber)).String()
		exchange := models.Exchange{}
		err := o.Raw(sql).QueryRow(&exchange)
		if err != nil || exchange.Id == 0 {
			return
		}
		this._checkDaifu(exchange.Platform)
	}else{
		o := orm.NewOrm()
		o.Using("jinliu_agent2")
		queryBuilder, _ := orm.NewQueryBuilder("mysql")
		sql := queryBuilder.Select("platform").From("spread_tx").Where(fmt.Sprintf("order_num=%s",withdraw.OrderNumber)).String()
		spreadTx := models.SpreadTx{}
		err := o.Raw(sql).QueryRow(&spreadTx)
		if err != nil || spreadTx.TxId == 0 {
			return
		}
		this._checkDaifu(spreadTx.Platform)
	}
}

func (this Daifu) _checkDaifu(platform string){
	v := reflect.ValueOf(Daifu{}).Elem() //task需要是引用
	m := v.MethodByName(fmt.Sprintf("call_%s",platform))
	m.Call([]reflect.Value{})
}

func (this Daifu) call123(){

}
