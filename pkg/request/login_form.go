package request

type LoginForm struct {
	User string	`json:"user" binding:"required" gorm:"primaryKey"`
	Password string	`json:"password" binding:"required"`
}
