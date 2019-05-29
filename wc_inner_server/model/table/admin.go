package model

type Admins struct {
	Id      		     uint	 `gorm:"primary_key;column:id"`
	No            		 string  `gorm:"column:no"`
    Password             string  `gorm:"column:password"`
	Mobile               string  `gorm:"column:mobile"`
	RealName             string  `gorm:"column:realname"`
	LastLogin            string  `gorm:"column:last_login"`
	Signature            string  `gorm:"column:signature"`
	Role                 string  `gorm:"column:role"`
}

