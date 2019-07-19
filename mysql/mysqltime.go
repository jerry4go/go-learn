package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    //   parseTime=true  把获取到的string类型的时间转换为时间类型
	//   loc=Local       把默认的UTC +0 时区转换为 本地时区
	db, err := sql.Open("mysql",
		"root:a123456@tcp(x.x.x.x:3306)/gintestdb?parseTime=true&loc=Local")
	var myTime time.Time
	rows, err := db.Query("SELECT current_timestamp()")
	fmt.Println(time.Now())
	if rows.Next() {
		if err = rows.Scan(&myTime); err != nil {
			panic(err)
		}
	}
	fmt.Println(myTime)
}
