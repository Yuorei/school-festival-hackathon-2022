package operateDb

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// データベースに入れるtable
type User struct {
	gorm.Model
	Uid string
}
type Rent_list struct {
	gorm.Model
	Uuid string `json:"uuid"`
	User_id     int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image_url   string `json:"Image_url"`
	Deadline    string `json:"Deadline"`
}
type Lend_list struct {
	gorm.Model
	Uuid string `json:"uuid"`
	User_id     int    `gorm:"foreignKey:User_id"`
	Name        string `json:"name"`
	Description string `json:"weight"`
	Image_url   string `json:"at"`
	Deadline    string `json:"Deadline"`
}

func Init() error {
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	TZ := os.Getenv("TZ")

	dsn := "host=" + POSTGRES_HOST + " user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " port=5432 sslmode=disable TimeZone=" + TZ
	//データベースを開ける
	//CloseはPostgreSQLの使用で不要
	dbCon, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//ローカル変数のdbCOnをグローバル変数のdbに入れる
	db = dbCon
	//作成するtableの構造体を引数に入れてテーブルを作る
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Rent_list{})
	db.AutoMigrate(&Lend_list{})
	if err != nil {
		return err
	}
	return nil
}

func GetConnect() *gorm.DB {
	return db
}
