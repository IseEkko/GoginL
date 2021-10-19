package service

import (
	"github.com/dgrijalva/jwt-go"
	"singo/middleware"
	"singo/model"
	"singo/serializer"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", err)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	j := middleware.NewJWT()
	claims := model.CustomClaims{
		Id:             user.ID,
		UserName:         user.UserName,
		PasswordDigest:   user.PasswordDigest,
		Nickname :        user.Nickname,
		Status :         user.Status,
		Avatar  :         user.Avatar,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),            //签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24, //30天过期
			Issuer:    "user-web",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return serializer.ParamErr("生成token失败", nil)
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.LoginResponse{
         Data: serializer.BuildUser(user),
         Token: token,
		},
		Msg: "登录成功",
	}
}
