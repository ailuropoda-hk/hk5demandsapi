package data

import (
  "io/ioutil"
  "net/http"
  "sort"

  // "github.com/davecgh/go-spew/spew"
  // "github.com/golang/glog"
  "gopkg.in/yaml.v2"
)

func VisualDataFindByFilename(filename string) *VisualDataSourceStruct {
  for _, data := range VisualDataSources {
    if data.Filename == filename {
      return &data
    }
  }
  return nil
}

func LoadVisualData(filepath string) error {
  var yamlFile []byte
  var err error
  yamlFile, err = ioutil.ReadFile(filepath)
  if err != nil {
    return err
  }
  err = yaml.Unmarshal(yamlFile, &VisualDataSources)
  if err != nil {
    return err
  }
  for i:=0; i<len(VisualDataSources); i++ {
    processSingleData(&VisualDataSources[i])
  }
  
  for i:=0; i<len(Locales); i++{
    var locale = Locales[i]
    sort.Slice(VisualData[locale][:], func(i, j int) bool {
      return VisualData[locale][i].Date < VisualData[locale][j].Date
    })
  }
  return nil
}

func processSingleData(origData *VisualDataSourceStruct) {
  var resp *http.Response
  var idx int
  var err error

  resp, err = http.Get(origData.Thumbnail)
  if err != nil || resp.StatusCode < 200 || resp.StatusCode > 304 {
    origData.Thumbnail = ""
  }
  // glog.Info(origData.Thumbnail, err, resp.StatusCode)
  // spew.Dump(resp)
  for i:=0; i<len(Locales); i++{
    visualdata := VisualDataStruct {}
    visualdata.Filename           = origData.Filename
    visualdata.CsvNo              = origData.CsvNo
    visualdata.Type               = origData.Type
    visualdata.Token              = origData.Token
    visualdata.Event              = origData.Event
    visualdata.Category           = origData.Category
    visualdata.Date               = origData.Date
    if origData.Source != "" {
      visualdata.Source = origData.Author
    }
    if (visualdata.Source == "") {
      visualdata.Source = "Public"
    }
    switch(origData.Type) {
    case "youtube":
      visualdata.Source = visualdata.Source + " on YouTube"
    case "facebook-video":
      visualdata.Source = visualdata.Source + " on Facebook"
    case "facebook-photo":
      visualdata.Source = visualdata.Source + " on Facebook"
    }
    visualdata.Url                = origData.Url
    visualdata.Embedded           = origData.Embedded
    visualdata.Thumbnail          = origData.Thumbnail
    if len(origData.Title) > 0 {
      visualdata.Title            = origData.Title[0].Content
    }
    if visualdata.Title == "" {
      visualdata.Title            = origData.ProviderTitle
    }
    if len(origData.Desc) > 0 {
      visualdata.Desc             = origData.Desc[0].Content
    }
    visualdata.Dimension.Height   = origData.Dimension.Height
    visualdata.Dimension.Width    = origData.Dimension.Width
    visualdata.Location.Address   = origData.Location.Address
    visualdata.Location.Latitude  = origData.Location.Latitude
    visualdata.Location.Longitude = origData.Location.Longitude
    visualdata.Duration           = origData.Duration
    visualdata.Tags               = origData.Tags
    visualdata.Timestamp          = origData.Timestamp
    
    
    VisualData[Locales[i]] = append(VisualData[Locales[i]], visualdata)
  }

  for i:=0; i<len(origData.Title); i++ {
    idx = len(VisualData[origData.Title[i].Locale]) - 1
    VisualData[origData.Title[i].Locale][idx].Title = origData.Title[i].Content
  }
  for i:=0; i<len(origData.Desc); i++ {
    idx = len(VisualData[origData.Desc[i].Locale]) - 1
    VisualData[origData.Desc[i].Locale][idx].Desc = origData.Desc[i].Content
  }

  // spew.Dump(VisualData)
}

func GetVisualData(locale string, event string, cat string) []VisualDataStruct {
  var result []VisualDataStruct
  locale = checkLocale(locale)
  if (event != "" && cat != "") {
    for i:=0; i<len(VisualData[locale]); i++{
      if (VisualData[locale][i].Event == event && VisualData[locale][i].Category == cat) {
        result = append(result, VisualData[locale][i])
      }
    }
  } else if (event != ""){
    for i:=0; i<len(VisualData[locale]); i++{
      if (VisualData[locale][i].Event == event) {
        result = append(result, VisualData[locale][i])
      }
    }
  } else if (cat != "") {
    for i:=0; i<len(VisualData[locale]); i++{
      if (VisualData[locale][i].Category == cat) {
        result = append(result, VisualData[locale][i])
      }
    }
  } else {
    result = VisualData[locale]
  }
  return result
}



func checkLocale(locale string) string {
  for i:=0; i<len(Locales); i++ {
    if Locales[i] == locale {
      return locale
    }
  }
  return DefaultLang
}

