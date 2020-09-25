package main

import (
	"encoding/json"
	rw "github.com/reujab/wallpaper"
	"io/ioutil"
	"log"
	"wallpaper/wall"
)

type Config struct {
	Sort  []string
	Bing  ConfigBing
	Baidu ConfigBaidu
	Zol   ConfigZol
	Sogou ConfigSogou
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

type ConfigSogou struct {
	Sort     []string
	Download bool
}

func getConfig(filename string) (setting Config) {
	defer func() {
		if err := recover(); err != nil {
			if len(setting.Sort) == 0 {
				setting.Sort = append(setting.Sort, "bing")
			}

			if setting.Bing.Mode != "today" {
				setting.Bing.Mode = "random"
			}

			setting.Log = true
		}
	}()

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("配置文件获取错误")

	}
	err = json.Unmarshal(content, &setting)
	if err != nil {
		panic("配置文件解析错误")
	}
	return setting
}

func main() {
	path := wall.GetCurrentPath()

	setting := getConfig(path + "/config.json")

	if setting.Log == true {
		wall.OpenLog(path + "/log.txt")
	}

	sort := wall.GetRandomWord(setting.Sort)
	switch sort {
	case "bing":
		url, filename := wall.GetBingImageURL()
		wall.DownloadImage(url, filename)

		var file string
		if setting.Bing.Mode == "today" {
			file = wall.GetWallpaperSavePath() + filename
		} else {
			file = wall.GetRandomFile(wall.GetWallpaperSavePath())
		}
		rw.SetFromFile(file)
	case "baidu":
		words := wall.GetRandomWord(setting.Baidu.Word)
		url, filename := wall.GetBaiduImageURL(words)
		if setting.Baidu.Download == false {
			rw.SetFromURL(url)
		} else {
			wall.DownloadImage(url, filename)
			rw.SetFromFile(wall.GetWallpaperSavePath() + filename)
		}
	case "zol":
		sort := wall.GetRandomWord(setting.Zol.Sort)
		url, filename := wall.GetZolImageURL(sort)
		if setting.Zol.Download == false {
			rw.SetFromURL(url)
		} else {
			wall.DownloadImage(url, filename)
			rw.SetFromFile(wall.GetWallpaperSavePath() + filename)
		}
	case "sogou":
		sort := wall.GetRandomWord(setting.Sogou.Sort)
		url, filename := wall.GetSogouImageURL(sort)
		if setting.Sogou.Download == false {
			rw.SetFromURL(url)
		} else {
			wall.DownloadImage(url, filename)
			rw.SetFromFile(wall.GetWallpaperSavePath() + filename)
		}
	default:
		log.Fatal("配置错误")
	}
}
