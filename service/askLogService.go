package services

import (
	"github.com/beego/beego/v2/client/orm"
	"go-chatgpt-api/models"
)

// AskLogService struct
type AskLogService struct {
	BaseService
}

// Record 获取admin_log 总数
func (*AskLogService) Record(log *models.AskLog) int {
	count, err := orm.NewOrm().Insert(log)
	if err != nil {
		return 0
	}
	return int(count)
}

// QueryRencetAsk 最近一次的ask
func (*AskLogService) QueryRecentAsk(log *models.AskLog) *models.AskLog {
	qb, _ := orm.NewQueryBuilder("mysql")
	result := log
	// 构建查询对象
	qb.Select("*").
		From("ask_log").
		Where("request_ip = ?").And("method = 'ask'").And("request = ?").And("createTime >= DATE_SUB(NOW(), INTERVAL 10 MINUTE)").
		OrderBy("createTime").Desc().
		Limit(1)

	// 导出 SQL 语句
	sql := qb.String()

	// 执行 SQL 语句
	o := orm.NewOrm()
	//o.Raw(sql, &log.RequestIp, &log.Request).QueryRow(result)
	o.Raw(sql, log.RequestIp, log.Request).QueryRow(result)

	return result
}
