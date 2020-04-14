package times

import "time"

func TimeStr(fmt string) string {
	return TimeFmt(time.Now(), fmt)
}

func TimeFmt(t time.Time, fmt string) string {
	return t.Format(fmt)
}

func Parse(t string, fmt string) time.Time {
	result, err := time.Parse(fmt, t)
	if err != nil {
		return time.Now()
	}
	return result
}
