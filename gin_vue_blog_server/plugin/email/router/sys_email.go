package router

import (
	"github.com/gin-gonic/gin"
)

type EmailRouter struct{}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	//emailRouter := Router.Use()
	//EmailApi := api.ApiGroupApp.EmailApi.EmailTest
	//SendEmail := api.ApiGroupApp.EmailApi.SendEmail
	//{
	//	emailRouter.POST("emailTest", EmailApi)  // 发送测试邮件
	//	emailRouter.POST("sendEmail", SendEmail) // 发送邮件
	//}
}
