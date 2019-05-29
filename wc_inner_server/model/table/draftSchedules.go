package model

type DraftSchedules struct {
	No      			string `gorm:"column:no"`
	DraftMerchantsNo    string `gorm:"column:draft_merchants_no"`
	Status              string `gorm:"column:status"`
	Pointer             string `gorm:"column:pointer"`
	Handler		        string `gorm:"column:handler"`
	SalesNo             string `gorm:"column:sales_no"`
	Modified            string `gorm:"column:modified"`
	Sales               Sales  `gorm:"ForeignKey:No;AssociationForeignKey:SalesNo"`
	DraftMerchants      DraftMerchants `gorm:"ForeignKey:No;AssociationForeignKey:DraftMerchantsNo"`
}