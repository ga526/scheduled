本项目用于对 定时任务做统一处理
auth

1. 可以对外部提供 立即执行 定时任务 的API
2. 通过配置文件的方式,配置执行 执行指令
3. 可以通过外部 api 接口方式 停用/启用定时任务

依赖
1. github.com/go-sql-driver/mysql
2. github.com/go-redis/redis

使用
1. 在 game 后台添加任务（task_name,method） 必须在 task 目录下存在
2. 正常来说，只需要在 task 目录下添加自己的任务文件
3. 需要在 task/TaskFactory 下添加相应方法

配置文件地址
1. main.go   数据库地址
2. singleton\redis.go  redis 地址
3. task\list\MultiplyDomain  短信配置

Docker 运行
1. docker/build.sh 需要执行权限
2. docker-compose up --build -d

数据库脚本
1.  CREATE TABLE IF NOT EXISTS `crontab` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_name` varchar(50) DEFAULT NULL COMMENT '类名',
  `method` varchar(50) DEFAULT NULL COMMENT '方法名',
  `rule` varchar(50) DEFAULT NULL COMMENT '定时任务规则',
  `status` tinyint(4) DEFAULT NULL COMMENT '运行状态 1 运行中  2 启动  3 暂停',
  `entry_id` int(11) DEFAULT NULL,
  `error` varchar(50) DEFAULT NULL COMMENT '错误信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COMMENT='定时任务';

2. INSERT INTO `crontab` (`id`, `task_name`, `method`, `rule`, `status`, `entry_id`, `error`) VALUES
	(1, 'FirstTask', 'Todo', '* * * * * *', 2, 1, '');