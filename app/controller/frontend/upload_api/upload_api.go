package upload_api

import (
	"eat/global"
	"eat/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"path"
	"time"
)

type UploadData struct {
	Url  string `json:"url"`
	Alt  string `json:"alt"`
	Href string `json:"href"`
}

func (UploadApi) Image(c *gin.Context) {
	// 获取文件名单文件
	file, _ := c.FormFile("uploaded-image")
	// 获取后缀名
	extString := path.Ext(file.Filename)
	// 使用UUID生成一个唯一的文件名，并拼接上扩展名
	fileName := uuid.NewString() + extString
	// 使用当前日期（年/月/日）生成文件夹名
	folderName := time.Now().Format("2006/01/02")
	// 定义文件的完整路径
	dst := "resource/upload/images/" + folderName + "/" + fileName
	// 上传文件至指定的完整文件路径
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		global.Logger.Println(err)
	}
	// 返回数据
	data := UploadData{
		Url: "http://" + global.Config.System.Addr() + "/" + dst,
	}
	response.SuccessWithData(data, c)
}
