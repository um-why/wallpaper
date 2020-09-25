package wall

import (
	"github.com/reujab/wallpaper"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
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

	if strings.Index(filename, "/") != -1 {
		secondPath := filepath + filename[:strings.Index(filename, "/")]
		dirExists, err = exists(secondPath)
		if err != nil {
			panic(err)
		}
		if !dirExists {
			err = os.Mkdir(secondPath, 0777)
			if err != nil {
				panic(err)
			}
		}
	}

	imageFilePath := filepath + filename
	fileExists, _ := exists(imageFilePath)
	if fileExists {
		log.Println("图片已存在，跳过下载")
		return
	}

	fileExists, _ = downloadexists(imageFilePath)
	if fileExists {
		log.Println("图片已存在，跳过下载.")
		return
	}

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

func downloadexists(path string) (bool, error) {
	fileName := path[strings.LastIndex(path, "/")+1:]

	paths := GetWallpaperSavePath()

	files, _ := ioutil.ReadDir(paths)
	if len(files) == 0 {
		return false, nil
	}

	var isExist = false
	for _, file := range files {
		if file.IsDir() == false {
			if file.Name() == fileName {
				isExist = true
				break
			}
		}

		filepath.Walk(paths+"/"+file.Name(), func(spath string, fi os.FileInfo, err error) error {
			if fi.IsDir() {
				return nil
			}
			if err != nil {
				return nil
			}
			if isExist {
				return filepath.SkipDir
			}
			if fi.Name() == fileName {
				isExist = true
				return filepath.SkipDir
			}
			return nil
		})
	}
	if isExist {
		return true, nil
	} else {
		return false, nil
	}
}

func GetRandomFile(paths string) string {
	files, _ := ioutil.ReadDir(paths)
	if len(files) == 0 {
		log.Fatal("未找到本地壁纸图片")
	}

	var lists []string

	for _, file := range files {
		if file.IsDir() == false {
			lists = append(lists, paths+file.Name())
			continue
		}

		filepath.Walk(paths+"/"+file.Name(), func(path string, fi os.FileInfo, err error) error {
			if fi.IsDir() {
				if fi.Name() == ".del" {
					return filepath.SkipDir
				} else {
					return nil
				}
			}
			if err != nil {
				return err
			}
			lists = append(lists, path)
			return nil
		})
	}

	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(len(lists))
	log.Println("随机获取的壁纸文件为:" + lists[randIndex])
	return lists[randIndex]
}

func SetRandomWall() {
	if err := recover(); err != nil {
		file := GetRandomFile(GetWallpaperSavePath())
		wallpaper.SetFromFile(file)
		log.Fatal(err)
	}
}
