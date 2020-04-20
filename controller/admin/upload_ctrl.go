package admin

import (
	"errors"
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminUploadCtrl struct {
	UploadType  string               `val:"${upload.type}"`
	LocalUpload *service.LocalUpload `ioc:"auto"`
	UpyunUpload *service.UpyunUpload `ioc:"auto"`
}

func (s *AdminUploadCtrl) Upload(ctx *ctrl.HTTPContext) error {
	var err error
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.ResultErr(err)
	}
	var result *model.Upload
	if s.UploadType == "local" {
		result, err = s.LocalUpload.Upload(file)
	} else if s.UploadType == "upyun" {
		result, err = s.UpyunUpload.Upload(file)
	} else {
		err = errors.New("Unsupport uploader type " + s.UploadType)
	}
	if err != nil {
		return ctx.ResultErr(err)
	}
	ctx.Put("path", result.URL)
	return ctx.ResultOK()
}

type base64File struct {
	Base64 string `json:"base64"`
}

func (s *AdminUploadCtrl) UploadBase64(ctx *ctrl.HTTPContext) error {
	var err error
	file := new(base64File)
	if err := ctx.Bind(file); err != nil {
		return ctx.ResultErrMsg("上传失败：" + err.Error())
	}
	var result *model.Upload
	if s.UploadType == "local" {
		result, err = s.LocalUpload.UploadBase64(file.Base64)
	} else if s.UploadType == "upyun" {
		result, err = s.UpyunUpload.UploadBase64(file.Base64)
	} else {
		err = errors.New("Unsupport uploader type " + s.UploadType)
	}
	if err != nil {
		return ctx.ResultErr(err)
	}
	ctx.Put("path", result.URL)
	return ctx.ResultOK()
}

func (s *AdminUploadCtrl) MarkdownUpload(ctx *ctrl.HTTPContext) error {
	form, err := ctx.Context.MultipartForm()
	//file, err := ctx.FormFile("file[]")
	if err != nil {
		return ctx.ResultErr(err)
	}
	var results []*model.Upload
	for _, files := range form.File {
		for _, file := range files {
			var result *model.Upload
			if s.UploadType == "local" {
				result, err = s.LocalUpload.Upload(file)
			} else if s.UploadType == "upyun" {
				result, err = s.UpyunUpload.Upload(file)
			} else {
				err = errors.New("Unsupport uploader type " + s.UploadType)
			}
			if err != nil {
				return ctx.ResultErr(err)
			}
			results = append(results, result)
		}
	}
	ctx.Put("path", results)
	return ctx.ResultOK()
}
