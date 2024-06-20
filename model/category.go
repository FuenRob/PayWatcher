package model

type Category struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Priority  uint   `json:"priority"`
	Recurrent bool   `json:"recurrent"`
	Notify    bool   `json:"notify"`
}
