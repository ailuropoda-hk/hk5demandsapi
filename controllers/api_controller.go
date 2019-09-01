package controllers

import (
  // "log"
  "github.com/astaxie/beego"
  "hk5demandsapi/lib/data"
)

type ApiController struct {
  beego.Controller
}

type ResponseStruct struct {
  Data            interface{}  `json:"data"`
}

func (c *ApiController) GetVisualData() {
  var locale = c.GetString("locale")
  var event = c.GetString("event")
  var cat = c.GetString("cat")
  c.Data["json"] = ResponseStruct{
    Data: data.GetVisualData(locale, event, cat),
  }
  c.ServeJSON()
}
