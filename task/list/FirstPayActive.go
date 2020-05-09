package list

import (
	"github.com/astaxie/beego/orm"
	"scheduled/common"
)

/**
 * 首充活动过期后，设置为过期
 * 每 5 分钟执行一次
 * 原mysql定时任务  jlmj_config.e_pay_active_status
 */
type FirstPayActive struct{
	r common.R
	taskName string
	method string
}

func (this FirstPayActive) Todo(){
	o := orm.NewOrm()
	o.Raw("update pay_active set `status` = 1 where `status`=0 AND start_time>now();").Exec()
	o.Raw("update pay_active set `status`= 2 where end_time<=now() AND `status`=1;").Exec()
}
