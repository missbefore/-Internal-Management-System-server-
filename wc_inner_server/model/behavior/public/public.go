package public

import (
	"wc_inner_server/model/table"
	"wc_inner_server/model/tool"
)

var db = tool.DB
func GetPostUserInfo(whereName string, whereValue string, admin *model.Admins) error {

	if err := db.Select("mobile, password, no, signature, role, realname").
		Where(whereName+"=?", whereValue).Find(&admin).Error; err != nil {
			return err
	}

	return nil
}

func UpdateUserInfo(field string, info string, no string, admin *model.Admins) error {
	if err := db.Model(&admin).Update(field,  info).Where("no = ?", no).Error; err != nil {
		return err
	}

	return nil
}
