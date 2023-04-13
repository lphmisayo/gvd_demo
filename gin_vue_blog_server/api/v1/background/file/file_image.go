package file

import (
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model"
	"gin_vue_blog_server/model/common"
	"gin_vue_blog_server/model/response"
	"gin_vue_blog_server/service/file"
	"github.com/gin-gonic/gin"
)

func (*FileApi) ImageListView(c *gin.Context) {
	var pageInit model.PageInfo
	err := c.ShouldBindJSON(&pageInit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	fileList, count, err := file.GetFileList(common.Image{}, file.Option{
		PageInfo: pageInit,
		Debug:    global.MysqlLogConfig.Debug,
	})
	if err != nil {
		response.FailWithMessage("图片列表获取失败！错误原因："+err.Error(), c)
	}

	response.OkWithList(fileList, count, "图片列表获取成功！", c)
}
