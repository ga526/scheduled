package singleton

import (
	"gopkg.in/robfig/cron.v2"
	"sync"
)

var cronInstance *cron.Cron
var cronOnce sync.Once

/**
# ┌───────────── second (0 - 59)
# | ┌───────────── minute (0 - 59)
# | │ ┌───────────── hour (0 - 23)
# | │ │ ┌───────────── day of the month (1 - 31)
# | │ │ │ ┌───────────── month (1 - 12)
# | │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday;
# | │ │ │ │ │                                   7 is also Sunday on some systems)
# | │ │ │ │ │
# | │ │ │ │ │
# * * * * * * command to execute
*/
//http://labix.org/mgo
func CronSingleton() *cron.Cron {
	cronOnce.Do(func() {
		cronInstance = cron.New()
	})
	return cronInstance
}
