package entity

import (
	"math"
	"strings"
	"time"
	"wc_inner_server/common/code"
	"wc_inner_server/model/table"
)

func DealMerchantList(merchants *[]model.Merchants) error  {
	//开始处理数据
	merchantsArray :=  *merchants
	for k, v := range merchantsArray {
		begin, _ := time.Parse(Time, strings.Split(v.BeginDate, "T")[0])
		end, _ := time.Parse(Time, strings.Split(v.EndDate, "T")[0])



		if v.Status == "ACTIVE"  {
			unixEnd := end.Unix()
			nowTime := time.Now().Unix()
			distance := math.Ceil(float64(unixEnd-nowTime))/86400

			if  distance > 0 {
				merchantsArray[k].Status = "ON_EXPIRE"
			}  else {
				merchantsArray[k].Status = "EXPIRE"
			}
		}

		merchantsArray[k].BeginDate = begin.Format(Time) + " 至 " + end.Format(Time)
		merchantsArray[k].CountryCode = code.CountryCode.ChineseNameSimple[v.CountryCode]
	}

	return nil
}
