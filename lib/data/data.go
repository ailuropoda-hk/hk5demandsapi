package data

import (
  "encoding/json"
  "errors"
  "io/ioutil"
  // "log"
  "net/http"
  "sort"

  // "github.com/davecgh/go-spew/spew"
  "gopkg.in/yaml.v2"
)

type VisualDataSourceStruct struct {
  Md5             string    `json:"md5"         yaml:"md5"`
  Type            string    `json:"type"        yaml:"type"`
  Token           string    `json:"token"       yaml:"token"`
  Event           string    `json:"event"       yaml:"event"`
  Category        string    `json:"cat"         yaml:"cat"`
  Title [] struct {
    Content       string    `json:"content"     yaml:"content"`
    Locale        string    `json:"locale"      yaml:"locale"`
  }                         `json:"title"       yaml:"title"`
  Desc [] struct {
    Content       string    `json:"content"     yaml:"content"`
    Locale        string    `json:"locale"      yaml:"locale"`
  }                         `json:"desc"        yaml:"desc"`
  Dimension struct{
    Height        int       `json:"height"      yaml:"height"`
    Width         int       `json:"width"       yaml:"width"`
  }                         `json:"dim"         yaml:"dim"`
  Location struct {
    Latitude      float32   `json:"lat"         yaml:"lat"`
    Longitude     float32   `json:"lng"         yaml:"lng"`
  }                         `json:"location"    yaml:"location"`
  Tags            []string  `json:"tags"        yaml:"tags"`
  Timestamp       int       `json:"ts"          yaml:"ts"`
}

type VisualDataStruct struct {
  Md5             string    `json:"md5"         yaml:"md5"`
  Type            string    `json:"type"        yaml:"type"`
  Event           string    `json:"event"       yaml:"event"`
  Category        string    `json:"cat"         yaml:"cat"`
  Source          string    `json:"source"      yaml:"source"`
  Url             string    `json:"url"         yaml:"url"`
  Embedded        string    `json:"embedded"    yaml:"embedded"`
  Thumbnail       string    `json:"thumbnail"   yaml:"thumbnail"`
  Title           string    `json:"title"       yaml:"title"`
  Desc            string    `json:"desc"        yaml:"desc"`
  Dimension struct {
    Height        int       `json:"height"      yaml:"height"`
    Width         int       `json:"width"       yaml:"width"`
  }                         `json:"dim"         yaml:"dim"`
  Location struct {
    Latitude      float32   `json:"lat"         yaml:"lat"`
    Longitude     float32   `json:"lng"         yaml:"lng"`
  }                         `json:"location"    yaml:"location"`
  Tags            []string  `json:"tags"        yaml:"tags"`
  Timestamp       int       `json:"ts"          yaml:"ts"`
}

type YouTubeJsonStruct struct {
  Title           string    `json:"title"`
  AuthorName      string    `json:"author_name"`
  ThumbnailUrl    string    `json:"thumbnail_url"`
  Height          int       `json:"height"`
  Width           int       `json:"width"`
}

var (
  DefaultLang       string
  Locales           []string
  DataSources       []VisualDataSourceStruct
  VisualData        map[string][]VisualDataStruct
)

func init() {
  DefaultLang = "zh-Hant" 
  Locales = make([]string, 3)
  Locales[0] = "zh-Hant"
  Locales[1] = "zh-Hans"
  Locales[2] = "en"

  VisualData = make(map[string][]VisualDataStruct)
  loadVisualData("./data/visualdata.yaml")
  // spew.Dump(VisualData)
}

func loadYaml(filepath string, v interface{}) error {
  yamlFile, err := ioutil.ReadFile(filepath)
  if err != nil {
    return err
  }
  err = yaml.Unmarshal(yamlFile, v)
  if err != nil {
    return err
  }
  return nil
}


