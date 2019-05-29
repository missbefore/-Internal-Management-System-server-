package model

type Agents struct {
	Id        		     uint	 `gorm:"primary_key;column:id"`
	No            		 string  `gorm:"column:no"`
	Account        		 string  `gorm:"column:account"`
	Name          		 string  `gorm:"column:name"`
	CountryCode   		 string  `gorm:"column:country_code"`
	ChargeRate      	 string  `gorm:"column:charge_rate"`
	AlipaycnflRate 		 string  `gorm:"column:alipaycnfl_rate"`
	AlipaycnflHkrate     string  `gorm:"column:alipaycnfl_hk_rate"`
	AlipaycnRate         string  `gorm:"column:alipaycn_rate"`
	AlipaycnHkRate       string  `gorm:"column:alipaycn_hk_rate"`
	WechatpaycnRate      string  `gorm:"column:wechatpaycn_rate"`
	WechatpaycnflRate    string  `gorm:"column:wechatpaycnfl_rate"`
	SupayalicnflRate     string  `gorm:"column:supayalicnfl_rate"`
	SupayalicnRate       string `gorm:"column:supayalicn_rate"`
	BeginDate            string `gorm:"column:begin_date"`
	EndDate              string `gorm:"column:begin_date"`
	ContactName          string  `gorm:"column:contact_name"`
	ContactEmail         string  `gorm:"column:contact_email"`
	ContactPhone         string  `gorm:"column:contact_phone"`
	Status               string  `gorm:"column:status"`
	SettleCurrency       string  `gorm:"column:settle_currency"`
	SettleAmount         string  `gorm:"column:settle_amount"`
	RegisterCountry      string  `gorm:"column:register_country"`
	Merchants            []Merchants `gorm:"ForeignKey:AgentNo;AssociationForeignKey:No"`
}

