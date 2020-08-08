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
	"syscall"
	"time"
)

func getWinScreenSize(nIndex int) int {
	rs, _, _ := syscall.NewLazyDLL("User32.dll").NewProc("GetSystemMetrics").Call(uintptr(nIndex))
	return int(rs)
}

func GetBaiduImageURL(word string) (imgURL string, imgFilename string) {
	searchUrl := "https://image.baidu.com/search/index?tn=baiduimage&word="
	searchUrl += url.QueryEscape(word)
	searchUrl += "&width=" + strconv.Itoa(getWinScreenSize(0))
	searchUrl += "&height=" + strconv.Itoa(getWinScreenSize(1))

	response, err := http.Get(searchUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	re := regexp.MustCompile("  \"objURL\":\".+?\"")
	htmlData, _ := ioutil.ReadAll(response.Body)
	objectUrl := re.FindAllString(string(htmlData), -1)
	if len(objectUrl) <= 0 {
		log.Fatal("未找到百度壁纸图片")
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

	return imgURL, imgFilename + imgExt
}
