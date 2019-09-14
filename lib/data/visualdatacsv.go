package data

import (
  "encoding/csv"
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "os/exec"
  "strconv"
  "strings"
  "time"
  "github.com/golang/glog"
  "gopkg.in/yaml.v2"
  // "github.com/davecgh/go-spew/spew"
)

func ProcessCsvFile(csvfilename string, vdfilename string) {
  var youTubeMeta *YouTubeMetaStruct
  var facebookVideoMeta *FacebookVideoMetaStruct
  var facebookPhotoMeta *FacebookPhotoMetaStruct
  var data *VisualDataSourceStruct
  var csvList []VisualDataCsvStruct
  var filename string
  var outByte []byte
  var isNew bool
  var err error
  
  csvList, err = readCsvFile(csvfilename)
  if err != nil {
    fmt.Println(err)
  }

  for i:=0; i<len(csvList); i++ {
  // for i:=0; i<60; i++ {
    csvitem := csvList[i]
    filename, err = downloadVisualData(csvitem)
    if err != nil {
      glog.Error(err)
    }
    glog.Info("Item: ", i, " Processed: ", filename)
    switch(csvitem.Type) {
    case "youtube":
      youTubeMeta, err = getYouTubeMeta(csvitem.Token, filename)
    case "facebook-video":
      facebookVideoMeta, err = getFacebookVideoMeta(csvitem.Token, filename)
    case "facebook-photo":
      facebookPhotoMeta, err = getFacebookPhotoMeta(csvitem.Token, filename)
    }
    
    if err != nil {
      glog.Error(err)
      continue
    }

    data = VisualDataFindByFilename(filename)
    if data != nil {
      isNew = false
    } else {
      isNew = true
      data = &VisualDataSourceStruct {
        Filename:   filename,
      }
    }
  
    if csvitem.Number != 0 {
      data.CsvNo              = csvitem.Number
    }
    if csvitem.Type != "" {
      data.Type               = csvitem.Type
    }
    if csvitem.Token != "" {
      data.Token              = csvitem.Token
    }
    if csvitem.Event != "" {
      data.Event              = csvitem.Event
    }
    if csvitem.Category != "" {
      data.Category           = csvitem.Category
    }
    if csvitem.Date != "" {
      data.Date               = csvitem.Date
    }
    if csvitem.Source != "" {
      data.Source             = csvitem.Source
    }
    if csvitem.Timestamp != 0 {
      data.Timestamp          = csvitem.Timestamp
    }
    if csvitem.Location != "" {
      data.Location.Address   = csvitem.Location
    }
    if csvitem.Latitude != 0 {
      data.Location.Latitude  = csvitem.Latitude
    }
    if csvitem.Longitude != 0 {
      data.Location.Longitude  = csvitem.Longitude
    }
    if csvitem.TitleZhHant != "" {
      data.Title = append(data.Title, LocaleContentStruct{
        Content: csvitem.TitleZhHant,
        Locale:  "zh-Hant",
      })
    }
    if csvitem.TitleEn != "" {
      data.Title = append(data.Title, LocaleContentStruct{
        Content: csvitem.TitleEn,
        Locale:  "en",
      })
    }
    if csvitem.DescZhHant != "" {
      data.Desc = append(data.Desc, LocaleContentStruct{
        Content: csvitem.DescZhHant,
        Locale:  "zh-Hant",
      })
    }
    if csvitem.DescEn != "" {
      data.Desc = append(data.Desc, LocaleContentStruct{
        Content: csvitem.DescEn,
        Locale:  "en",
      })
    }

    var dataTags []string
    if csvitem.Tag01 != "" { dataTags = append(dataTags, csvitem.Tag01) }
    if csvitem.Tag02 != "" { dataTags = append(dataTags, csvitem.Tag02) }
    if csvitem.Tag03 != "" { dataTags = append(dataTags, csvitem.Tag03) }
    if csvitem.Tag04 != "" { dataTags = append(dataTags, csvitem.Tag04) }
    if csvitem.Tag05 != "" { dataTags = append(dataTags, csvitem.Tag05) }
    if csvitem.Tag06 != "" { dataTags = append(dataTags, csvitem.Tag06) }
    if csvitem.Tag07 != "" { dataTags = append(dataTags, csvitem.Tag07) }
    if csvitem.Tag08 != "" { dataTags = append(dataTags, csvitem.Tag08) }
    if csvitem.Tag09 != "" { dataTags = append(dataTags, csvitem.Tag09) }
    if csvitem.Tag10 != "" { dataTags = append(dataTags, csvitem.Tag10) }
    if csvitem.Tag11 != "" { dataTags = append(dataTags, csvitem.Tag11) }
    if csvitem.Tag12 != "" { dataTags = append(dataTags, csvitem.Tag12) }
    if csvitem.Tag13 != "" { dataTags = append(dataTags, csvitem.Tag13) }
    if csvitem.Tag14 != "" { dataTags = append(dataTags, csvitem.Tag14) }
    if csvitem.Tag15 != "" { dataTags = append(dataTags, csvitem.Tag15) }

    if (len(dataTags) > len(data.Tags)) {
      data.Tags = dataTags
    }
    switch(csvitem.Type) {
    case "youtube":
      data.Duration           = youTubeMeta.Duration
      data.Dimension.Height   = youTubeMeta.Height
      data.Dimension.Width    = youTubeMeta.Width
      data.Author             = youTubeMeta.AuthorName
      data.ProviderTitle      = youTubeMeta.Title
      data.Url                = youTubeMeta.Url
      data.Embedded           = youTubeMeta.Embedded
      data.Thumbnail          = youTubeMeta.Thumbnail
    case "facebook-video":
      data.Duration           = facebookVideoMeta.Duration
      data.Dimension.Height   = facebookVideoMeta.Height
      data.Dimension.Width    = facebookVideoMeta.Width
      data.Url                = facebookVideoMeta.Url
      data.Embedded           = facebookVideoMeta.Embedded
      data.Thumbnail          = facebookVideoMeta.Thumbnail
    case "facebook-photo":
      data.Dimension.Height   = facebookPhotoMeta.Height
      data.Dimension.Width    = facebookPhotoMeta.Width
      data.Url                = facebookPhotoMeta.Url
      data.Embedded           = facebookPhotoMeta.Embedded
      data.Thumbnail          = facebookPhotoMeta.Thumbnail
    }
    // spew.Dump(data)
    if isNew {
      VisualDataSources = append(VisualDataSources, *data)
    }
    
  }
  outByte, err = yaml.Marshal(VisualDataSources)
  if err != nil {
    glog.Error(err)
  }
  err = ioutil.WriteFile(vdfilename, outByte, 0644)
    

}

