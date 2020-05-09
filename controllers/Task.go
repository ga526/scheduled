package controllers

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	cron2 "gopkg.in/robfig/cron.v2"
	"scheduled/common"
	"scheduled/models"
	"scheduled/singleton"
	"scheduled/task"
)

type Task struct {
	parent
}

/**
 * 关闭定时任务
 */
func (this *Task) Stop() {
	type taskForm struct {
		ID int `form:"id"`
	}
	taskParam := taskForm{}
	this.ParseForm(&taskParam)

	if taskParam.ID == 0 {
		this.Ctx.WriteString(this.r.Error("参数错误").ToString())
		return
	}
	this.Ctx.WriteString(this._stop(taskParam.ID).ToString())
}

/**
 * 开启定时任务
 */
func (this *Task) Start() {
	type taskForm struct {
		ID int `form:"id"`
	}
	taskParam := taskForm{}
	this.ParseForm(&taskParam)

	if taskParam.ID == 0 {
		this.Ctx.WriteString(this.r.Error("参数错误").ToString())
		return
	}
	this.Ctx.WriteString(this._start(taskParam.ID).ToString())
}

/**
 * 立即执行定时任务
 */
func (this *Task) Todo() {
	type taskForm struct {
		ID int `form:"id"`
	}
	taskParam := taskForm{}
	this.ParseForm(&taskParam)

	if taskParam.ID == 0 {
		this.Ctx.WriteString(this.r.Error("参数错误").ToString())
		return
	}
	this.Ctx.WriteString(this._todo(taskParam.ID).ToString())
}

func (this *Task) _stop(cronID int) common.R {
	o := orm.NewOrm()
	cron := models.Cron{}
	cron.Id = cronID
	err := o.Read(&cron, "id")
	if err != nil {
		return this.r.Error(err.Error())
	}
	if cron.EntryId == 0 {
		return this.r.Error("任务信息错误，实体ID 为 0")
	}
	if cron.Status == 3 {
		return this.r.Error("已经是暂停状态，无需再次暂停")
	}

	cronO := singleton.CronSingleton()
	cron.Status = 3
	cron.Error = ""
	entryID := cron.EntryId
	cron.EntryId = 0
	_, err = o.Update(&cron)
	if err != nil {
		return this.r.Error(err.Error())
	}
	cronEntryID := cron2.EntryID(entryID)
	cronO.Remove(cronEntryID)
	return this.r.Ok("暂停成功")
}
func (this *Task) _start(cronID int) common.R {
	o := orm.NewOrm()
	cron := models.Cron{}
	cron.Id = cronID
	err := o.Read(&cron, "id")
	if err != nil {
		return this.r.Error(err.Error())
	}
	if cron.Status == 2 || cron.Status == 1 {
		return this.r.Error("已经是启动状态，无需再次启动")
	}
	cronO := singleton.CronSingleton()
	cronStr := cron.Rule
	taskName := cron.TaskName
	methodName := cron.Method
	taskFactory := task.TaskFactory{}
	entryID, err := cronO.AddFunc(cronStr, func() {
		taskFactory.Todo(taskName, methodName)
	})
	cron.Status = 2
	cron.Error = ""
	cron.EntryId = int(entryID)
	_, err = o.Update(&cron)
	if err != nil {
		return this.r.Error(err.Error())
	}

	return this.r.Ok("启动成功")
}
func (this *Task) _todo(cronID int) common.R {
	o := orm.NewOrm()
	cron := models.Cron{}
	cron.Id = cronID
	err := o.Read(&cron, "id")
	if err != nil {
		return this.r.Error(err.Error())
	}
	if cron.Status == 1 {
		return this.r.Error("程序执行处理中")
	}
	taskName := cron.TaskName
	methodName := cron.Method
	taskFactory := task.TaskFactory{}
	taskFactory.Todo(taskName, methodName)
	cron.Status = cron.Status
	_, err = o.Update(&cron)
	if err != nil {
		return this.r.Error(err.Error())
	}

	return this.r.Ok("执行成功")
}