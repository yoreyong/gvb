package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvb_server/config"
)

// 全局变量
var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
