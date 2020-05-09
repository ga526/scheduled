package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
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
	config, _ := config.NewConfig("ini", "./conf/mysql.conf")
	host := config.String("host")
	port := config.String("port")
	username := config.String("username")
	password := config.String("password")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/jlmj_config?charset=utf8",username,password,host,port))
	orm.RegisterDataBase("recharge", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/recharge?charset=utf8",username,password,host,port))
	//打印 orm 调试信息 ,如 sql
	orm.Debug = true
}

func doTask() {
	//1. 读取配置文件( 绝对路径 )
	mainTask := task.MainTask{}
	mainTask.Todo()
}
