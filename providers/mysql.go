package providers

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"kafka2mysql/configs"
)

func NewMySql(config configs.DatabaseConfig) *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				config.User, config.Password, config.Host, config.Port, config.Name)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
