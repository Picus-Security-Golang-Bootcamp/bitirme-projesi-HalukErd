package entity

type User struct {
	Item
	Email    string `gorm:"uniqueIndex"`
	Password string
	Roles    []string `gorm:"type:text"`
}
