package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils"
	"io"
	"path"
	"path/filepath"
	"strings"
)

var (
	// WhiteImageList 图片上传的白名单
	WhiteImageList = []string{
		".jpg", ".png", ".jpeg", ".ico", ".tiff", ".gif", ".svg", ".webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

func (ImagesApi) ImagesUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("image files does not exist", c)
		return
	}

	var resList []FileUploadResponse

	for _, file := range fileList {
		// 判断图片白名单
		fileExt := strings.ToLower(filepath.Ext(file.Filename))
		_, ok := utils.InList(fileExt, WhiteImageList)
		if !ok {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "the file extension is illegal",
			})
			continue
		}

		filePath := path.Join(global.Config.Upload.Path, file.Filename)
		// 判断文件大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("file size: %.2f MD, is over the limitation: %d MB", size, global.Config.Upload.Size),
			})
			continue
		}

		// MD5 加密处理
		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		md5String := utils.Md5(byteData)
		//fmt.Println(md5String)

		// 去数据库查看图片是否已经存在
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "hash = ?", md5String).Error
		fmt.Println(err)
		if err == nil {
			// 找到相同的图片文件
			resList = append(resList, FileUploadResponse{
				FileName:  bannerModel.Path,
				IsSuccess: true,
				Msg:       "Image is already existed!",
			})
			continue
		}

		// 保存文件至指定目录
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err.Error())
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: true,
				Msg:       err.Error(),
			})
			continue
		}
		resList = append(resList, FileUploadResponse{
			FileName:  file.Filename,
			IsSuccess: true,
			Msg:       "File upload successfully!",
		})
		// 将图片存入数据库
		global.DB.Create(&models.BannerModel{
			Path: filePath,
			Hash: md5String,
			Name: file.Filename,
		})
	}
	res.OkWithData(resList, c)
}
