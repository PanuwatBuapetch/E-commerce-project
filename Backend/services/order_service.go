package services

import (
	"ecommerce-backend/configs"
)

func CreateOrder(totalAmount float64) (int, error) {
	// บันทึกลงตาราง orders
	query := "INSERT INTO orders (user_id, total_amount, status, created_at) VALUES (?, ?, 'pending', NOW())"

	// Hardcode user_id = 1 ไปก่อน (ตามที่คุยกัน)
	result, err := configs.DB.Exec(query, 1, totalAmount)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return int(id), nil
}
