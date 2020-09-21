package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRep(t *testing.T) {
	hasReplaceRegex, _ := regexp.Compile(`\$\{(.+?)\}`)
	vals := hasReplaceRegex.FindAllStringSubmatch("ssss${ddd}${dfdfq}", -1)
	fmt.Println(vals)
}