func isDownloaded(datatype string, token string) (bool, string) {
  var filename, filepath, newfilename string
  var isExist bool 
  isExist = false
  switch (datatype) {
  case "youtube":
    filename = "ytub-" + token
  case "facebook-video":
    filename = "fbvi-" + token
  case "facebook-photo":
    filename = "fbph-" + token
  }

  for _, filetype := range VisualFileTypes {
    newfilename = filename + "." + filetype
    filepath = VisualDataPath + "/" + newfilename
    if _, err := os.Stat(filepath); err == nil {
      isExist = true
      break
    } 
    // else {
    //   glog.Error(err)
    // }
  }
  if isExist {
    return true, newfilename
  } else {
    return false, ""
  }
}

func downloadVisualData(csvitem VisualDataCsvStruct) ( string, error) {
  var cmdOutput []byte
  var filepath, newfilename, newfilepath, fileSuffix, url string
  var cmd *exec.Cmd
  var isExist bool
  var err error

  isExist, newfilename = isDownloaded(csvitem.Type, csvitem.Token)
  // glog.Info("Is Downloaded", csvitem.Type, csvitem.Token, isExist, newfilename)
  if isExist {
    return newfilename, nil
  }

  switch(csvitem.Type) {
  case "youtube":
    url = "https://www.youtube.com/watch?v=" + csvitem.Token
    cmd = exec.Command("youtube-dl", "--output", VisualDataPath + "/temp", url)
    newfilename = "ytub-" + csvitem.Token
  case "facebook-video":
    url = "https://www.facebook.com/video.php?v=" + csvitem.Token
    cmd = exec.Command("youtube-dl", "--output", VisualDataPath + "/temp", url)
    newfilename = "fbvi-" + csvitem.Token
  case "facebook-photo":
    url = "https://graph.facebook.com/" + csvitem.Token + "/picture"
    newfilename = "fbph-" + csvitem.Token
    cmd = exec.Command("wget", "-O", VisualDataPath + "/temp.jpg", url)
  default:
    return "", errors.New("Invalid Data Type")
  }
  // glog.Info("Download ", csvitem.Type, " ", csvitem.Token)
  cmdOutput, err = cmd.CombinedOutput()
  glog.Info(string(cmdOutput))
  if err != nil {
    return "", err
  }

  for _, filetype := range VisualFileTypes {
    fileSuffix = filetype
    filepath = VisualDataPath + "/temp." + fileSuffix
    if _, err := os.Stat(filepath); err == nil {
      break
    }
  }

  newfilename = newfilename + "." + fileSuffix
  newfilepath = VisualDataPath + "/" + newfilename
  err = os.Rename(filepath, newfilepath)
  if err != nil {
    return "", err
  }
  
  return newfilename, nil
}

