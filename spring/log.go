package spring

type SpringLog struct {
	Level string `val:${log.level}`
}

func (s *SpringLog) Info(fmt string, args ...interface{}) {
}

func (s *SpringLog) Warn(fmt string, args ...interface{}) {
}

func (s *SpringLog) Error(fmt string, args ...interface{}) {
}

func (s *SpringLog) Debug(fmt string, args ...interface{}) {
}
