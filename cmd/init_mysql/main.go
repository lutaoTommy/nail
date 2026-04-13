package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"nail/config"
	"nail/handler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	now := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println("1. 开始读取配置文件")
	/*加载端口配置*/
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("2. 开始连接数据库: %s\n", config.GetMysqlUrl())

	db := getMysqlConn(config.GetMysqlUrl())

	fmt.Println("3. 开始初始化用户")

	// 按你的要求：初始化用户直接用 Save 覆盖现有记录
	adminPwd, err := handler.HashPassword("fndroid")
	if err != nil {
		panic(err)
	}
	admin := handler.User{
		Phone:    "admin",
		Token:    "admin",
		Passwd:   adminPwd,
		UserId:   "admin",
		Language: "en-US",
		Nickname: "Admin",
		Status:   1,
		Email:    "admin@example.com",
	}
	if err := db.Save(&admin).Error; err != nil {
		panic(err)
	}

	touristPwd, err := handler.HashPassword("fndroid")
	if err != nil {
		panic(err)
	}
	tourist := handler.User{
		Phone:    "tourist",
		Token:    "tourist",
		Passwd:   touristPwd,
		UserId:   "tourist",
		Language: "en-US",
		Nickname: "Tourist",
		Status:   1,
	}
	if err := db.Save(&tourist).Error; err != nil {
		panic(err)
	}

	fmt.Println("4. 开始初始化色系")

	colorDescData, err := os.ReadFile("config/color_desc.json")
	if err != nil {
		panic(err)
	}
	var colorDescs []handler.ColorDesc
	if err = json.Unmarshal(colorDescData, &colorDescs); err != nil {
		panic(err)
	}
	for i := range colorDescs {
		colorDescs[i].CreateTime = now
		if err = db.Save(&colorDescs[i]).Error; err != nil {
			panic(err)
		}
	}

	fmt.Println("5. 开始初始化颜色")

	colorsData, err := os.ReadFile("config/colors.json")
	if err != nil {
		panic(err)
	}
	var colors []handler.Color
	if err = json.Unmarshal(colorsData, &colors); err != nil {
		panic(err)
	}
	for i := range colors {
		colors[i].Count = 61
		colors[i].CreateTime = now
		if err = db.Save(&colors[i]).Error; err != nil {
			panic(err)
		}
	}
}

/*连接数据库*/
func getMysqlConn(database string) *gorm.DB {
	dsn := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s", database)
	mysqlSession, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	err = mysqlSession.Set("gorm:table_options", "DEFAULT CHARSET=utf8").AutoMigrate(&handler.CirclePost{},
		&handler.PostImage{}, &handler.Comment{}, &handler.Like{}, &handler.Collect{}, &handler.User{},
		&handler.Color{}, &handler.ColorRecog{}, &handler.PostInfo{}, &handler.Device{}, &handler.Avatar{},
		&handler.Suggest{}, &handler.SuggestImage{}, &handler.SensitiveWord{}, &handler.Follow{},
		&handler.OtaVersion{}, &handler.ColorDesc{}, &handler.PostColor{}, &handler.ColorFavorite{}, &handler.LutData{},
		&handler.ApkVersion{},
	)
	if err != nil {
		panic(err)
	}
	return mysqlSession
}
