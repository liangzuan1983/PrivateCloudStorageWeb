package main

import (
	_ "github.com/oikomi/PrivateCloudStorageWeb/web/routers"
	"github.com/oikomi/PrivateCloudStorageWeb/web/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("api/v1/webconfig", &controllers.WebConfigController{})
	beego.SetStaticPath("/views", "/mh/mygo/src/github.com/oikomi/PrivateCloudStorageWeb/web/views")
	beego.Run()
}

