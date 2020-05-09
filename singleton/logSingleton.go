package singleton

import (
	"github.com/astaxie/beego/logs"
	"sync"
)

var logInstance *logs.BeeLogger
var logOnce sync.Once
//http://labix.org/mgo
func LogSingleton() *logs.BeeLogger{
	cronOnce.Do(func() {
		logInstance = logs.NewLogger(10000)
		logInstance.SetLogger(logs.AdapterFile, `{"filename":"go_log/test.log"}`)
	})
	return logInstance
}
