package wechat_api

import (
	"eat/app/model"
	"eat/global"
	"eat/utils/jwt"
	"eat/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
)

type SessionParams struct {
	Code          string `json:"code"`
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
}

func (WechatApi) User(c *gin.Context) {
	var params SessionParams
	err := c.ShouldBind(&params)
	if err != nil {
		fmt.Println(err)
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	global.Logger.Info(params)

	wc := wechat.NewWechat()

	// 这里本地内存保存access_token，也可选择redis或者自定cache
	// memory := cache.NewMemory()
	// wc.SetCache(memory)

	// 获取小程序配置信息
	redisConf := global.Config.Redis
	redisOpts := &cache.RedisOpts{
		Host: redisConf.Addr(),
	}
	// 设置缓存
	redisCache := cache.NewRedis(c, redisOpts)
	wc.SetCache(redisCache)

	// 设置小程序信息
	wxappConf := global.Config.Wxapp
	cfg := &config.Config{
		AppID:     wxappConf.AppID,
		AppSecret: wxappConf.AppSecret,
		// Token:     "xxx",
		// EncodingAESKey: "xxxx",
		// Cache: memory,
	}
	mini := wc.GetMiniProgram(cfg)
	auth := mini.GetAuth()
	// 根据code获取sessionKey
	session, _ := auth.Code2Session(params.Code)
	global.Logger.Info(session)
	// 解析用户信息
	encryptor := mini.GetEncryptor()
	user, _ := encryptor.Decrypt(session.SessionKey, params.EncryptedData, params.Iv)
	global.Logger.Info(user)

	// 根据open_id查找用户信息
	var userModel model.UserModel
	result := global.DB.Where("open_id = ?", user.OpenID).Select("user_id,nickname,avatar").First(&userModel)
	// 找不到则新增用户信息
	if result.RowsAffected == 0 {
		userModel.UnionId = user.UnionID
		userModel.OpenId = user.OpenID
		userModel.Nickname = user.NickName
		userModel.Avatar = user.AvatarURL
		global.DB.Save(&userModel)
	}

	// 生成 jwt token
	jwtPayLoad := jwt.JwtPayLoad{
		UserId:   userModel.UserId,
		Nickname: userModel.Nickname,
		Avatar:   userModel.Avatar,
	}
	token, _ := jwt.GenToken(jwtPayLoad)

	response.SuccessWithData(UserInfo{
		UserId:   userModel.UserId,
		Nickname: userModel.Nickname,
		Avatar:   userModel.Avatar,
		Token:    token,
	}, c)

}

type UserInfo struct {
	UserId   uint   `json:"user_id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Token    string `json:"token,omitempty"`
}