func getFacebookPhotoMeta(token string, filename string) (*FacebookPhotoMetaStruct, error) {
  var result FacebookPhotoMetaStruct
  var filepath string 
  var height, width int
  var err error

  filepath =  VisualDataPath + "/" + filename
  height, width, err = getPhotoDim(filepath)
  if err != nil {
    return nil, err
  }

  result.Url      = "https://www.facebook.com/" + token
  result.Embedded = "https://graph.facebook.com/" + token + "/picture"
  result.Thumbnail= "https://graph.facebook.com/" + token + "/picture"
  result.Height   = height
  result.Width    = width

  return &result, nil
} 

func getFacebookVideoMeta(token string, filename string) (*FacebookVideoMetaStruct, error) {
  var result FacebookVideoMetaStruct
  var ffmpegmeta *FFMpegMetaStruct
  var duration float64
  var filepath string 
  var err error

  filepath =  VisualDataPath + "/" + filename
  ffmpegmeta, err = getFfmpegMeta(filepath)
  if err != nil {
    return nil, err
  }
  result.Height   = ffmpegmeta.Streams[0].Height
  result.Width    = ffmpegmeta.Streams[0].Width
  duration, _ = strconv.ParseFloat(ffmpegmeta.Fromat.Duration, 32)
  result.Duration = (float32)(duration)

  result.Url      = "https://www.facebook.com/watch/?v=" + token
  result.Embedded = "https://www.facebook.com/v2.3/plugins/video.php?href=https://www.facebook.com/redbull/videos/" + token 
  result.Thumbnail= "https://graph.facebook.com/" + token + "/picture"

  return &result, nil
}

func getYouTubeMeta(token string, filename string) (*YouTubeMetaStruct, error) {
  var result YouTubeMetaStruct
  var jsonData YouTubeJsonStruct
  var resp *http.Response
  var ffmpegmeta *FFMpegMetaStruct
  var byteBody []byte
  var duration float64
  var filepath string 
  var err error
  resp, err = http.Get("https://www.youtube.com/oembed?url=https://www.youtube.com/watch?v=" + token + "&format=json")
  if err!= nil {
    glog.Error(err)
    return nil, err
  }
  byteBody, err = ioutil.ReadAll(resp.Body)
  if err != nil {
    glog.Error(err)
    return nil, err
  }
  err = json.Unmarshal([]byte(byteBody), &jsonData)
  if err != nil {
    glog.Error(err)
    return nil, err
  }
  result.Title      = jsonData.Title
  result.AuthorName = jsonData.AuthorName
  result.Thumbnail  = jsonData.ThumbnailUrl
  result.Height     = jsonData.Height
  result.Width      = jsonData.Width

  filepath =  VisualDataPath + "/" + filename
  ffmpegmeta, err = getFfmpegMeta(filepath)
  if err != nil {
    glog.Error(err)
    return nil, err
  }
  duration, _ = strconv.ParseFloat(ffmpegmeta.Fromat.Duration, 32)
  result.Duration = (float32)(duration)

  result.Url      = "https://www.youtube.com/watch?v=" + token
  result.Embedded = "https://www.youtube.com/embed/" + token + "?rel=0"
  
  return &result, nil
}

