package routers

import (
	"gin_vue_blog_server/routers/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp RouterGroup

/*func BackServer() *http.Server {
	backPort := global.Config.Server.BackPort
	log.Printf("后台服务启动于 %s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      BackRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
*/
