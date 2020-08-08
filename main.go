package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"wallpaper/wall"
	rw "github.com/reujab/wallpaper"
)

type Config struct {
	Sort     string
	Bing     ConfigBing
	Baidu    ConfigBaidu
	Log      bool
	WallPath string
}
type ConfigBing struct {
	Mode string
}
type ConfigBaidu struct {
	Word string
}

func getConfig(filename string) Config {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("配置文件获取错误")
	}
	var setting Config
	err = json.Unmarshal(content, &setting)
	if err != nil {
		fmt.Println("配置文件解析错误")
	}
	return setting
}

func main() {
	path := wall.GetCurrentPath()

	setting := getConfig(path + "/config.json")

	if setting.Log == true {
		wall.OpenLog(path + "/log.txt")
	}

	var file string
	switch setting.Sort {
	case "bing":
		url, filename := wall.GetBingImageURL()
		wall.DownloadImage(url, path+"/"+setting.WallPath+"/", filename)
		if setting.Bing.Mode == "today" {
			file = path + "/" + setting.WallPath + "/" + filename
		} else {
			file = wall.GetRandomFile(path + "/" + setting.WallPath + "/")
		}
	case "baidu":
		fallthrough
	default:
		log.Fatal("配置错误")
	}

	rw.SetFromFile(file)
}
