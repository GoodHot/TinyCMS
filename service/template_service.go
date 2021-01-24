package service

import (
	"errors"
	"fmt"
	"github.com/GoodHot/TinyCMS/common"
	"github.com/GoodHot/TinyCMS/config"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/Masterminds/sprig"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

type TemplateService struct {
	templates     *template.Template
	DictService   *DictService `ioc:"auto"`
	PostService   *PostService `ioc:"auto"`
	CodeService   *CodeService `ioc:"auto"`
	DefLayoutFile string       `val:"default"`
	DefLayout     string       `val:"layout"`
	DefSuffix     string       `val:".html"`
	Placeholder   string       `val:"$${content_placeholder}$$"`
	TmplDir       string       `val:"${props.template.dir}"`
	AssetsDir     string       `val:"${props.template.assetsDir}"`
	StaticDir     string       `val:"${server.static}"`
	funcMap       *template.FuncMap
	tmplName      string
}

func (ts *TemplateService) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmplID := ts.tmplID(name)
	tmpl := ts.templates.Lookup(tmplID)
	if tmpl == nil {
		return errors.New("not found template " + name)
		//if tmp, err := ts.LoadTemplate(name); err != nil {
		//	return err
		//} else {
		//	tmpl = tmp
		//}
	}
	err := tmpl.ExecuteTemplate(w, tmplID, data)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (ts *TemplateService) LoadTemplate(name string) (*template.Template, error) {
	fileData, err := ts.loadTemplateFile(name + ts.DefSuffix)
	if err != nil {
		return nil, err
	}
	layout := ts.getLayout(fileData)
	if layout != ts.DefLayoutFile {
		fileData = fileData[strings.Index(fileData, "\n")+1:]
	}
	layoutData, err := ts.loadTemplateFile(ts.DefLayout + "/" + layout + ts.DefSuffix)
	if err != nil {
		return nil, err
	}
	html := strings.ReplaceAll(layoutData, ts.Placeholder, fileData)
	tmplID := ts.tmplID(name)
	return ts.templates.Parse(`{{define "` + tmplID + `"}}` + html + `{{end}}`)
}

func (ts *TemplateService) loadTemplateFile(path string) (string, error) {
	filePath := fmt.Sprintf("%v/%v/%v", ts.TmplDir, ts.tmplName, path)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (ts *TemplateService) tmplID(name string) string {
	return ts.tmplName + "__" + name
}

func (ts *TemplateService) buildFuncMap() template.FuncMap {
	if ts.funcMap != nil {
		return *ts.funcMap
	}
	funcMap := sprig.FuncMap()
	assetsURLPrefix := fmt.Sprintf("/%v/%v/%v", ts.StaticDir, ts.AssetsDir, ts.tmplName)
	funcMap["assetsURL"] = func(src string) string {
		return assetsURLPrefix + src
	}
	funcMap["page"] = func(page int, size int) *trait.Page {
		result, err := ts.PostService.Page(&trait.Query{
			Page: page,
			Size: size,
			Order: map[string]string{
				"id": "desc",
			},
		})
		if err != nil {
			return nil
		}
		return result
	}
	funcMap["code"] = func(key string) *trait.Code {
		code, _ := ts.CodeService.GetByKey(key)
		if code == nil {
			return &trait.Code{}
		}
		return code
	}
	funcMap["html"] = func(html string) template.HTML {
		return template.HTML(html)
	}
	ts.funcMap = &funcMap
	return *ts.funcMap
}

func (ts *TemplateService) Startup() error {
	dict, err := ts.DictService.Get(config.DictKeySkinTemplate)
	if err != nil {
		return err
	}
	ts.tmplName = dict.Value
	funcMap := ts.buildFuncMap()
	ts.templates = template.New(ts.tmplName)
	ts.templates.Funcs(funcMap)
	root := fmt.Sprintf("%v/%v", ts.TmplDir, ts.tmplName)
	if dirs, err := ioutil.ReadDir(root); err != nil {
		return err
	} else {
		for _, file := range dirs {
			if file.IsDir() {
				continue
			}
			name := file.Name()
			split := strings.LastIndex(name, ".")
			if split < 0 {
				continue
			}
			suffix := name[split:]
			if suffix != ts.DefSuffix {
				continue
			}
			name = name[:split]
			ts.LoadTemplate(name)
		}
	}
	// 查找并移动模板目录中的assets
	if err := ts.copyAssets(); err != nil {
		return err
	}
	return nil
}

func (ts *TemplateService) copyAssets() error {
	assetsDir := fmt.Sprintf("%v/%v/%v", ts.TmplDir, ts.tmplName, ts.AssetsDir)
	if !common.FileExists(assetsDir) || !common.IsDir(assetsDir) {
		return nil
	}
	desPath := fmt.Sprintf("%v/%v/%v", ts.StaticDir, ts.AssetsDir, ts.tmplName)
	if !common.FileExists(desPath) || !common.IsDir(desPath) {
		if err := common.MakeDir(desPath); err != nil {
			return err
		}
	}
	return common.CopyDir(assetsDir, desPath)
}

func (ts *TemplateService) getLayout(data string) string {
	line := strings.Split(data, "\n")
	firstLine := line[0]
	reg, _ := regexp.Compile("\\$\\${layout:(.+?)}\\$\\$")
	rst := reg.FindAllStringSubmatch(firstLine, -1)
	if rst != nil {
		return rst[0][1]
	}
	return ts.DefLayoutFile
}
