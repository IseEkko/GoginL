package global

import (
	ut "github.com/go-playground/universal-translator"
	"singo/conf"
)

var (
	Trans ut.Translator

	ServerConfig *conf.ServerConfig  = &conf.ServerConfig{}
	UserConfig   *conf.UserSrvConfig = &conf.UserSrvConfig{}

)
