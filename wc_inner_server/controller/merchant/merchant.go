package merchant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"wc_inner_server/common"
	"wc_inner_server/common/code"
	"wc_inner_server/model/behavior/merchant"
	"wc_inner_server/model/table"
)

var sendError = common.SendError
var optionDeal = common.DealGetOptions
var merchants  []model.Merchants
var onLineAdmin model.UserAdmin
func List(c *gin.Context)  {

	fmt.Println(onLineAdmin.No)
	option := make(map[string]map[string]string)

	option["country"] = map[string]string{
		"data": c.Query("country"),
		"field":"country_code =",
		"other": "AND",
	}
	option["agentNo"] = map[string]string{
		"data":c.Query("agents"),
		"field":"agent_no = ",
		"other": "AND",
	}
	option["merchantInfo"] = map[string]string{
		"data": "%" + c.Query("merchant_info") + "%",
		"other":"OR",
		"field":"name LIKE~short_name LIKE~no = ",
	}
	option["startTime"] = map[string]string{
		"data": c.Query("time_range"),
		"field":"begin_date >= ~end_date <=",
		"other": "AND",
	}
	option["wxMerchantsId"] = map[string]string{
		"data":  c.Query("wx_merchant_id"),
		"field":"wx_merchant_id = ",
		"other": "AND",
	}


	//deal the Disorderly data structure
	optionStatusArray := strings.Split(c.Query("Status"), ",")
	var statusWhereString string
	var statusWhereValue string
	for i:=0;i<len(optionStatusArray) ;i++ {
		switch optionStatusArray[i] {
		case "ACTIVE":
			statusWhereString += "status = | end_date > ~"
			aMonthLater := time.Now().AddDate(0,1,0).Format("2006-01-02")
			statusWhereValue += "ACTIVE, " + aMonthLater + ","
			break
		case "ON_EXPIRE":
			statusWhereString += "status = | end_date > | end_date < ~"
			aMonthLaterGreater := time.Now().Format("2006-01-02")
			aMonthLaterLess := time.Now().AddDate(0,1,0).Format("2006-01-02")
			statusWhereValue += "ACTIVE, " + aMonthLaterGreater + "," +  aMonthLaterLess + ","
			break
		case "EXPIRE":
			statusWhereString += "status = | end_date < ~"
			itsNow := time.Now().Format("2006-01-02")
			statusWhereValue += "ACTIVE, " + itsNow + ","
			break
		case "DENIED":
			statusWhereString += "status = ~"
			statusWhereValue += "DENIED,"
			break
		case "DELETE":
			statusWhereString += "status = ~"
			statusWhereValue += "DELETE,"
			break
		}
	}

	option["status"] = map[string]string{
		"data": strings.TrimRight(statusWhereValue, ","),
		"field": strings.Trim(statusWhereString, "~"),
		"son_deal" : "ok",
		"other": "AND",
		"son_deal_connect" : "AND",
		"deal_connect" : "OR",
	}

	sqlOptionNames, sqlValOptions, joinWhere := optionDeal(option)

	err := merchant.GetDataList(sqlOptionNames, sqlValOptions,joinWhere, &merchants)
	if err != nil {
		sendError("查询商户列表出错", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": code.SuccessCode.SUCCESS,
		"msg": "success",
		"data": gin.H{
			"merchants": merchants,
		},
	})
}

func GetOptionsData(c *gin.Context)  {

	results, errQuery := merchant.GetOptionsData()
	if errQuery != nil {
		fmt.Println(errQuery.Error())
		sendError("查询出错", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": code.SuccessCode.SUCCESS,
		"msg": "success",
		"data": results,
	})
}
