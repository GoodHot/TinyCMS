package times

import "time"

func TimeStr(fmt string) string {
	return time.Now().Format(fmt)
}
