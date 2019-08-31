package main

import (
  "log"
  "github.com/astaxie/beego"

  _ "hk5demandsapi/routers"
)

func main() {
  var err error

  

  beego.BConfig.Listen.HTTPPort = 8080 //, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
  if err != nil {
    log.Printf("Error loading HTTP_PORT")
  }

  beego.BConfig.CopyRequestBody = true
  beego.BConfig.WebConfig.ViewsPath = "web/views"
  beego.SetStaticPath("/vendor", "web/vendor")
  beego.SetStaticPath("/img", "web/img")
  beego.SetStaticPath("/js", "web/js")
  beego.SetStaticPath("/css", "web/css")
  beego.BConfig.AppName = "HK5demands"

  beego.Run()
}
