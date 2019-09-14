package controllers

import (
  "github.com/astaxie/beego"
)

type WebController struct {
  beego.Controller
}

func (c *WebController) Prepare() {
  c.Layout = "web/defaultlayout.html"
}

func (c *WebController) Get() {
  c.TplName = "web/dashboard.html"
}

func (c *WebController) GetMapEvent() {
  c.TplName = "web/mapevent.html"
}