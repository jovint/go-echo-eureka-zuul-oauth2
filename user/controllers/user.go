package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"school/models"
	"github.com/google/uuid"
	"log"
	"school/utils"
)

func getConnection()  (*gorm.DB){
	db, err := gorm.Open("mysql", "root:root@/user")
	fmt.Println(err)
		 
	 return db
}


func HandlePost(c echo.Context) error {
	var user models.User
	var role models.Roles
	var userDetails models.UserDetails
	user.UserId = uuid.New().String()
	role.RoleId = "1"
	userDetails.DetailsId = uuid.New().String()
	userDetails.FullName= "Jovin Thariyath"
	userDetails.UserId = user.UserId
	user.UserDetails = userDetails
	if err := c.Bind(&user); err != nil {
		return err
	}
	db:= getConnection()
	db.AutoMigrate(&user, &role, &userDetails)

	db.Create(&user)
	return c.String(http.StatusOK, "User created!")
}

func HandleGet(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	username, err := utils.VerifyToken(token)
	if username != "false"{
		fmt.Println(username)
	}

    if err != nil {
        log.Fatal(err)
    }


	db:= getConnection()
	if db != nil {
		fmt.Println(db)
	}
	return c.String(http.StatusOK, "Hello, World!")
}
