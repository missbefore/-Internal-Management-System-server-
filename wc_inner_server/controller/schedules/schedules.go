package schedules

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"wc_inner_server/common"
	"wc_inner_server/common/code"
	"wc_inner_server/model/behavior/schedules"
	"wc_inner_server/model/table"
)

var sendError = common.SendError
var optionDeal = common.DealGetOptions
var draftSchedules  []model.DraftSchedules

func List(c *gin.Context)  {

	option := make(map[string]map[string]string)

	statusArray := strings.Split(c.Query("status"), ",")

	for k, v := range statusArray {
		option["status"+ strconv.Itoa(k)] = map[string]string{
			"other":"OR",
		}
		if vData := strings.Split(v, "_"); len(vData) > 1 {
			option["status"+ strconv.Itoa(k)]["data"] = vData[0] + "," + vData[1]
			option["status"+ strconv.Itoa(k)]["field"] = "status = | handler = ~"
			option["status"+ strconv.Itoa(k)]["son_deal"] = "ok"
			option["status"+ strconv.Itoa(k)]["son_deal_connect"] = "AND"
		} else {
			option["status"+ strconv.Itoa(k)]["data"] = v
			option["status"+ strconv.Itoa(k)]["field"] = "status= ~"
		}
	}

	sqlOptionNames, sqlValOptions, _ := optionDeal(option)
	sqlOptionNames += " AND pointer = ?"
	sqlValOptions = append(sqlValOptions, 1)
	err := schedules.GetDataList(sqlOptionNames, sqlValOptions, &draftSchedules)
	if err != nil {
		sendError("查询进件失败", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": code.SuccessCode.SUCCESS,
		"msg": "success",
		"data": gin.H{
			"schedules": draftSchedules,
		},
	})
}
