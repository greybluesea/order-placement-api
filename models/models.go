package models

import (
	/* "time" */
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerName
}

type CustomerName struct {
	FirstName string `jason:"first_name"`
	LastName  string `jason:"last_name"`
}

type Product struct {
	gorm.Model
	Name  string `jason:"name"`
	Price uint   `jason:"price"`
}

type OrderRequest struct {
	CustomerID uint //`json:"CustomerID"`
	ProductID  uint //`json:"ProductID"`
}

type Order struct {
	gorm.Model
	CustomerID uint     `json:"CustomerID"`
	Customer   Customer `gorm:"foreignKey:CustomerID"`
	ProductID  uint     `json:"ProductID"`
	Product    Product  `gorm:"foreignKey:ProductID"`
}
