package database

import (
	"TikTok/model"
	"TikTok/utility"
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
)

func Init() {
	dbUser, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbUser.AutoMigrate(&model.User{})
	dbVideo, err := gorm.Open(sqlite.Open("video.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbVideo.AutoMigrate(&model.Video{})
} //以用户唯一id为名称创建该用户关注列表的数据库文件
func CreateFollowList(ListId int64) {
	dbFollowList, err := gorm.Open(sqlite.Open(strconv.Itoa(int(ListId))), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbFollowList.AutoMigrate(&struct {
		ID int64 `json:"id"`
	}{})
}
func SaveUser(user *model.User) error {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	db.Create(user)
	return nil
}
func VerifyLogin(login *model.LoginRequest) model.User {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic(err)
		return model.User{}
	}
	var userFromDB model.User
	db.Where("name=?", login.Username).First(&userFromDB)
	login.Password, err = utility.HashPassword(login.Password)
	if err != nil {
		errors.New("unable to hash password")
		return model.User{}
	}
	if login.Password != userFromDB.Password {
		return model.User{}
	}
	return userFromDB
}
func QueryUser(req *model.UserRequest) model.User {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic(err)
		return model.User{}
	}
	var userFromDB model.User
	db.Where("ID=?", req.UserID).First(&userFromDB)
	if (userFromDB == model.User{}) {
		return model.User{}
	}
	return userFromDB
}
