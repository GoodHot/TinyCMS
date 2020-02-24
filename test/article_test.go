package test

import (
	"bytes"
	"fmt"
	"github.com/go-ego/gse"
	"github.com/mozillazg/go-pinyin"
	"regexp"
	"strings"
	"testing"
)

func TestCover(t *testing.T)  {
	html := `<h1>你好，世界</h1>
<p><img src="http://localhost:9000/static/upload/20200224/e5758a839583421e90414e34d30e0912.png" alt="e5758a839583421e90414e34d30e0912.png" /><img src="http://localhost:9000/static/upload/20200224/7f709ae150d045cb8a8963616635f818.jpeg" alt="7f709ae150d045cb8a8963616635f818.jpeg" /><img src="http://localhost:9000/static/upload/20200224/d24f571514484b6c9c1b55a2eededa0e.jpeg" alt="d24f571514484b6c9c1b55a2eededa0e.jpeg" /></p>
`

	compile := regexp.MustCompile("<img\\s+src=\"(.+?)\"\\s+.+?/>")
	submatch := compile.FindAllStringSubmatch(html, -1)
	fmt.Println(submatch[0][1])
}

func TestDescription(t *testing.T)  {
	html := `<h1>你好，世界</h1>
<p><img src="http://localhost:9000/static/upload/20200224/ebebb3aab491435b844650d891a9cab4.jpeg" alt="ebebb3aab491435b844650d891a9cab4.jpeg" /></p>
<p>adminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadminadmin</p>

`
	tagReg := regexp.MustCompile("(<.+?>.+?</.+?>|<.+?/>)")
	compile := regexp.MustCompile("<p>(.+?)</p>")
	submatch := compile.FindAllStringSubmatch(html, -1)
	for _, m := range submatch {
		s := tagReg.ReplaceAllString(m[1], "")
		s = strings.TrimSpace(s)
		fmt.Println(s)
	}
}

func getPinyin(han string, a pinyin.Args) string {
	rst := pinyin.Paragraph(han)
	rst = strings.ReplaceAll(rst, " ", "")
	return rst
}

func TestTitle(t *testing.T)  {
	symbolReg := regexp.MustCompile("\\p{P}")
	title := "我佛了：香蕉是”莓“，草莓不是”莓“"
	dict := gse.New("zh,resource/dict/dictionary.txt", "alpha")
	hmm := dict.Cut(title,true)
	fmt.Println(hmm)
	a := pinyin.NewArgs()
	result := bytes.NewBufferString("")
	for i, s := range hmm {
		trim := strings.TrimSpace(s)
		if trim == "" {
			continue
		}
		findSymbol := symbolReg.FindAllString(trim, -1)
		if len(findSymbol) > 0 {
			trim = symbolReg.ReplaceAllString(trim, "-")
		}
		py := getPinyin(trim, a)
		fmt.Println(trim, "-->", py)
		result.WriteString(py)
		if i != len(hmm) - 1 {
			result.WriteString("-")
		}
	}
	compile := regexp.MustCompile("\\-{1,}")
	s := compile.ReplaceAllString(result.String(), "-")
	if s[len(s)-1:] == "-" {
		s = s[:len(s)-1]
	}
	fmt.Println(s)
	// 1. 分词
	// 2. 转拼音
	// 3. 拼接
	//fmt.Println(title)
	//var seg gse.Segmenter
	//seg.LoadDict("zh,resource/dict/dictionary.txt")
	//fmt.Println(seg.String([]byte(title), true))
	//segments := seg.Segment([]byte(title))
	//fmt.Println(gse.ToString(segments))
	//tagReg := regexp.MustCompile("(<.+?>.+?</.+?>|<.+?/>)")
	//compile := regexp.MustCompile("<p>(.+?)</p>")
	//submatch := compile.FindAllStringSubmatch(html, -1)
	//for _, m := range submatch {
	//	s := tagReg.ReplaceAllString(m[1], "")
	//	s = strings.TrimSpace(s)
	//	fmt.Println(s)
	//}
}