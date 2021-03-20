package file

import (
	"project/utils"
	"project/utils/config"
)

func Init() error {
	return utils.IsNotExistMkDir(config.ApplicationConfig.StaticPath)
}
