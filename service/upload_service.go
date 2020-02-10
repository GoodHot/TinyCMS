package service

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/GoodHot/TinyCMS/common/files"
	"github.com/GoodHot/TinyCMS/common/strs"
	"github.com/GoodHot/TinyCMS/common/times"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	uuid "github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

type LocalUpload struct {
	StaticDir string   `val:"${server.static}"`
	UploadDir string   `val:"${upload.local_dir}"`
	MaxSize   int      `val:"${upload.max_size}"`
	ORM       *orm.ORM `ioc:"auto"`
}

func (s *LocalUpload) Upload(file *multipart.FileHeader) (*model.Upload, error) {
	// check size
	if file.Size > int64(s.MaxSize) {
		return nil, errors.New("Upload file must smaller than " + strconv.Itoa(s.MaxSize) + "k")
	}
	tmpFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	currDate := times.TimeStr("20060102")
	uploadDir := strs.Fmt("%s/%s/%s", s.StaticDir, s.UploadDir, currDate)
	err = s.createUploadDir(uploadDir)
	if err != nil {
		return nil, err
	}
	fileType, _ := getFileType(file.Header.Get("Content-Type"))
	fileName := s.fileName(file.Filename)
	URL := strs.Fmt("/%s/%s", uploadDir, fileName)
	Path := strs.Fmt("%s/%s", uploadDir, fileName)
	upload := &model.Upload{
		Type: fileType,
		URL:  URL,
		Path: Path,
	}
	// Destination
	dst, err := os.Create(Path)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, tmpFile); err != nil {
		return nil, err
	}
	return upload, s.ORM.DB.Create(upload).Error
}

func (s *LocalUpload) UploadBase64(base64Data string) (*model.Upload, error) {
	head := base64Data[strings.Index(base64Data, ":")+1 : strings.Index(base64Data, ";")]
	fileType, suffix := getFileType(head)
	code := base64Data[strings.Index(base64Data, ",") + 1:]
	decodeBytes, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return nil, err
	}
	if len(decodeBytes) > s.MaxSize {
		return nil, errors.New("Upload file must smaller than " + strconv.Itoa(s.MaxSize) + "k")
	}
	currDate := times.TimeStr("20060102")
	uploadDir := strs.Fmt("%s/%s/%s", s.StaticDir, s.UploadDir, currDate)
	err = s.createUploadDir(uploadDir)
	if err != nil {
		return nil, err
	}
	fileName := s.fileName("tmp" + suffix)
	URL := strs.Fmt("/%s/%s", uploadDir, fileName)
	Path := strs.Fmt("%s/%s", uploadDir, fileName)
	upload := &model.Upload{
		Type: fileType,
		URL:  URL,
		Path: Path,
	}
	// Destination
	dst, err := os.Create(Path)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, bytes.NewBuffer(decodeBytes)); err != nil {
		return nil, err
	}
	return upload, s.ORM.DB.Create(upload).Error
}

func (s *LocalUpload) fileName(fileName string) string {
	suffix := fileName[strings.LastIndex(fileName, "."):]
	return strings.ReplaceAll(uuid.Must(uuid.NewV4()).String(), "-", "") + suffix
}

func (s *LocalUpload) createUploadDir(uploadDir string) error {
	if files.DirExists(uploadDir) {
		return nil
	}
	return os.MkdirAll(uploadDir, os.ModePerm)
}

type UpyunUpload struct {
}

func (s *UpyunUpload) Upload(file *multipart.FileHeader) (*model.Upload, error) {
	return nil, nil
}

func (s *UpyunUpload) UploadBase64(base64 string) (*model.Upload, error) {
	return nil, nil
}

func getFileType(fileType string) (fType, suffix string) {
	switch fileType {
	case "image/jpeg":
		return "image", ".jpg"
	case "image/x-icon":
		return "image", ".icon"
	case "image/gif":
		return "image", ".gif"
	case "image/png":
	case "application/x-png":
		return "image", ".png"
	}
	return "unknow", ".unknow";
}
