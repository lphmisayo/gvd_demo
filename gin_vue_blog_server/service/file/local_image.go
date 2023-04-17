package file

import (
	"errors"
	"gin_vue_blog_server/global"
	"gin_vue_blog_server/model"
	ctype "gin_vue_blog_server/model/cType"
	"gin_vue_blog_server/model/common"
	"gin_vue_blog_server/model/request"
	"gin_vue_blog_server/model/response"
	"gin_vue_blog_server/utils"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"strings"
)

type LocalImageService struct {
}

var (
	whiteImageList = []string{
		"jpg", "png", "jpeg", "webp", "ai", "svg",
	}
)

func (LocalImageService) ImageUpload(file *multipart.FileHeader, filePath string) (res response.FileResponse) {

	fileName := file.Filename

	//
	nameList := strings.Split(filePath, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	//判断文件类型
	if !utils.CheckStrInList(suffix, whiteImageList) {
		res = utils.InitFileResponse(file.Filename, -1, response.ImageUploadTypeErr, response.ImageMsg, nil)
		return res
	}
	//
	fileSize := float64(file.Size) / float64(1024*1024) //文件大小 MB
	if fileSize > float64(global.Config.Upload.ImageSize) {
		res = utils.InitFileResponse(file.Filename, -1, response.ImageUploadTypeErr, response.ImageMsg, nil)
		return res
	}

	//生成文件MD5
	fileObject, err := file.Open()
	if err != nil {
		global.Logger.Error(err.Error(), zap.Error(err))
		res = utils.InitFileResponse(file.Filename, -1, response.ErrorMsg, response.ErrMsg, err)
		return res
	}
	byteData, err := io.ReadAll(fileObject)
	if err != nil {
		global.Logger.Error(err.Error(), zap.Error(err))
		res = utils.InitFileResponse(file.Filename, -1, response.ErrorMsg, response.ErrMsg, err)
		return res
	}
	md5Str := utils.InitMd5Str(string(byteData))

	var image common.Image
	dbErr := global.DB.Take(&image, "hash = ?", md5Str).Error
	if dbErr == nil {
		res = utils.InitFileResponse(file.Filename, -1, response.FileUploadMd5Err, response.FileMsg, err)
		return res
	}

	image = common.Image{
		Name:      fileName,
		Hash:      md5Str,
		Path:      filePath,
		ImageType: ctype.Local,
	}
	image.Model = model.Model{CreateAt: utils.GetNowTimeHasFormat()}
	global.DB.Create(&image)

	return utils.InitFileResponse(fileName, 0, response.ImageUploadSuccess, response.SucMsg, nil)
}

func (LocalImageService) ImageUpdate(fileReq request.FileRequest) error {
	if global.DB == nil {
		return errors.New("数据库未初始化！")
	}
	var updateImage common.Image
	updateImage.Name = fileReq.Name
	updateImage.ID = uint(fileReq.ID)
	err := global.DB.Model(common.Image{}).Where("id = ?", updateImage.ID).Update("name", updateImage.Name).Error
	if err != nil {
		return err
	}
	return nil
}

func (LocalImageService) CheckImageIsExist(fileReq request.FileRequest) (bool, int, error) {
	if global.DB == nil {
		return false, -1, nil
	}
	err := global.DB.Model(common.Image{}).Where("id = ?", fileReq.ID).Error
	if err != nil {
		return false, -2, err
	}
	return true, 0, nil
}

func (LocalImageService) ImageDelete(fileReq request.FileRequest) error {
	if global.DB == nil {
		return errors.New("数据库未初始化！")
	}
	err := global.DB.Delete(&common.Image{}, fileReq.ID).Error
	if err != nil {
		return err
	}
	return nil
}
