package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"scheduled/common"
)

type parent struct {
	beego.Controller
	r common.R
}