package main

import (
  "flag"
  "log"
  "github.com/astaxie/beego"
  "github.com/golang/glog"
  "hk5demandsapi/lib/data"
  _ "hk5demandsapi/routers"
)

var isReadCSV bool
func init() {
  flag.Set("stderrthreshold", "INFO")
  readcsvPtr := flag.Bool("readcsv", false, "a bool")
  flag.Parse()

  isReadCSV = *readcsvPtr
  glog.Info("readcsv:", isReadCSV)
}

func main() {
  var err error
  
  
  if isReadCSV {
    data.ProcessCsvFile("./data/data.csv", "./data/visualdata.yaml")
    return
  }
  
  data.LoadVisualData("./data/visualdata.yaml")
  

  

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
