package services

import (
	"database/sql"
	"ecommerce-backend/configs"
	"ecommerce-backend/models"
	"errors"
	"log"
)

// 1. READ ALL: ดึงข้อมูลสินค้าทั้งหมด
func GetAllProducts() ([]models.Product, error) {
	rows, err := configs.DB.Query("SELECT id, name, price, image_url, category, rating FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		// Scan ค่าเข้าตัวแปร (ต้องเรียงลำดับให้ตรงกับ SELECT)
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.ImageURL, &p.Category, &p.Rating)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		products = append(products, p)
	}

	return products, nil
}

// 2. CREATE: เพิ่มสินค้าใหม่
func CreateProduct(p models.Product) (models.Product, error) {
	query := "INSERT INTO products (name, price, image_url, category, rating, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())"

	result, err := configs.DB.Exec(query, p.Name, p.Price, p.ImageURL, p.Category, p.Rating)
	if err != nil {
		return p, err
	}

	// ดึง ID ที่เพิ่งสร้างมาใส่กลับใน Struct
	id, _ := result.LastInsertId()
	p.ID = int(id)

	return p, nil
}

// 3. READ ONE: ดึงสินค้าตาม ID
func GetProductByID(id string) (models.Product, error) {
	var p models.Product
	query := "SELECT id, name, price, image_url, category, rating FROM products WHERE id = ?"

	// QueryRow ใช้สำหรับดึงข้อมูลแค่แถวเดียว
	row := configs.DB.QueryRow(query, id)

	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.ImageURL, &p.Category, &p.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, errors.New("product not found")
		}
		return p, err
	}

	return p, nil
}

// 4. UPDATE: แก้ไขข้อมูลสินค้า
func UpdateProduct(id string, p models.Product) (models.Product, error) {
	query := "UPDATE products SET name=?, price=?, image_url=?, category=?, rating=?, updated_at=NOW() WHERE id=?"

	result, err := configs.DB.Exec(query, p.Name, p.Price, p.ImageURL, p.Category, p.Rating, id)
	if err != nil {
		return p, err
	}

	// เช็คว่ามีแถวไหนถูกอัปเดตไหม
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return p, errors.New("product not found or no changes made")
	}

	// (Optional) แปลง id string เป็น int เพื่อใส่กลับใน struct (ถ้าจำเป็น)
	// p.ID = ...

	return p, nil
}

// 5. DELETE: ลบสินค้า
func DeleteProduct(id string) error {
	query := "DELETE FROM products WHERE id=?"

	result, err := configs.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}
