package models

import (
	/* "time" */
	"gorm.io/gorm"
)

/* type Fact struct {
	gorm.Model
	Question string `jason:"question" gorm:"text;not null;default:null"`
	Answer   string `jason:"answer"   gorm:"text;not null;default:null" `
} */

type Customer struct {
	/* ID        uint `jason:"id" gorm:"primaryKey"`
	CreatedAt time.Time */
	gorm.Model
	CustomerName
}

type CustomerName struct {
	/* ID        uint `jason:"id" gorm:"primaryKey"`
	CreatedAt time.Time */
	FirstName string `jason:"first_name"`
	LastName  string `jason:"last_name"`
}

type Product struct {
	/* ID           uint `jason:"id" gorm:"primaryKey"`
	CreatedAt    time.Time */
	gorm.Model
	Name  string `jason:"name"`
	Price uint   `jason:"price"`
	//Quantity uint   `jason:"quantity"`
}

type Order struct {
	/* ID         uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time */
	gorm.Model
	CustomerID uint     `json:"customer_id"`
	Customer   Customer `gorm:"foreignKey: CustomerID"`
	ProductID  uint     `json:"product_id"`
	Product    Product  `json:"product" gorm:"foreignKey: ProductID"`
}
