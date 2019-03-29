package models

import "golang.org/x/crypto/bcrypt"

// AdminLogin 로그인 => 반환 (bool, 에러 혹은 이름)
func AdminLogin(studentNumber string, password string) (bool, string) {
	user := User{}
	err := db.Table("users").Where("student_number = ?", studentNumber).First(&user).Error
	if err != nil {
		return false, "없는 아이디입니다"
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+SALT))
	if err != nil {
		return false, "비밀번호가 틀렸습니다"
	}

	return true, user.Name
}
