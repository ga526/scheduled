package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "scheduled/common"
	_ "scheduled/routers"
	"scheduled/task"
)

func main() {
	/*dom := list.MultiplyDomain{}*/
	//初始化 mysql 配置
	initOrm()
	// 执行定时任务
	doTask()
	//beego 主程序
	beego.Run()
}

func initOrm() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1q@W3e$R@tcp(192.168.10.136:15642)/jlmj_config?charset=utf8")
	orm.RegisterDataBase("recharge", "mysql", "root:1q@W3e$R@tcp(192.168.10.136:15642)/recharge?charset=utf8")
	//打印 orm 调试信息 ,如 sql
	orm.Debug = true
}

func doTask() {
	//1. 读取配置文件( 绝对路径 )
	mainTask := task.MainTask{}
	mainTask.Todo()
}