func getFfmpegMeta(filepath string) (*FFMpegMetaStruct, error) {
  var ffmpegmeta FFMpegMetaStruct
  var cmd *exec.Cmd
  var cmdOutput []byte
  var err error
  cmd = exec.Command("ffprobe", "-v", "quiet", "-print_format", "json",
    "-show_format", "-show_streams", filepath)
  cmdOutput, err = cmd.CombinedOutput()
  if err != nil {
    return nil, err
  }
  // glog.Info(string(cmdOutput))
  err = json.Unmarshal(cmdOutput, &ffmpegmeta)
  if err != nil {
    return nil, err
  }
  return &ffmpegmeta, nil
}

func getPhotoDim(filepath string) (int, int, error) {
  var cmd *exec.Cmd
  var cmdOutput []byte
  var height, width int
  var strOutput []string
  var err error
  cmd = exec.Command("identify", "-format", "%hx%w", filepath)
  cmdOutput, err = cmd.CombinedOutput()
  if err != nil {
    return 0, 0, err
  }
  strOutput = strings.Split((string)(cmdOutput),"x")
  if len(strOutput) != 2 {
    return 0, 0, errors.New("Invalid identify result")
  }
  height, err = strconv.Atoi(strOutput[0])
  if err != nil {
    return 0, 0, err
  }
  width, err = strconv.Atoi(strOutput[1])
  if err != nil {
    return 0, 0, err
  }
  return height, width, nil
}

func readCsvFile(filepath string) ([]VisualDataCsvStruct, error) {
  var result []VisualDataCsvStruct
  
  csvfile, err := os.Open(filepath)
  if err != nil {
    return result, err
  }

  lines, err := csv.NewReader(csvfile).ReadAll()
  if err != nil {
    return result, err
  }

  for _, line := range lines {

    datatype := strings.ToLower(strings.TrimSpace(line[1]))
    isValidType := false
    for _, v := range VisualDataTypes {
      if datatype == v {
        isValidType = true
      }
    } 
    
    if (isValidType) {
      var number, ts int
      var lat, lng float64
      var t time.Time
      var date string
      number, _ = strconv.Atoi(line[0])
      ts, _ = strconv.Atoi(line[14])
      lat, _ = strconv.ParseFloat(line[12], 32)
      lng, _ = strconv.ParseFloat(line[13], 32)
      err = nil
      t, err = time.Parse("01-02-2006", line[5])
      if err != nil {
        date = "unknown"
      } else {
        date = t.Format("01-02-2006")
      }
      data := VisualDataCsvStruct {
        Number:         number,
        Type:           datatype,
        Token:          strings.TrimSpace(line[2]),
        Category:       strings.ToLower(strings.TrimSpace(line[3])),
        Event:          strings.ToLower(strings.TrimSpace(line[4])),
        Date:           date,
        Source:         strings.TrimSpace(line[6]),
        TitleZhHant:    strings.TrimSpace(line[7]),
        TitleEn:        strings.TrimSpace(line[8]),
        DescZhHant:     strings.TrimSpace(line[9]),
        DescEn:         strings.TrimSpace(line[10]),
        Location:       strings.TrimSpace(line[11]),
        Latitude:       float32(lat),
        Longitude:      float32(lng),
        Timestamp:      ts,
        Tag01:          strings.TrimSpace(line[15]),
        Tag02:          strings.TrimSpace(line[16]),
        Tag03:          strings.TrimSpace(line[17]),
        Tag04:          strings.TrimSpace(line[18]),
        Tag05:          strings.TrimSpace(line[19]),
      }
      result = append(result, data)
    }
  }
  return result, nil
}