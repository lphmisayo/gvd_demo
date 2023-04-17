package file

import (
	"errors"
	"fmt"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model"
	"gin_vue_blog_server/model/common"
	"gin_vue_blog_server/model/request"
	"gin_vue_blog_server/model/response"
	"gin_vue_blog_server/service/file"
	"gin_vue_blog_server/utils"
	"github.com/gin-gonic/gin"
	"os"
	"path"
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

func (*FileApi) ImageUpdateView(c *gin.Context) {
	var fileReq request.FileRequest
	err := c.ShouldBindJSON(&fileReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	isExist, errType, dbErr := localImageService.CheckImageIsExist(fileReq)
	if !isExist {
		if errType == -1 {
			response.FailWithMessage(errors.New("数据库未初始化").Error(), c)
			return
		} else if errType == -2 {
			response.FailWithMessage(dbErr.Error(), c)
			return
		} else {
			response.FailWithMessage(errors.New("未知错误，无法检查图片是否存在！").Error(), c)
		}
	}

	err = localImageService.ImageUpdate(fileReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功！", c)
}

func (*FileApi) ImageUploadView(c *gin.Context) {
	images, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := images.File["images"]
	if !ok {
		response.FailWithCode(response.FailNotExistError, c)
		return
	}
	dir, err := os.ReadDir(global.Config.Upload.Path)
	if err != nil {
		err = os.MkdirAll(global.Config.Upload.Path, os.ModePerm)
		fmt.Println(dir, err)
	}
	UploadFiles := make([]response.FileResponse, 0)
	for _, file := range fileList {
		filePath := path.Join(global.Config.Upload.StorePath, file.Filename)

		fileResponse := localImageService.ImageUpload(file, filePath)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Logger.Error(err.Error())
			fileResponse = utils.InitFileResponse(file.Filename, -1, response.ErrorMsg, response.ErrMsg, err)
			continue
		}
		UploadFiles = append(UploadFiles, fileResponse)
	}

	response.OkWithData(UploadFiles, c)
}

func (*FileApi) ImageDeleteView(c *gin.Context) {
	var fileReq request.FileRequest
	err := c.ShouldBindJSON(&fileReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = localImageService.ImageDelete(fileReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功！", c)
}
