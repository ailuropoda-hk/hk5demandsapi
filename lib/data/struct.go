package data

type LocaleContentStruct struct {
  Content         string    `json:"content"     yaml:"content"`
  Locale          string    `json:"locale"      yaml:"locale"`
}

type VisualDataSourceStruct struct {
  Filename        string    `json:"filename"    yaml:"filename"`
  CsvNo           int       `json:"csvno"       yaml:"csvno"`
  Type            string    `json:"type"        yaml:"type"`
  Token           string    `json:"token"       yaml:"token"`
  Event           string    `json:"event"       yaml:"event"`
  Category        string    `json:"cat"         yaml:"cat"`
  Date            string    `json:"date"        yaml:"date"`
  Source          string    `json:"source"      yaml:"source"`
  Url             string    `json:"url"         yaml:"url"`
  Embedded        string    `json:"embedded"    yaml:"embedded"`
  Thumbnail       string    `json:"thumbnail"   yaml:"thumbnail"`
  Author          string    `json:"author"      yaml:"author"`
  ProviderTitle   string    `json:"providerTitle"     yaml:"providerTitle"`
  Title   []LocaleContentStruct  `json:"title"  yaml:"title"`
  Desc    []LocaleContentStruct  `json:"desc"   yaml:"desc"`
  Dimension struct{
    Height        int       `json:"height"      yaml:"height"`
    Width         int       `json:"width"       yaml:"width"`
  }                         `json:"dim"         yaml:"dim"`
  Location struct {
    Address       string    `json:"address"     yaml:"address"`
    Latitude      float32   `json:"lat"         yaml:"lat"`
    Longitude     float32   `json:"lng"         yaml:"lng"`
  }                         `json:"location"    yaml:"location"`
  Duration        float32   `json:"duration"    yaml:"duration"`
  Tags            []string  `json:"tags"        yaml:"tags"`
  Timestamp       int       `json:"ts"          yaml:"ts"`
}

type VisualDataStruct struct {
  Filename        string    `json:"filename"    yaml:"filename"`
  CsvNo           int       `json:"csvno"       yaml:"csvno"`
  Type            string    `json:"type"        yaml:"type"`
  Token           string    `json:"token"       yaml:"token"`
  Event           string    `json:"event"       yaml:"event"`
  Category        string    `json:"cat"         yaml:"cat"`
  Date            string    `json:"date"        yaml:"date"`
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
    Address       string    `json:"address"     yaml:"address"`
    Latitude      float32   `json:"lat"         yaml:"lat"`
    Longitude     float32   `json:"lng"         yaml:"lng"`
  }                         `json:"location"    yaml:"location"`
  Duration        float32   `json:"duration"    yaml:"duration"`
  Tags            []string  `json:"tags"        yaml:"tags"`
  Timestamp       int       `json:"ts"          yaml:"ts"`
}

type FacebookPhotoMetaStruct struct {
  Url             string
  Embedded        string
  Thumbnail       string
  Height          int
  Width           int
}

type FacebookVideoMetaStruct struct {
  Url             string
  Embedded        string
  Thumbnail       string
  Height          int
  Width           int
  Duration        float32
}

type YouTubeMetaStruct struct {
  Title           string
  AuthorName      string
  Url             string
  Embedded        string
  Thumbnail       string
  Height          int
  Width           int
  Duration        float32
}

type YouTubeJsonStruct struct {
  Title           string    `json:"title"`
  AuthorName      string    `json:"author_name"`
  ThumbnailUrl    string    `json:"thumbnail_url"`
  Height          int       `json:"height"`
  Width           int       `json:"width"`
}

type FFMpegMetaStruct struct {
  Streams []struct {
    Index         int       `json:"index"`
    Height        int       `json:"height"`
    Width         int       `json:"width"`
  }                         `json:"streams"`
  Fromat struct {
    Duration      string    `json:"duration"`
  }                         `json:"format"`
}

type VisualDataCsvStruct struct {
  Number          int
  Type            string
  Token           string
  Category        string
  Event           string
  Date            string
  Source          string
  TitleZhHant     string
  TitleEn         string
  DescZhHant      string
  DescEn          string
  Location        string
  Latitude        float32
  Longitude       float32
  Timestamp       int
  Tag01           string
  Tag02           string
  Tag03           string
  Tag04           string
  Tag05           string
  Tag06           string
  Tag07           string
  Tag08           string
  Tag09           string
  Tag10           string
  Tag11           string
  Tag12           string
  Tag13           string
  Tag14           string
  Tag15           string
}