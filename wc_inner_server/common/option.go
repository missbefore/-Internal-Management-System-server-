package common

import (
	"strings"
)

func DealGetOptions(args map[string]map[string]string) (string, []interface{}, []interface{})  {

	var whereName string
	var valueFind []interface{}
	var connectString string

	var joinWhere []interface{}
	if args != nil {
		for k, v := range args{
			if len(v["data"]) == 0 ||
				v["data"] == "%%" ||
				v["data"] == "()"{
				delete(args, k)
			} else {

				if _, ok:= v["son_deal"]; ok {
					whereSonData := strings.Split(v["data"], ",")
					for s := 0;s < len(whereSonData);s++ {
						valueFind = append(valueFind, whereSonData[s])
					}
				}
				whereField := v["field"]
				whereFieldArray := strings.Split(whereField, "~")
				connectString = " " + v["other"] + " "
				whereFieldArray = removeDuplicatesAndEmpty(whereFieldArray)
				var n = 0
				for x:=0;x<len(whereFieldArray);x ++ {
					if _, ok:= v["son_deal"]; ok {
						sonArray := strings.Split(whereFieldArray[x], "|")
						if len(sonArray) > 1 {
							for xSon:= 0;xSon<len(sonArray) ;xSon++ {
								whereName += sonArray[xSon] + " ? " + v["son_deal_connect"]
							}
						}
						if _,ok := v["deal_connect"];ok {
							selfContent := " " + v["deal_connect"] + " "
							whereName = strings.TrimRight(whereName, v["son_deal_connect"]) + selfContent
						} else {
							whereName = strings.TrimRight(whereName, v["son_deal_connect"])
						}
					} else {
						whereName += whereFieldArray[x] + "?"
						isTimeArray := strings.Split(v["data"], ",")
						if len(isTimeArray) > 1 {
							valueFind = append(valueFind, isTimeArray[n])
							n++
						} else {
							valueFind = append(valueFind, v["data"])
						}
					}

					if (x+1) == len(whereFieldArray) {
						whereName = strings.TrimRight(whereName," " +v["deal_connect"]+" ")
					}
					if _, ok := v["son_deal_connect"];!ok || (x+1) == len(whereFieldArray){
						whereName += connectString
					}
				}

				//联表条件join处理 field_other : name like
				if _, ok:= v["field_other"];ok {
					joinWhere = append(joinWhere, v["field_other"] + " ? ")
					joinWhere = append(joinWhere, v["data"])
				}
			}
		}

		whereName = strings.TrimRight(whereName, connectString)
	}

	return  whereName, valueFind, joinWhere
}


func removeDuplicatesAndEmpty(a []string) (ret []string) {
	aLen := len(a)
	for i := 0; i < aLen; i++{
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}
