package main

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"collect_manage/Libraries/Databases"
	"collect_manage/Models"
)

var ROOT_PATH string

var DB_GORM_SQLITE *gorm.DB

func init() {
	var err error
	ROOT_PATH, err = os.Getwd()
	if err != nil {
		panic("failed to get path")
	}
}

func main() {
	var err error
	DB_GORM_SQLITE, err = Databases.NewDB(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	DB_GORM_SQLITE.AutoMigrate(&Models.User{})

	// Create
	// eamil := "zk001@qq.com"
	// birthday := time.Now()
	// DB_GORM_SQLITE.Create(&Models.User{
	// 	ID:           1,
	// 	Name:         "zk001",
	// 	Email:        &eamil,
	// 	Age:          18,
	// 	Birthday:     &birthday,
	// 	MemberNumber: sql.NullString{"0002", true},
	// })

	// Read
	var user Models.User
	DB_GORM_SQLITE.First(&user, 1)           // 根据整型主键查找
	DB_GORM_SQLITE.First(&user, "id = ?", 1) // 查找 code 字段值为 D42 的记录

	fmt.Println(user)

	// Update - 将 product 的 price 更新为 200
	DB_GORM_SQLITE.Model(&user).Update("Age", 20)
	// // Update - 更新多个字段
	DB_GORM_SQLITE.Model(&user).Updates(Models.User{Age: 21, MemberNumber: sql.NullString{"0003", true}}) // 仅更新非零值字段
	// DB_GORM_SQLITE.Model(&user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// DB_GORM_SQLITE.Delete(&product, 1)

	// fmt.Println(ROOT_PATH)
}
