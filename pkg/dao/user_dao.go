package dao

import "github.com/githcorrado/fake-admagic/pkg/request"

func GetAllUserNames() ([]request.LoginForm, error) {
	var loginForms []request.LoginForm
	res := db.Find(&loginForms)
	return loginForms, res.Error
}

func CreateUser(form request.LoginForm) (request.LoginForm, error) {
	res := db.Create(&form)
	return form, res.Error
}