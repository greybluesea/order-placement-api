package models

import (
	/* "time" */
	"gorm.io/gorm"
)

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

type OrderRequest struct {
	CustomerID uint //`json:"CustomerID"`
	ProductID  uint //`json:"ProductID"`
}

type Order struct {
	/* ID         uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time */
	gorm.Model
	CustomerID uint     `json:"CustomerID"`
	Customer   Customer `gorm:"foreignKey:CustomerID"`
	ProductID  uint     `json:"ProductID"`
	Product    Product  `gorm:"foreignKey:ProductID"`
}
