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
	Sort  string
	Bing  ConfigBing
	Baidu ConfigBaidu
	Zol ConfigZol
	Log   bool
}
type ConfigBing struct {
	Mode string
}
type ConfigBaidu struct {
	Word     []string
	Download bool
}

type ConfigZol struct {
	Sort     []string
	Download bool
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

	switch setting.Sort {
	case "bing":
		url, filename := wall.GetBingImageURL()
		wall.DownloadImage(url, path+"/.wall/", filename)

		var file string
		if setting.Bing.Mode == "today" {
			file = path + "/.wall/" + filename
		} else {
			file = wall.GetRandomFile(path + "/.wall/")
		}
		rw.SetFromFile(file)
	case "baidu":
		words := wall.GetRandomWord(setting.Baidu.Word)
		url, filename := wall.GetBaiduImageURL(words)
		if setting.Baidu.Download == false {
			rw.SetFromURL(url)
		} else {
			wall.DownloadImage(url, path+"/.wall/", filename)
			rw.SetFromFile(path + "/.wall/" + filename)
		}
	case "zol":
		sort := wall.GetRandomWord(setting.Zol.Sort)
		url, filename := wall.GetZolImageURL(sort)
		fmt.Println(url)
		fmt.Println(filename)

	default:
		log.Fatal("配置错误")
	}
}
