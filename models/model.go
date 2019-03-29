package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

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

// Login 로그인 => 반환 (bool, 에러 혹은 이름)
func Login(studentNumber string, password string) (bool, string) {
	user := User{}
	err := db.Table("users").Where("student_number = ?", studentNumber).First(&user).Error
	if err != nil {
		return false, "없는 학번입니다"
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+SALT))
	if err != nil {
		return false, "비밀번호가 틀렸습니다"
	}

	return true, user.Name
}
