package install

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common/files"
)

type InstallService struct {
}

/**
 * check application is install
 */
func (s *InstallService) CheckInstalled(path string) bool {
	isInstall := files.FileExists(path)
	if isInstall {
		fmt.Println("application installed")
		return true
	}
	return false
}
