package service

import (
	"github.com/astaxie/beego/orm"
	"scheduled/common"
)

/**
 * 任务主框架 -- 主要用于外部API 调用
 * 方法  stop  run  start
 */
type Schema struct {
	r common.R
}

/**
 * 判断表是否存在
 */
func (this *Schema) TableExists(tableName string,dbName string) bool{
	if(tableName == ""){
		return false
	}

	schemaStr := "TABLE_NAME='"+tableName+"'"
	if(dbName != ""){
		schemaStr = schemaStr+" AND TABLE_SCHEMA='"+dbName+"'"
	}

	type result struct{
		Count int
	}
	o := orm.NewOrm()
	var res result
	err := o.Raw("SELECT count(*) count FROM information_schema.TABLES WHERE "+schemaStr).QueryRow(&res)
	if(err != nil){
		return false
	}
	return res.Count == 1
}

