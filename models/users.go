package models

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

func GetUser(username string) User {
	var user User
	DB.Where(&User{Username: username}).First(&user)
	return user
}

func GetUserById(id uint) User {
	var user User
	DB.Where(&User{ID: id}).First(&user)
	return user
}

func GetUsers() []User {
	var users []User
	DB.Find(&users)
	return users
}
