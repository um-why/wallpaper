package wall

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetBaiduImageURL(word string) (imgURL string, imgFilename string) {
	defer SetRandomWall()

	if word == "" {
		word = "壁纸"
	}

	searchUrl := "https://image.baidu.com/search/index?tn=baiduimage&word="
	searchUrl += url.QueryEscape(word)
	searchUrl = strings.Replace(searchUrl, "+", "%20", -1)
	width := GetWinScreenSize(0)
	if width != 0 {
		searchUrl += "&width=" + strconv.Itoa(width)
	}
	height := GetWinScreenSize(1)
	if height != 0 {
		searchUrl += "&height=" + strconv.Itoa(height)
	}

	response, err := http.Get(searchUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	re := regexp.MustCompile("  \"objURL\":\".+?\"")
	htmlData, _ := ioutil.ReadAll(response.Body)
	objectUrl := re.FindAllString(string(htmlData), -1)
	if len(objectUrl) <= 0 {
		panic("未找到搜索关键词: " + word + " 的百度壁纸图片")
	}

	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(len(objectUrl))

	imgURL = objectUrl[randIndex]
	imgURL = imgURL[12 : len(imgURL)-1]
	log.Println("本次随机的百度壁纸图片URL:" + imgURL)

	imgExt := imgURL[strings.LastIndexAny(imgURL, "."):]

	h := md5.New()
	h.Write([]byte(imgURL))
	cipherStr := h.Sum(nil)
	imgFilename = hex.EncodeToString(cipherStr)

	return imgURL, "baidu/" + imgFilename + imgExt
}
