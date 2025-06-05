package models

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	DB = db
	return db
}

func TestCheckLogin_Success(t *testing.T) {
	setupTestDB(t)
	password := "secret"
	hash := GeneratePasswordHash(password)
	user := User{Username: "john", Password: hash}
	DB.Create(&user)

	u, err := CheckLogin(LoginForm{Username: "john", Password: password})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.ID == 0 || u.Username != user.Username {
		t.Fatalf("returned user not correct")
	}
}

func TestCheckLogin_NonexistentUser(t *testing.T) {
	setupTestDB(t)
	_, err := CheckLogin(LoginForm{Username: "nouser", Password: "pass"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	expected := "Формат пароля неверный"
	if err.Error() != expected {
		t.Fatalf("expected error %q, got %v", expected, err)
	}
}