func processVisualLink(mediaType string, mediaId string) (string, string, error) {
  var embedded, url string
  // var err error = nil
  switch mediaType {
  case "youtube":
    embedded = "https://www.youtube.com/embed/" + mediaId + "?rel=0"
    url = "https://www.youtube.com/watch?v=" + mediaId
  case "facebook-video":
    embedded = "https://www.facebook.com/v2.3/plugins/video.php?href=https://www.facebook.com/redbull/videos/" + mediaId 
    url = "https://www.facebook.com/watch/?v=" + mediaId
  case "facebook-photo":
    embedded = "https://graph.facebook.com/" + mediaId + "/picture"
    url = "https://www.facebook.com/watch/?v=" + mediaId
  default:
    return "", "", errors.New("Invalid Visual Type")
  }
  return embedded, url, nil
}

func processYouTubeMeta(mediaId string) (string, string, int, int, error) {
  var resp *http.Response
  var byteBody []byte
  var jsonData YouTubeJsonStruct
  var err error
  resp, err = http.Get("https://www.youtube.com/oembed?url=https://www.youtube.com/watch?v=" + mediaId + "&format=json")
  if err != nil {
    return "", "", 0, 0, err
  }
  byteBody, err = ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", "", 0, 0, err
  }
  err = json.Unmarshal([]byte(byteBody), &jsonData)
  if err != nil {
    return "", "", 0, 0, err
  }
  return jsonData.AuthorName, jsonData.ThumbnailUrl, jsonData.Height, jsonData.Width, nil
}

func processFacebookVideoMeta(mediaId string) (string, string, int, int, error) {
  return "", "https://graph.facebook.com/" + mediaId + "/picture", 128 , 226, nil
}

func processFacebookPhotoMeta(mediaId string) (string, string, int, int, error) {
  return "", "https://graph.facebook.com/" + mediaId + "/picture", 128 , 226, nil
}

func processVisualMeta(mediaType string, mediaId string) (string, string, int, int, error) {
  var source, thumbnail string
  var height, width int
  var err error = nil
  switch mediaType {
  case "youtube":
    source, thumbnail, height, width, err = processYouTubeMeta(mediaId)
    source += " on YouTube"
  case "facebook-video":
    source, thumbnail, height, width, err = processFacebookVideoMeta(mediaId)
    source += " on Facebook"
  case "facebook-photo":
    source, thumbnail, height, width, err = processFacebookPhotoMeta(mediaId)
    source += " on Facebook"
  default:
    return "", "", 0, 0, errors.New("Invalid Visual Type")
  }
  if err != nil {
    return "", "", 0, 0, err
  }
  return source, thumbnail, height, width, nil
}

func processSingleData(origData *VisualDataSourceStruct) {
  var embedded, url, source, thumbnail string
  var height, width, idx int
  var err error
  embedded, url, err = processVisualLink(origData.Type, origData.Token)
  if err != nil {
    
  }
  source, thumbnail, height, width, err = processVisualMeta(origData.Type, origData.Token)
  if err != nil {
    
  }

  for i:=0; i<len(Locales); i++{
    visualdata := VisualDataStruct {}
    visualdata.Md5 = origData.Md5
    visualdata.Type = origData.Type
    visualdata.Event = origData.Event
    visualdata.Category = origData.Category
    visualdata.Source = source
    visualdata.Url = url
    visualdata.Embedded = embedded
    visualdata.Thumbnail = thumbnail
    visualdata.Title = origData.Title[0].Content
    visualdata.Desc = origData.Desc[0].Content
    visualdata.Dimension.Height = height
    visualdata.Dimension.Width = width
    visualdata.Timestamp = origData.Timestamp
    visualdata.Tags = origData.Tags
    visualdata.Location.Latitude = origData.Location.Latitude
    visualdata.Location.Longitude = origData.Location.Longitude
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

  // spew.Dump(visualdata)
}

func loadVisualData(filepath string) error {
  var err error
  err = loadYaml(filepath, &DataSources)
  if err != nil {
    return err
  }
  for i:=0; i<len(DataSources); i++ {
    processSingleData(&DataSources[i])
  }
  
  for i:=0; i<len(Locales); i++{
    var locale = Locales[i]
    sort.Slice(VisualData[locale][:], func(i, j int) bool {
      return VisualData[locale][i].Timestamp < VisualData[locale][j].Timestamp
    })
  }
  return nil
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
