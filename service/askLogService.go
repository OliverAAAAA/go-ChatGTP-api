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
