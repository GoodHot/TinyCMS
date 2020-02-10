package times

import "time"

func TimeStr(fmt string) string {
	return time.Now().Format(fmt)
}

func Parse(t string, fmt string) time.Time {
	result, err := time.Parse(fmt, t)
	if err != nil {
		return time.Now()
	}
	return result
}
