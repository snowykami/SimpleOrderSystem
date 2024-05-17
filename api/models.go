package api

// gorm
import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID        int64  `json:"id" gorm:"primaryKey"`
	Username  string `json:"nickname"`
	Password  string `json:"password"` // Password is hashed
	Token     string `json:"token"`    // JWT token
	SecretKey string `json:"secret"`   // SecretKey for JWT
	IsAdmin   bool   `json:"is_admin"` // Is admin
}

type Order struct {
	ID         int64 `json:"id" gorm:"primaryKey"`
	UserID     int64 `json:"user_id"`     // 创建者id
	CreatedAt  int64 `json:"created_at"`  // Unix timestamp
	ItemID     int64 `json:"item_id"`     // 商品id
	ActivityID int64 `json:"activity_id"` // 活动id
}

type Item struct {
	ID    int64   `json:"id" gorm:"primaryKey"` // 商品id，固定不变，用于储存库存
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Store string  `json:"store"` // 商店名称,店主username
	Stock int64   `json:"stock"` // 库存
}

type Activity struct {
	// 优惠活动
	ID          int64   `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Discount    float64 `json:"discount"`   // 折扣
	StartTime   int64   `json:"start_time"` // 起始时间戳
	EndTime     int64   `json:"end_time"`   // 结束时间戳
	Description string  `json:"description"`
}

func init() {
	// 迁移数据库
	DB, _ = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	err := DB.AutoMigrate(&User{}, &Order{}, &Item{}, &Activity{}, &Order{})
	if err != nil {
		Logger.Error("auto migrate error: " + err.Error())
		return
	}
}
