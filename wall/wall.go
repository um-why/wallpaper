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
)

func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return filepath.Dir(path)
}

func OpenLog(file string) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}

func DownloadImage(url string, filepath string, filename string) {
	dirExists, err := exists(filepath)
	if err != nil {
		log.Fatal("寻找壁纸目录错误\n", err)
	}

	if !dirExists {
		err = os.Mkdir(filepath, 0777)

		if err != nil {
			log.Fatal("壁纸目录创建错误\n", err)
		}
	}

	imageFilePath := filepath + filename
	fileExists, err := exists(imageFilePath)
	if !fileExists {
		response, err := http.Get(url)
		if err != nil {
			log.Fatal("图片下载错误\n", err)
		}
		defer response.Body.Close()

		i := len(url) - 1
		for i >= 0 && url[i] != '.' {
			i--
		}

		file, err := os.Create(imageFilePath)
		if err != nil {
			log.Fatal("图片文件创建错误\n", err)
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal("图片保存错误\n", err)
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
	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(len(files))
	log.Println("随机获取的壁纸文件为:" + files[randIndex].Name())
	return path + files[randIndex].Name()
}
