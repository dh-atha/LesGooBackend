package mysql

import (
	"fmt"
	"lesgoobackend/config"
	chatsData "lesgoobackend/feature/chats/data"
	groupUsersData "lesgoobackend/feature/group_users/data"
	groupData "lesgoobackend/feature/groups/data"
	"lesgoobackend/feature/users/data"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username, cfg.Password, cfg.Address, cfg.Port, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func MigrateData(db *gorm.DB) {
	db.AutoMigrate(data.User{}, groupData.Group{}, groupUsersData.Group_User{}, chatsData.Chat{})
}
