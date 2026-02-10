package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // สำคัญมาก: ต้องมีบรรทัดนี้เพื่อโหลด Driver
)

// เปลี่ยน Type จาก *gorm.DB เป็น *sql.DB เพื่อให้ใช้ .Query และ .Exec แบบ 2 ค่าได้
var DB *sql.DB

func ConnectDB() {
	var err error
	// ใส่ DSN ของคุณ (User:Pass@tcp(127.0.0.1:3306)/DbName)
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce_db?parseTime=true"

	// ใช้ sql.Open แทน gorm.Open
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// เช็คการเชื่อมต่อ
	if err = DB.Ping(); err != nil {
		log.Fatal("Database unreachable: ", err)
	}

	fmt.Println("Connected to Database via Standard SQL!")
}
