package models

type User struct {
	Id      string `gorm:"column:id" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Email   string `gorm:"column:email" json:"email"`
	Mobile  string `gorm:"column:mobile" json:"mobile"`
	Address string `gorm:"column:address" json:"address"`
}

func (b *User) String() string {
	return "users"
}
