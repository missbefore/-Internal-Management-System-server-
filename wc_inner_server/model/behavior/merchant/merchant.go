package merchant

import (
	"wc_inner_server/common/code"
	"wc_inner_server/model/entity"
	"wc_inner_server/model/table"
	"wc_inner_server/model/tool"
)

var db = tool.DB
var table = "merchants"
func GetDataList(whereNames string, whereValues []interface{}, joinWhere []interface{},
merchants *[]model.Merchants, ) error {
	
	if err := db.Select("no, agent_no, name, short_name, country_code, begin_date, end_date, contact_name, status").
		Preload("Agent", joinWhere...).
		Where(whereNames, whereValues...).Find(&merchants).Error;err != nil {
		return err
	}
	if errEntity := entity.DealMerchantList(merchants);errEntity != nil {
		return errEntity
	}

	return nil
}

func GetOptionsData() (map[string]interface{}, error) {

	var results = make(map[string]interface{})

	var country []string
	if err := db.Table(table).Pluck("DISTINCT(country_code)", &country).Error; err != nil {
		return nil, err
	}

	var resultCountry []model.OptionCountry
	for _, v := range country  {
		resultCountry = append(resultCountry, model.OptionCountry{Cn:code.CountryCode.ChineseNameSimple[v], En:v})
	}
	results["country"] = resultCountry


	var resultsAgent []model.OptionResult

	db.Raw("SELECT distinct(agent_no), agents.name " +
		"FROM merchants INNER JOIN agents ON merchants.agent_no=agents.no").Scan(&resultsAgent)
	results["agents"] = resultsAgent

	return results, nil
}