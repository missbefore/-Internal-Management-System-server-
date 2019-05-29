package schedules

import (
	"wc_inner_server/model/entity"
	"wc_inner_server/model/table"
	"wc_inner_server/model/tool"
)

var db = tool.DB
func GetDataList(whereNames string, whereValues []interface{},
	schedules *[]model.DraftSchedules) error {

	if err := db.Select([]string{"draft_merchants_no", "modified", "handler", "status", "sales_no", "no"}).
		Preload("Sales").
		Where(whereNames, whereValues...).Find(&schedules).Error;err != nil {
		return err
	}
	var agent model.Agents
	var merchant model.DraftMerchants
	var schedulesArray =  *schedules
	for k, v := range schedulesArray {
		agentNo := v.Sales.AgentNo
		db.Select([]string{"no", "name", "country_code"}).Where("no=?", agentNo).First(&agent)
		schedulesArray[k].Sales.Agent = agent

		merchantsNo := v.DraftMerchantsNo
		db.Select([]string{"no", "short_name"}).Where("no=?", merchantsNo).First(&merchant)
		schedulesArray[k].DraftMerchants = merchant
	}

	if errSchedule := entity.DealSchedulesList(schedules); errSchedule != nil {
		return errSchedule
	}

	return nil
}
