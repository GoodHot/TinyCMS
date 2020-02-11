package render

import (
	"errors"
	"fmt"
	"github.com/GoodHot/TinyCMS/common/strs"
	"html/template"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

type HTMLRenderer struct {
	LayoutDir    string // layout dir name
	ComponentDir string // component dir name
	TemplateDir  string // template dir
	Suffix       string // template file suffix
	tmplCacheMap map[string]string
	tmpl         *template.Template
	funcMap      template.FuncMap
	Compress     bool // compress html result
}

func (s *HTMLRenderer) Render(w io.Writer, name string, data interface{}) error {
	tmplName := s.tmplCacheMap[name]
	if tmplName == "" {
		return errors.New("can't found template file：" + name)
	}
	err := s.executeTmpl(w, tmplName, data)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (s *HTMLRenderer) executeTmpl(writer io.Writer, name string, data interface{}) error {
	if data == nil {
		data = make(map[string]interface{})
	}
	return s.tmpl.ExecuteTemplate(writer, name, data)
}

func (s *HTMLRenderer) loadComponent() error {
	if s.ComponentDir == "" {
		return nil
	}
	dirs, _ := ioutil.ReadDir(s.TemplateDir + "/" + s.ComponentDir)
	if dirs == nil || len(dirs) == 0 {
		return nil
	}
	var fileNames []string
	for _, v := range dirs {
		if v.IsDir() {
			continue
		}
		name := v.Name()[0:strings.LastIndex(v.Name(), ".")]
		fileNames = append(fileNames, name)
	}
	if fileNames == nil || len(fileNames) == 0 {
		return nil
	}
	for _, v := range fileNames {
		bt, err := ioutil.ReadFile(s.TemplateDir + "/" + s.ComponentDir + "/" + v + s.Suffix)
		if err != nil {
			return err
		}
		html := "{{define \"comp:" + v + "\"}}" + string(bt) + "{{end}}"
		if s.Compress {
			html = s.compressHTML(html)
		}
		if _, err = s.tmpl.Parse(html); err != nil {
			return err
		}
	}
	return nil
}

func (s *HTMLRenderer) compressHTML(src string) string {
	var re = regexp.MustCompile(`([\r\n]| {3,})`)
	html := re.ReplaceAllString(src, "")
	return html
}

func (s *HTMLRenderer) loadPage() error {
	if s.TemplateDir == "" {
		return nil
	}
	dirs, _ := ioutil.ReadDir(s.TemplateDir)
	if dirs == nil || len(dirs) == 0 {
		return nil
	}
	var fileNames []string
	for _, v := range dirs {
		if v.IsDir() {
			continue
		}
		name := v.Name()[0:strings.LastIndex(v.Name(), ".")]
		fileNames = append(fileNames, name)
	}
	if fileNames == nil || len(fileNames) == 0 {
		return nil
	}
	regex, _ := regexp.Compile("##layout:(.+?)##")
	for _, v := range fileNames {
		bt, err := ioutil.ReadFile(s.TemplateDir + "/" + v + s.Suffix)
		if err != nil {
			return err
		}
		html := string(bt)
		var firstLine string
		if strings.Index(html, "\n") != -1 {
			firstLine = html[0:strings.Index(html, "\n")]
		}
		layout := "default"
		layoutGroup := regex.FindStringSubmatch(firstLine)
		if len(layoutGroup) > 0 {
			layout = layoutGroup[1]
		}
		html, err = s.readLayoutFile(layout, v)
		if err != nil {
			return err
		}
		html = regex.ReplaceAllString(html, "")
		_, err = s.tmpl.Parse(html)
		if err != nil {
			return err
		}
		s.tmplCacheMap[v] = s.templateName(layout, v)
	}
	return nil
}

func (s *HTMLRenderer) readLayoutFile(layout, page string) (string, error) {
	layoutFile := strs.Fmt("%s/%s/%s%s", s.TemplateDir, s.LayoutDir, layout, s.Suffix)
	layoutHTML, _ := ioutil.ReadFile(layoutFile)
	if layoutHTML == nil {
		layoutHTML = []byte("##content##")
	}
	pageFile := strs.Fmt("%s/%s%s",s.TemplateDir, page, s.Suffix)
	pageHTML, err := ioutil.ReadFile(pageFile)
	if err != nil {
		return "", err
	}
	html := strings.Replace("{{define \""+s.templateName(layout, page)+"\"}}"+string(layoutHTML)+"{{end}}", "##content##", string(pageHTML), -1)
	if s.Compress {
		html = s.compressHTML(html)
	}
	return html, nil
}

func (*HTMLRenderer) templateName(layout, page string) string {
	return layout + "__" + page
}

func (s *HTMLRenderer) Init(funcMap template.FuncMap) {
	s.tmpl = template.New("TinyCMS").Funcs(funcMap)
	s.tmplCacheMap = make(map[string]string)
	s.loadComponent()
	s.loadPage()
}
