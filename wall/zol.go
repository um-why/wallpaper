package wall

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func sortDict(sort string) string {
	sorts := map[string]string{
		"风景":  "fengjing",
		"动漫":  "dongman",
		"美女":  "meinv",
		"创意":  "chuangyi",
		"卡通":  "katong",
		"汽车":  "qiche",
		"游戏":  "youxi",
		"可爱":  "keai",
		"明星":  "mingxing",
		"建筑":  "jianzhu",
		"植物":  "zhiwu",
		"动物":  "dongwu",
		"静物":  "jingwu",
		"影视":  "yingshi",
		"车模":  "chemo",
		"体育":  "tiyu",
		"模特":  "model",
		"手抄报": "shouchaobao",
		"美食":  "meishi",
		"星座":  "xingzuo",
		"节日":  "jieri",
		"品牌":  "pinpai",
		"背景":  "beijing",
		"其他":  "qita"}
	if _, ok := sorts[sort]; ok {
		return sorts[sort]
	}
	return "pc"
}

func getSuitablePixel() string {
	width := GetWinScreenSize(0)
	height := GetWinScreenSize(1)
	pixel := strconv.Itoa(width) + "x" + strconv.Itoa(height)
	pixels := [12]string{"4096x2160", "2880x1800", "2560x1600", "2560x1440", "1920x1080", "1680x1050", "1600x900", "1440x900", "1366x768", "1280x1024", "1280x800", "1024x768"}
	isExists := false
	for _, pix := range pixels {
		if pix == pixel {
			isExists = true
		}
	}
	if isExists == true {
		return pixel
	} else {
		return ""
	}
}

func GetZolImageURL(sort string) (imgURL string, imgFilename string) {
	defer SetRandomWall()

	sort = sortDict(sort)
	homeUrl := ""
	if sort == "pc" {
		homeUrl = "http://desk.zol.com.cn/"
	} else {
		homeUrl = "http://desk.zol.com.cn/" + sort + "/"
	}

	response, err := http.Get(homeUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	re := regexp.MustCompile("photo-list-padding\"><a class=\"pic\" href=\".+?\"")
	htmlData, _ := ioutil.ReadAll(response.Body)
	objectUrl := re.FindAllString(string(htmlData), -1)
	if len(objectUrl) <= 0 {
		panic("未找到ZOL " + sort + " 分类的壁纸图片")
	}

	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(len(objectUrl))

	detailUrl := objectUrl[randIndex]
	detailUrl = detailUrl[41 : len(detailUrl)-1]
	detailUrl = "http://desk.zol.com.cn" + detailUrl

	response, err = http.Get(detailUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	re = regexp.MustCompile("oriSize\":\".+?\"}")
	htmlData, _ = ioutil.ReadAll(response.Body)
	objectUrl = re.FindAllString(string(htmlData), -1)
	if len(objectUrl) <= 0 {
		panic("ZOL壁纸图片解析错误")
	}

	rand.Seed(time.Now().Unix())
	randIndex = rand.Intn(len(objectUrl))

	imgURL = objectUrl[randIndex]
	if strings.Index(imgURL, "##SIZE##") == -1 {
		panic("ZOL壁纸图片规则错误")
	}
	size := getSuitablePixel()
	if size == "" {
		size = imgURL[10:19]
		if strings.Index(size, "\"") != -1 {
			size = size[:8]
		}
	}
	imgURL = strings.Replace(imgURL, "\\", "", -1)
	imgURL = imgURL[strings.Index(imgURL, "imgsrc\":\"")+9:]
	imgURL = imgURL[:len(imgURL)-2]
	imgURL = strings.Replace(imgURL, "##SIZE##", size, 1)
	log.Println("本次随机的ZOL壁纸图片URL:" + imgURL)

	imgExt := imgURL[strings.LastIndexAny(imgURL, "."):]

	h := md5.New()
	h.Write([]byte(imgURL))
	cipherStr := h.Sum(nil)
	imgFilename = hex.EncodeToString(cipherStr)

	return imgURL, "zol/" + imgFilename + imgExt
}
