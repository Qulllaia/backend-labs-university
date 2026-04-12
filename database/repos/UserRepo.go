package repos

import (
	"gorm.io/gorm"
)

type User struct {
	User_ID uint   `gorm:"primaryKey"`
	Login   string `gorm:"size:100"`
}

func (User) TableName() string {
	return "user"
}

type UserRepo struct {
	DB *gorm.DB
}

func (ur *UserRepo) CreateUser(login string) *User {
	user := &User{Login: login}
	ur.DB.Create(user)
	return user
}

func (ur *UserRepo) DeleteUser(id string, login string) *User {
	var user User
	ur.DB.Delete(map[string]any{
		"user_id": id,
	})
	return &user
}

func (ur *UserRepo) UpdateUser(id string, login string) *User {
	user := &User{Login: login}
	ur.DB.Where(map[string]any{
		"user_id": id,
	}).Update("login", user)
	return user
}

func (ur *UserRepo) GetUser(id string) *User {
	var user User
	ur.DB.Where(map[string]any{
		"user_id": id,
	}).First(&user)
	return &user
}
