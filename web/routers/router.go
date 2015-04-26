package routers

import (
	"github.com/oikomi/PrivateCloudStorageWeb/web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
