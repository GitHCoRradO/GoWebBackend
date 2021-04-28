package dao

import (
	"fmt"
	"github.com/githcorrado/fake-admagic/pkg/database"
	"github.com/githcorrado/fake-admagic/pkg/request"
	"reflect"
)

var db = database.Db

func InitTables() {
	createTableIfNotExist(request.LoginForm{})
}

func createTableIfNotExist(modelType interface{}) {
	val := reflect.Indirect(reflect.ValueOf(modelType))
	modelName := val.Type().Name()
	err := db.AutoMigrate(modelType)
	fmt.Printf("Migrating Table of %s ...\n", modelName)
	if err != nil {
		fmt.Errorf("AutoMigrate failed with error: %v", err)
	}
}