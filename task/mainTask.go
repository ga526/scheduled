package task

import (
	"github.com/astaxie/beego/orm"
	"scheduled/models"
	"scheduled/singleton"
)
type MainTask struct{}

func (this *MainTask) Todo() {
	//1. 数据库查询定时任务
	o := orm.NewOrm()
	CronO := singleton.CronSingleton()

	queryBuilder, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		panic(err.Error())
	}
	sql := queryBuilder.Select("*").From("crontab").String()
	var lists []models.Cron
	_, err = o.Raw(sql).QueryRows(&lists)
	if err != nil {
		singleton.LogSingleton().Error("定时任务启动失败,原因为:"+err.Error())
	}
	//2 类是否存在
	taskFactory := TaskFactory{}
	for _, value := range lists {
		if value.Status == 0 {
			continue
		}

		cronStr := value.Rule
		TaskName := value.TaskName
		MethodName := value.Method
		//3. 将任务插入到执行列表
		entryID,err := CronO.AddFunc(cronStr, func() {

			task := taskFactory.GetTask(TaskName)
			taskFactory.RunTask(task,TaskName,MethodName)

		})

		if err != nil {
			value.Status = 4
			value.Error = err.Error()
			o.Update(&value)
			singleton.LogSingleton().Error("更新域名状态错误,原因为"+ err.Error())
		}else{
			value.Status = 2
			value.EntryId = int(entryID)
			o.Update(&value)
		}
	}

	CronO.Start()
}
