package models

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/mail"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

func ValidateColor(colorSlug, colorTypeSlug string) error {
	var color Color
	var colorType ColorType
	DB.Where("slug = ?", colorSlug).First(&color)
	DB.Where("slug = ?", colorTypeSlug).First(&colorType)
	if color.ID == 0 {
		return errors.New("color not found")
	}
	if colorType.ID == 0 {
		return errors.New("color type not found")
	}
	if color.TypeID != colorType.ID {
		return errors.New("color type mismatch")
	}
	return nil
}

func ValidateRack(companySlug, sectionSlug string, rackId int) error {
	rack := GetRack(companySlug, sectionSlug, rackId)
	if rack.ID == 0 {
		return errors.New("rack not found")
	}
	return nil
}

func ValidateRemoveForm(RemoveForm RemoveForm) error {
	isDefect := RemoveForm.Defect == "on"
	if RemoveForm.OrderNumber == "" && !isDefect {
		return errors.New("Для удаления отрезка нужно указать или номер заказа или указать дефект")
	}
	if RemoveForm.OrderNumber == "" {
		if isDefect && RemoveForm.Description == "" {
			return errors.New("При указании дефекта нужно указать описание")
		}
	}
	if RemoveForm.OrderNumber != "" && GetOrderNumber(RemoveForm.OrderNumber).ID != 0 {
		// TODO: order number should be unique within company
		return errors.New("Такой номер заказа уже есть в базе")
	}
	if RemoveForm.OrderNumber == "" && !isDefect {
		return errors.New("no order number or defect")
	}
	return nil
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

func ValidateCreateUserForm(form CreateUserForm) error {
	match, _ := regexp.MatchString("^[a-z.]{2,}$", form.Username)
	if !match {
		return errors.New("Username must only contain ASCII lowercase letters and dots, and be at least 2 characters long")
	}

	if len(form.Password) < 5 {
		return errors.New("Password must be at least 5 characters long")
	}

	if form.Email != "" {
		_, err := mail.ParseAddress(form.Email)
		if err != nil {
			return errors.New("Email is not valid")
		}
	}
	return nil
}
