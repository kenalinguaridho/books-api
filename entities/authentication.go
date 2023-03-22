package entities

type User struct {
	Id       uint   `gorm:"type:int" json:"id"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Username string `gorm:"type:varchar(100)" json:"username"`
	Password string `gorm:"type:varchar(100)" json:"password"`
}
