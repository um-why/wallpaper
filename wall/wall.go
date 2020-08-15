package wall

import (
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"github.com/reujab/wallpaper"
)

func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return filepath.Dir(path)
}

func GetWallpaperSavePath() string {
	return GetCurrentPath() + "/.wall/"
}

func OpenLog(file string) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}

func GetRandomWord(words []string) string {
	if len(words) == 0 {
		return ""
	}

	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(len(words))
	return words[randIndex]
}

func DownloadImage(url string, filename string) {
	defer SetRandomWall()

	filepath := GetWallpaperSavePath()

	dirExists, err := exists(filepath)
	if err != nil {
		panic(err)
	}

	if !dirExists {
		err = os.Mkdir(filepath, 0777)

		if err != nil {
			panic(err)
		}
	}

	imageFilePath := filepath + filename
	fileExists, err := exists(imageFilePath)
	if !fileExists {
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		i := len(url) - 1
		for i >= 0 && url[i] != '.' {
			i--
		}

		file, err := os.Create(imageFilePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			panic(err)
		}
	} else {
		log.Println("图文已存在，跳过下载")
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func GetRandomFile(path string) string {
	files, _ := ioutil.ReadDir(path)
	if len(files) == 0 {
		log.Fatal("未找到本地壁纸图片")
	}

	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(len(files))
	log.Println("随机获取的壁纸文件为:" + files[randIndex].Name())
	return path + files[randIndex].Name()
}

func SetRandomWall() {
	if err := recover(); err != nil {
		file := GetRandomFile(GetWallpaperSavePath())
		wallpaper.SetFromFile(file)
		log.Fatal(err)
	}
}
