package models

type Data struct {
	ID      int    `json:"id" gorm:"primary_key:auto_increment;"`
	Message string `json:"message" form:"message" gorm:"type: varchar(255)"`
}
