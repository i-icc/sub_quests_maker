package database

import (
	"fmt"
	"os"

    "github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// SQLConnect DB接続
func GormConnect() (database *gorm.DB) {
    // パスワード等を.envファイルから読み取る
    // program > go > .env
    err := godotenv.Load("../../.env")
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("env読み取り成功")
    }

    DBMS := "mysql"                   // MySQL
    PROTOCOL := "tcp(localhost:3306)" // db:3306
    DBNAME := os.Getenv("MYSQL_DATABASE")    // テーブル名
    USER := os.Getenv("MYSQL_USER")      // MySQLユーザー名
    PASS := os.Getenv("MYSQL_USER_PASSWORD")  // パスワード

    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("DB接続成功")
    }
    return db
}