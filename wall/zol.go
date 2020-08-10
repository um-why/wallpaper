package wall

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func sortDict(sort string) string {
	sorts := map[string]string{
		"风景":"fengjing",
		"动漫":"dongman",
		"美女":"meinv",
		"创意":"chuangyi",
		"卡通":"katong",
		"汽车":"qiche",
		"游戏":"youxi",
		"可爱":"keai",
		"明星":"mingxing",
		"建筑":"jianzhu",
		"植物":"zhiwu",
		"动物":"dongwu",
		"静物":"jingwu",
		"影视":"yingshi",
		"车模":"chemo",
		"体育":"tiyu",
		"模特":"model",
		"手抄报":"shouchaobao",
		"美食":"meishi",
		"星座":"xingzuo",
		"节日":"jieri",
		"品牌":"pinpai",
		"背景":"beijing",
		"其他":"qita"}
	if _,ok:=sorts[sort];ok{
		return sorts[sort]
	}
	return "pc"
}


func getUrl(sort string) string{
	sort = sortDict(sort)
	if sort == "pc" {
		return "http://desk.zol.com.cn/4096x2160/"
	}else{
		return "http://desk.zol.com.cn/"+sort+"/4096x2160/"
	}
}

func GetZolImageURL(sort string)  (imgURL string, imgFilename string){
	fmt.Println(sort)
	homeUrl := getUrl(sort)
	fmt.Println(homeUrl)

	response, err := http.Get(homeUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	re := regexp.MustCompile("photo-list-padding\"><a class=\"pic\" href=\".+?\"")
	htmlData, _ := ioutil.ReadAll(response.Body)
	objectUrl := re.FindAllString(string(htmlData), -1)
	if len(objectUrl) <= 0 {
		log.Fatal("未找到ZOL壁纸图片")
	}
	fmt.Println(objectUrl)

	return "", ""
}
