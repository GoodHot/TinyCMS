package strs

import "fmt"

func Fmt(str string, args ...interface{}) string {
	return fmt.Sprintf(str, args...)
}
