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

func GetSogouImageURL(sort string) (imgURL string, imgFilename string) {
	defer SetRandomWall()

	homeUrl := "https://pic.sogou.com/api/pic/searchList?tagQSign=&forbidqc="
	homeUrl += "&entityid=&preQuery=&rawQuery=&queryList=&ie=&query=" + url.QueryEscape("电脑壁纸")
	homeUrl += "&mode=1&start=0&xml_len=48&_=" + strconv.Itoa(int(time.Now().Unix()))
	response, err := http.Get(homeUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	htmlData, _ := ioutil.ReadAll(response.Body)
	_html := string(htmlData)
	_html = _html[strings.Index(_html, ",\"tag\":\"[")+9:]
	_html = _html[:strings.Index(_html, "]\",\"")]
	sort, sortSign := getSortSign(_html, sort)

	imgUrl := "https://pic.sogou.com/api/pic/searchList?tagQSign="
	imgUrl += url.QueryEscape(sort + "," + sortSign)
	imgUrl += "&forbidqc=&entityid=&preQuery=&rawQuery=&queryList=&ie=&query="
	imgUrl += url.QueryEscape("电脑壁纸")
	imgUrl += "&mode=6&start=0&xml_len=48&tagQ=" + url.QueryEscape(sort)
	imgUrl += "&_=" + strconv.Itoa(int(time.Now().Unix()))

	response, err = http.Get(imgUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	re := regexp.MustCompile("oriPicUrl\":\".+?\"")
	htmlData, _ = ioutil.ReadAll(response.Body)
	objectUrl := re.FindAllString(string(htmlData), -1)
	if len(objectUrl) <= 0 {
		panic("未找到分类为: " + sort + " 的搜狗图片")
	}

	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(len(objectUrl))
	imgURL = objectUrl[randIndex]
	imgURL = imgURL[12 : len(imgURL)-1]
	log.Println("本次随机的搜狗壁纸图片URL:" + imgURL)

	imgExt := imgURL[strings.LastIndexAny(imgURL, "."):]

	h := md5.New()
	h.Write([]byte(imgURL))
	cipherStr := h.Sum(nil)
	imgFilename = hex.EncodeToString(cipherStr)

	return imgURL, "sogou/" + imgFilename + imgExt
}

func getSortSign(sortHtml string, sortLable string) (string, string) {
	_index := strings.Index(sortHtml, sortLable)
	if _index == -1 {
		sort := sortHtml[strings.Index(sortHtml, "[\\\"")+3 : strings.Index(sortHtml, "\\\",")]
		sortSign := sortHtml[strings.Index(sortHtml, ",\\\"")+3 : strings.Index(sortHtml, "\\\"],")]
		return sort, sortSign
	} else {
		_html := sortHtml[_index:]
		sortSign := _html[strings.Index(_html, ",\\\"")+3 : strings.Index(_html, "\\\"],")]
		return sortLable, sortSign
	}
}
