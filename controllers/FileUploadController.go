package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"mime/multipart"
	"os"
	"io"
	"Blog/models"
	"path"
)

type FileUploadController struct {
	BaseController
}

//设置方法
func (this *FileUploadController) URLMapping(){
	this.Mapping("Upload",this.Upload)
}

//数据返回结构体设计
type ImageResult struct {
	Message string    `json:"message"`
	Success int        `json:"success"`
	Url  string         `json:"url"`
}
//文件上传
// @router /file/upload?:guid [post]
func (this *FileUploadController) Upload() {
	var files []*multipart.FileHeader  = this.Ctx.Request.MultipartForm.File["editormd-image-file"]
	file , err := files[0].Open();
	defer file.Close()
	if err != nil {
		beego.Error(err)
		return
	}
	imgUrl := this.SaveFile(file,files[0]);
	if imgUrl != "saveImageErr" && imgUrl != "createImageErr"  {
		this.SetData(&ImageResult{Message:"上传成功",Success:1,Url:imgUrl})
	}else{
		this.SetData(&ImageResult{Message:"上传失败",Success:0,Url:""})
	}
	this.ServeJSON()
}

//上传图片,返回值为图片id
func (this *FileUploadController) SaveFile(file multipart.File, fheader *multipart.FileHeader) (imgUrl string) {
	defer file.Close()
	fileName := fheader.Filename
	fileName = models.GetGuid()+path.Base(fileName)
	fW, err := os.Create("./download/" + fileName)
	if err != nil {
		fmt.Println("文件创建失败")
		imgUrl = "createImageErr"
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		imgUrl = "saveImageErr"
		fmt.Println("文件保存失败")
		return
	}
	website := this.GetConfigParameter("website")
	imgUrl = website+"/download/" + fileName
	return
}

