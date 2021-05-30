package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	CreatedOn string `gorm:"column:createdOn"`
	IsEnabled bool   `gorm:"column:isEnabled"`
	IsDeleted bool   `gorm:"column:isDeleted"`
	UpdatedOn string `gorm:"column:updatedOn"`
}


func (usr *User) ComparePassword(givenPassword string) bool{
	passwordMatches:=false

	err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(givenPassword))

	if err==nil{
		passwordMatches=true
	}
	return passwordMatches
}
