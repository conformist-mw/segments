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

type UserUpdateForm struct {
	Email       string `form:"email"`
	Username    string `form:"username"`
	FirstName   string `form:"first_name"`
	LastName    string `form:"last_name"`
	IsSuperuser string `form:"is_superuser"`
	IsActive    string `form:"is_active"`
}

func UpdateUser(user User, form UserUpdateForm) User {
	if form.Email != "" {
		user.Email = form.Email
	}
	if form.Username != "" {
		user.Username = form.Username
	}
	if form.FirstName != "" {
		user.FirstName = form.FirstName
	}
	if form.LastName != "" {
		user.LastName = form.LastName
	}
	if form.IsSuperuser == "on" {
		user.IsSuperuser = true
	} else {
		user.IsSuperuser = false
	}
	if form.IsActive == "on" {
		user.IsActive = true
	} else {
		user.IsActive = false
	}
	DB.Save(&user)
	return user
}
