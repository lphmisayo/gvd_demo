package system

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model/common"
	"gin_vue_blog_server/model/response"
	"gin_vue_blog_server/model/system"
	"gin_vue_blog_server/model/system/request"
	"gin_vue_blog_server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func (b *BaseApi) Register(c *gin.Context) {
	var register utils.Register
	err := c.ShouldBindJSON(&register)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	errType, stVer := register.RegistStVertify()
	if !stVer {
		var errMsg string
		if errType == utils.ParamError {
			errMsg = "用户名或者密码输入不能为空！"
		} else if errType == utils.StructError {
			errMsg = "入参结构出错，请检查参数！"
		}
		response.FailWithMessage(errMsg, c)
	}
	user := common.User{
		UserName: register.Username,
		NickName: register.Nickname,
		Email:    register.Email,
		PassWord: register.Password,
	}
	err = userService.Register(&user)
	if err != nil {
		response.FailWithDataAndMsg(&user, "注册失败！错误原因："+err.Error(), c)
		return
	}
	response.OkWithDataAndMsg(&user, "注册成功！", c)
}

func (b *BaseApi) LoginDefault(c *gin.Context) {
	var login utils.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	errType, stVer := login.LoginStVertify(true)
	if !stVer {
		var errMsg string
		if errType == utils.ParamError {
			errMsg = "用户名或者密码输入不能为空！"
		} else if errType == utils.StructError {
			errMsg = "入参结构出错，请检查参数！"
		}
		response.FailWithMessage(errMsg, c)
	}
	u := common.User{UserName: login.Username, PassWord: login.Password}
	user, err := userService.LoginDefault(&u)
	if err != nil {
		global.Logger.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		response.FailWithDataAndMsg(user, "登陆失败! 用户名不存在或者密码错误! 错误原因："+err.Error(), c)
		return
	}
	response.OkWithDataAndMsg(user, "登录成功！", c)
}

// Login 后续优化
func (b *BaseApi) Login(c *gin.Context) {
	var login utils.Login
	err := c.ShouldBindJSON(&login)
	key := c.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	errType, stVer := login.LoginStVertify(false)
	if !stVer {
		var errMsg string
		if errType == utils.ParamError {
			errMsg = "用户名或者密码输入不能为空！"
		} else if errType == utils.StructError {
			errMsg = "入参结构出错，请检查参数！"
		} else if errType == utils.CaptchaError {
			errMsg = "验证码输入不能为空！"
		}
		response.FailWithMessage(errMsg, c)
	}
	openCaptcha := global.Config.Captcha.OpenCaptcha               //是否开启防爆次数
	openCaptchaTimeOut := global.Config.Captcha.OpenCaptchaTimeOut //缓存超过时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	var oc bool = openCaptcha == 0 || openCaptcha < utils.InterfaceToInt(v)
	if !oc || utils.Store.Verify(login.CaptchaId, login.Captcha, true) {
		u := &system.SysUser{Username: login.Username, Password: login.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.Logger.Error("登录失败！用户名不存在或者密码错误！", zap.Error(err))
			//验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.Logger.Error("登陆失败！用户被禁止登录！")
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("登陆失败！用户被禁止登录！", c)
			return
		}
		b.TokenNext(c, *user)
		return
	}
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{TokenKey: []byte(global.Config.JWT.TokenKey)}
	claims := j.CreateClaims(request.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.Logger.Error("获取Token失败！", zap.Error(err))
		response.FailWithMessage("获取Token失败！", c)
		return
	}
	if !global.Config.System.UseMultipoint {
		response.OkWithDataAndMsg(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
}
