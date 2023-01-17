package controller

import (
	"TikTok/database"
	"TikTok/model"
	"TikTok/utility"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func Register(c *gin.Context) {
	var reg model.RegisterRequest
	var user model.User
	var err error
	c.ShouldBindJSON(&reg)
	//limit username and password in 0-32
	if len(reg.Username) < 0 || len(reg.Username) > 32 || len(reg.Password) < 0 || len(reg.Password) > 32 {
		c.JSON(200, &model.RegisterResponse{StatusCode: -1, StatusMSG: "wrong format of username or password", UserID: 0, Token: ""})
	}
	user.Name = reg.Username
	user.Password, err = utility.HashPassword(reg.Password)
	if err != nil {
		errors.New("unable to crypto password")
		return
	}
	user.CreateTime = time.Now().UnixNano()
	user.FollowerCount = 0
	user.FollowCount = 0
	//user.IsFollow = true
	uuid, err := uuid.NewUUID()
	user.ID = int64(uuid.ID())
	//save to database
	err = database.SaveUser(&user)
	if err != nil {
		panic(err)
	}
	token, err := utility.GenToken(&user)
	if err != nil {
		errors.New("unable to gen token")
		c.JSON(200, model.RegisterResponse{
			StatusCode: -1,
			StatusMSG:  "unable to gen token",
			UserID:     0,
			Token:      "",
		})
		return
	}
	c.JSON(200, model.RegisterResponse{
		StatusCode: 0,
		StatusMSG:  "success",
		UserID:     user.ID,
		Token:      token,
	})
	database.CreateFollowList(user.ID)
}
func Login(c *gin.Context) {
	var login model.LoginRequest
	c.ShouldBindJSON(&login)
	user := database.VerifyLogin(&login)
	if (user == model.User{}) {
		c.JSON(200, model.LoginResponse{
			StatusCode: -1,
			StatusMSG:  "username or password is not correct",
			UserID:     0,
			Token:      "",
		})
		return
	}
	token, err := utility.GenToken(&user)
	if err != nil {
		errors.New("unable to gen token")
		c.JSON(200, model.RegisterResponse{
			StatusCode: -1,
			StatusMSG:  "unable to gen token",
			UserID:     0,
			Token:      "",
		})
		return
	}
	c.JSON(200, model.LoginResponse{
		StatusCode: 0,
		StatusMSG:  "success",
		UserID:     user.ID,
		Token:      token,
	})

}
func User(c *gin.Context) {
	var req model.UserRequest
	c.ShouldBindJSON(&req)
	//handle token
	user := database.QueryUser(&req)
	if (user == model.User{}) {
		c.JSON(200, model.UserResponse{
			StatusCode: -1,
			StatusMSG:  "wrong user id",
			User:       model.User{},
		})
		return
	}
	c.JSON(200, model.UserResponse{
		StatusCode: 0,
		StatusMSG:  "success",
		User:       user,
	})
}
