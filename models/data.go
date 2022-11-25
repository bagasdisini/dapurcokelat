package models

type Data struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment;"`
	DataUser string `json:"dataUser" form:"dataUser" gorm:"type: varchar(255)"`
	Result   string `json:"result" gorm:"type: varchar(255)"`
}
