package task

import (
	"github.com/astaxie/beego/orm"
	"reflect"
	"scheduled/models"
	"scheduled/task/list"
)

/**
 * 任务工厂类
 */
type TaskFactory struct{}

func (this *TaskFactory) Todo(taskName string, method string) {
	/*  由于 golang 无法直接根据字符串反射出类型实例(https://stackoverflow.com/questions/23030884/is-there-a-way-to-create-an-instance-of-a-struct-from-a-string)
	 *  所以需要手动写入，或者遍历文件夹去写一个注册器（）
	 */
	var task TaskImpl
	switch taskName {
	case "FirstTask":
		task = &list.FirstTask{}
	case "FirstPayActive":
		task = &list.FirstPayActive{}
	case "Daifu":
		task = &list.Daifu{}
	default:
		panic("hasn't this task : " + taskName)
	}

	this.begin(taskName,method)
	v := reflect.ValueOf(task).Elem() //task需要是引用
	m := v.MethodByName(method)
	m.Call([]reflect.Value{})
	this.end(taskName,method)
}


func (this *TaskFactory) begin(taskName string,method string){
	cron := models.Cron{}
	cron.TaskName = taskName
	cron.Method = method
	o := orm.NewOrm()
	err := o.Read(&cron,"task_name","method")
	if err != nil{
		return
	}
	cron.Status = 1
	o.Update(&cron,"status")
}

func (this *TaskFactory) end(taskName string,method string){
	cron := models.Cron{}
	cron.TaskName = taskName
	cron.Method = method
	o := orm.NewOrm()
	err := o.Read(&cron,"task_name","method")
	if err != nil{
		return
	}
	cron.Status = 2
	o.Update(&cron,"status")
}