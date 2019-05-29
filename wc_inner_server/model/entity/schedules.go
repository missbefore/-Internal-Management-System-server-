package entity

import (
	"strings"
	"time"
	"wc_inner_server/common/code"
	model "wc_inner_server/model/table"
)

func DealSchedulesList(schedules *[]model.DraftSchedules) error {

	schedulesArray := *schedules

	for k, v := range schedulesArray {
		modified, _ := time.Parse(Time, strings.Split(v.Modified, "T")[0])
		schedulesArray[k].Modified = modified.Format(Time)

		if v.Handler != "1" {
			switch v.Status  {
			case "KYC":
				schedulesArray[k].Sales.Name = "风控"
				break
			case "SECONDKYC":
				schedulesArray[k].Sales.Name = "风控"
				break
			case "CONTRACT":
				schedulesArray[k].Sales.Name = "运营"
				break
			case "OPENING":
				schedulesArray[k].Sales.Name = "运营"
				break
			case "ACTIVE":
				schedulesArray[k].Sales.Name = "运营"
				break
			}
		}

		var status string
		switch v.Status {
		case "START":
			status = "资料提交中"
			break
		case "OPENING":
			status = "商户开通中"
			break
		case "ACTIVE":
			status = "已开通"
			break
		case "ABOLITION":
			status = "拒绝开通"
			break
		}

		if v.Handler=="0" {
			switch v.Status {
			case "KYC":
				status="风控初审中"
				break
			case "SECONDKYC":
				status = "风控制复审中"
				break
			case "CONTRACT":
				status = "合同审核中"
				break
			}
		} else {
			switch v.Status {
			case "KYC":
				status="风控驳回"
				break
			case "CONTRACT":
				status = "提交合同中"
				break
			}
		}

		schedulesArray[k].Status = status
		schedulesArray[k].Sales.Agent.CountryCode = code.CountryCode.ChineseNameSimple[v.Sales.Agent.CountryCode]
	}

	return nil

}







