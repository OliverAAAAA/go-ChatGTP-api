package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

// AskLog struct
type AskLog struct {
	Id         int       `orm:"column(id);auto;size(11)" description:"表ID"`
	UserId     int       `orm:"column(user_id);size(11)" description:"用户id"`
	Method     string    `orm:"column(method);size(30)" description:"请求方法"`
	Request    string    `orm:"column(request);type(text)" description:"问题"`
	RequestIp  string    `orm:"column(request_ip);size(30)" description:"请求iP"`
	Address    string    `orm:"column(address);size(30)" description:"请求地址"`
	Content    string    `orm:"column(data);type(text)" description:"日志内容"`
	CreateTime time.Time `orm:"column(createTime);type(datetime)" description:"创建时间"`
}

// TableName 自定义table 名称
func (*AskLog) TableName() string {
	return "ask_log"
}

// SearchField 定义模型的可搜索字段
func (*AskLog) SearchField() []string {
	return []string{}
}

// WhereField 定义模型可作为条件的字段
func (*AskLog) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*AskLog) TimeField() []string {
	return []string{}
}

// NoDeletionId 禁止删除的数据id
func (*AskLog) NoDeletionId() []int {
	return []int{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(AskLog))
}
