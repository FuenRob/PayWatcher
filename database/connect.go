package database

import (
	config "PayWatcher/config"
	"PayWatcher/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	configDB := config.GetConfigDataBase()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configDB.User, configDB.Pass, configDB.Host, configDB.Port, configDB.Name)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database because: %s!", err.Error()))
	}

	fmt.Println("Database connected!")
	// Migrate the schema
	fmt.Println("Migrating the schema...")
	DB.AutoMigrate(model.User{})
	DB.AutoMigrate(model.Payment{})
	DB.AutoMigrate(model.Category{})
	fmt.Println("Schema migrated!")
}
