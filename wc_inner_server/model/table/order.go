package model

type Order struct {
	No          string      `gorm:"column:no"`
	MerchantNo  string      `gorm:"column:merchant_no"`
	Currency    string      `gorm:"column:currency"`
	Amount      string      `gorm:"column:amount"`
	PayCurrency string      `gorm:"column:pay_currency"`
	PayAmount   string      `gorm:"column:pay_amount"`
	ExchangeRate string     `gorm:"column:exchange_rate"`
	SettleRate   string     `gorm:"column:settle_rate"`
	DeviceInfo   string     `gorm:"column:device_info"`
	ProductInfo  string     `gorm:"column:product_info"`
	PrepayDetail string     `gorm:"column:prepay_detail"`
	PayDetail    string     `gorm:"column:pay_detail"`
	ProductId    string     `gorm:"column:product_id"`
	LimitPay     string     `gorm:"column:limit_pay"`
	TradeType    string     `gorm:"column:trade_type"`
	AgentOrderId string     `gorm:"column:agent_order_id"`
	Status       string     `gorm:"column:status"`
	PayMethod    string     `gorm:"column:pay_method"`
	IsCredited   string     `gorm:"column:is_credited"`
	RefundStatus string     `gorm:"column:refund_status"`
	RefundableAmount string `gorm:"column:refundable_amount"`
	Channel      string     `gorm:"column:channel"`
	ChargeRate   string     `gorm:"column:charge_rate"`
	ChargeAmount string     `gorm:"column:charge_amount"`
	SettleCurrency string   `gorm:"column:settle_currency"`
	Subject      string     `gorm:"column:subject"`
	TransactionTime string  `gorm:"column:transaction_time"`
	IsSurcharge  string     `gorm:"column:is_surcharge"`
	SurchargeAmount string  `gorm:"column:surcharge_amount"`
	SurchargeCurrency string `gorm:"column:surcharge_currency"`
	SurchargeRate  string   `gorm:"column:surcharge_rate"`
	Merchant      Merchants `gorm:"ForeignKey:No;AssociationForeignKey:MerchantNo"`
}
