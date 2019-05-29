package model

type Merchants struct {
	ID        		uint	   `gorm:"primary_key;column:id"`
	No 		        string     `gorm:"column:no"`
	Account         string     `gorm:"column:account"`
	Password        string     `gorm:"column:password"`
	AgentNo         string     `gorm:"column:agent_no"`
	WxMerchantId    string     `gorm:"column:wx_merchant_id"`
	Name            string     `gorm:"column:name"`
	ShortName       string     `gorm:"column:short_name"`
	CountryCode     string     `gorm:"column:country_code"`
	ChargeRate      string     `gorm:"column:charge_rate"`
	AllowCurrency   string     `gorm:"column:allow_currency"`
	SettleCurrency  string     `gorm:"column:settle_currency"`
	SettledBalance  string     `gorm:"column:settled_balance"`
	TradeBalance    string     `gorm:"column:trade_balance"`
	MinBalance      string     `gorm:"column:min_balance"`
	RegisterCountry string     `gorm:"column:register_country"`
	BeginDate       string     `gorm:"column:begin_date"`
	EndDate         string     `gorm:"column:end_date"`
	Type            string     `gorm:"column:type"`
	ContactName     string     `gorm:"column:contact_name"`
	ContactEmail    string     `gorm:"column:contact_email"`
	ContactPhone    string     `gorm:"column:contact_phone"`
	Status          string     `gorm:"column:status"`
	ServicePhone    string     `gorm:"column:service_phone"`
	BeginTime       string     `gorm:"column:begin_time"`
	EndTime         string     `gorm:"column:end_time"`
	Agent           Agents     `gorm:"ForeignKey:No;AssociationForeignKey:AgentNo"`
}


type OptionCountry struct {
	Cn string
	En string
}


type OptionResult struct {
	AgentNo    string
	Name  string
}