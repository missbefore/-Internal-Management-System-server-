package model

type Sales struct {
	Id      		     uint	 `gorm:"primary_key;column:id"`
	No            		 string  `gorm:"column:no"`
	AgentNo              string  `gorm:"column:agent_no"`
	SaleNo               string  `gorm:"column:sale_no"`
	Name                 string  `gorm:"column:name"`
	Agent				 Agents  `gorm:"ForeignKey:No;AssociationForeignKey:AgentNo"`
}

