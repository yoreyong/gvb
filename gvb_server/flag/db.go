package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
)

func MakeMigrations() {
	var err error
	err = global.DB.SetupJoinTable(&models.UserModel{}, "SubscribesModels", &models.UserSubscribeModel{}) // 人为设置第3张表, User <-> Subscribes
	if err != nil {
		global.Log.Error("[ error ] Fail to setup join tables: SubscribesModels")
	}
	err = global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{}) // 认为设置第3张表, Menu <-> Images
	if err != nil {
		global.Log.Error("[ error ] Fail to setup join tables: MenuBannerModels")
	}
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.AdvertiseModel{},
			&models.ArticleModel{},
			&models.BannerModel{},
			&models.CommentModel{},
			&models.FeedbackModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.MessageModel{},
			&models.ModeModel{},
			&models.TagModel{},
			&models.UserSubscribeModel{},
			&models.UserModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] Faile to generate SQL tables.")
		return
	}
	global.Log.Info("[ success ] Success to generate SQL tables.")
}

func PVersion() {
	fmt.Println("Version")
}
