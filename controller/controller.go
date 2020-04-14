package controller

import (
	"github.com/GoodHot/TinyCMS/brain"
)

type Controller struct {
	Log brain.Logger `ioc:"auto"`
}

func (s *Controller) Startup() error {
	s.Log.Info("gogogogogo")
	return nil
}
