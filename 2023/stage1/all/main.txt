//package main
//
//import (
//	"gorm.io/gorm"
//	"time"
//)
//
//type AqQuery struct {
//	ID          uint         `gorm:"primaryKey"`
//	UserID      uint64       `gorm:"column:user_id"`
//	TaskQuery   *TaskQuery   `gorm:"embedded"`
//	HourQuery   *HourQuery   `gorm:"embedded"`
//	ReportQuery *ReportQuery `gorm:"embedded"`
//}
//
//type TaskQuery struct {
//	ProName  uint   `json:"proName"`
//	Status   uint   `json:"status"`
//	Versions []uint `json:"versions"`
//	UserIds  []uint `json:"userIds"`
//}
//
//type HourQuery struct {
//	StartCreatedAt time.Time `json:"startCreatedAt"`
//	EndCreatedAt   time.Time `json:"endCreatedAt"`
//	ProName        uint      `json:"proName"`
//	TaskName       string    `json:"taskName"`
//}
//
//type ReportQuery struct {
//	StartCreatedAt time.Time `json:"startCreatedAt"`
//	EndCreatedAt   time.Time `json:"endCreatedAt"`
//	ProName        []uint    `json:"proName"`
//}
//
//func main() {
//	db, err := gorm.Open(mysql.Open("dsn"), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	err = db.AutoMigrate(&AqQuery{})
//	if err != nil {
//		panic("failed to migrate table")
//	}
//}
