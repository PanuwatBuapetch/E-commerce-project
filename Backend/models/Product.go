package models

// Product สำหรับตารางสินค้า
type Product struct {
	ID       int     `json:"id"` // ใช้ int ปกติ (Database ส่วนใหญ่ ID เป็น int)
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"` // ชื่อใน Database: image_url
	Category string  `json:"category"`
	Rating   float64 `json:"rating"`
}
