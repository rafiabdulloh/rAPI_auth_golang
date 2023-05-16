package models

type Users struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(300)" json:"username"`
	Password   string `gorm:"type:text" json:"password"`
}
