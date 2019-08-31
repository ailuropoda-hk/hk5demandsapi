package routers

import (
  "hk5demandsapi/controllers"

  "github.com/astaxie/beego"
)

func init() {
  
  // API
  beego.Router("/api/visualdata", &controllers.ApiController{}, "get:GetVisualData")

}
