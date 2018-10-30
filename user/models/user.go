package models

import (
	"github.com/jinzhu/gorm"
)

//User type
type User struct {
	gorm.Model
	UserId      string      `gorm:"column:user_id;primary_key:true";`
	UserName    string      `gorm:"column:user_name"`
	Password    string      `gorm:"column:password"`
	UserDetails UserDetails  `gorm:"foreignkey:UserId"`
}

//UserDetails type
type UserDetails struct {
	gorm.Model
	DetailsId string `gorm:"column:details_id;primary_key:true"`
	FullName  string `gorm:"column:full_name"`
	Phone     string `gorm:"column:phone"`
	Address   string `gorm:"column:address"`
	UserId    string `gorm:"column:user_id"`
}

type Roles struct {
	gorm.Model
	RoleId string `gorm:"column:role_id;primary_key:true"`
	RoleName  string `gorm:"column:role_name"`
}

