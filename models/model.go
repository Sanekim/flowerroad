package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	// SQLConnectionString : MySQL Connection String
	SQLConnectionString = "*"
	// SALT : SALT
	SALT = "*"
)

// Database Connection
var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", SQLConnectionString)
	if err != nil {
		fmt.Println("db have error")
	}
}

// User models
type User struct {
	StudentNumber string `gorm:"type:VARCHAR(6); primary_key" json:"studentNumber"`
	Name          string `gorm:"type:VARCHAR(15)" json:"name"`
	Hr            string `gorm:"type:VARCHAR(10)" json:"hr"`
	Password      string `gorm:"type:VARCHAR(100)" json:"password"`
}
