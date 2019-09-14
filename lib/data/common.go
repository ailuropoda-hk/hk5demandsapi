package data

var (
  DefaultLang       string
  Locales           []string
  VisualDataTypes   []string
  VisualFileTypes   []string
  DataSources       []VisualDataSourceStruct
  VisualDataSources []VisualDataSourceStruct
  VisualData        map[string][]VisualDataStruct
  VisualDataPath    string
)

func init() {
  DefaultLang = "zh-Hant" 
  Locales = make([]string, 3)
  Locales[0] = "zh-Hant"
  Locales[1] = "zh-Hans"
  Locales[2] = "en"

  VisualDataTypes = make([]string, 3)
  VisualDataTypes[0] = "youtube"
  VisualDataTypes[1] = "facebook-video"
  VisualDataTypes[2] = "facebook-photo"
  VisualDataPath = "./data/visualdata"
  
  VisualFileTypes = make([]string, 4)
  VisualFileTypes[0] = "mp4"
  VisualFileTypes[1] = "mkv"
  VisualFileTypes[2] = "webm"
  VisualFileTypes[3] = "jpg"
  
  VisualData = make(map[string][]VisualDataStruct)

}