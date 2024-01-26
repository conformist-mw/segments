package models

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

type User struct {
	ID          uint   `gorm:"primarykey;not null" json:"id"`
	Email       string `gorm:"type:varchar(254);not null" json:"email"`
	Username    string `gorm:"type:varchar(150);not null;unique" json:"username"`
	FirstName   string `gorm:"type:varchar(150);not null" json:"first_name"`
	LastName    string `gorm:"type:varchar(150);not null" json:"last_name"`
	Password    string `gorm:"type:varchar(255);not null" json:"password"`
	DateJoined  string `gorm:"type:datetime;not null" json:"date_joined"`
	LastLogin   string `gorm:"type:datetime" json:"last_login"`
	IsSuperuser bool   `gorm:"type:bool;not null;default:false" json:"is_superuser"`
	IsStaff     bool   `gorm:"type:bool;not null;default:false" json:"is_staff"`
	IsActive    bool   `gorm:"type:bool;not null;default:true" json:"is_active"`
}

func (User) TableName() string {
	return "auth_user"
}

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func CheckLogin(userForm LoginForm) (User, error) {
	var user User
	DB.Where(&User{Username: userForm.Username}).First(&user)

	// Split the Django password into its components
	parts := strings.Split(user.Password, "$")
	if len(parts) != 4 {
		return User{}, errors.New("Формат пароля неверный")
	}

	iterations, err := strconv.Atoi(parts[1])
	if err != nil {
		return User{}, errors.New("Формат пароля неверный")
	}

	salt := parts[2]
	djangoHash := parts[3]

	// Hash the provided password using the same PBKDF2 algorithm
	hash := pbkdf2.Key([]byte(userForm.Password), []byte(salt), iterations, sha256.Size, sha256.New)
	goHash := base64.StdEncoding.EncodeToString(hash)

	// Compare the hashed password with the Django password
	if goHash == djangoHash {
		return user, nil
	}

	return User{}, errors.New("Неправильное имя пользователя или пароль")
}
