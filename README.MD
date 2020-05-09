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

目录结构
~~~
www  WEB部署目录（或者子目录）
├─common           共用方法目录
│  ├─aesEncrypt.go      aes加解密
│  ├─R.go               统一结果返回结构体
│  ├─readFile.go        文件读取
│  ├─utils.go           工具类
│  └─ ...               更多文件
│
├─conf              配置目录
│  ├─app.go             应用总配置
│  ├─mysql.go           数据库配置
│  └─redis.go           redis配置
│  └─ ...               更多文件
│
├─controllers       控制器目录
│  ├─default.go         默认控制器
│  ├─parent.go          父类控制器
│  └─Task.go            任务管理控制器
│  └─ ...               更多文件
│  
├─docker            docker运行目录
│  └─ ...               更多文件
│  
├─models            模型
│  └─ ...               更多文件

├─routers           路由
│  └─ ...               更多文件

├─services          服务
│  └─ ...               更多文件

├─singleton         单例服务
│  ├─cron.go            定时任务单例
│  ├─logSingleton.go    日志单例    
│  └─redis.go           redis单例
│  └─ ...               更多文件

├─task              定时任务模块
│  ├─mainTask.go        定时任务入口文件
│  ├─TaskFactory.go     定时任务工厂方法  
│  ├─TaskImpl.go        定时任务接口
│  └─redis.go           redis单例
│  └─ ...               更多文件

├─thirdparty        第三方组件
│  └─ ...               更多文件
~~~


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