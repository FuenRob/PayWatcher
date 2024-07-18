package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
