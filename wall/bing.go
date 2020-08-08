package wall

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func GetBingImageURL()(imgURL string, imgFilename string){
	response, err := http.Get("https://www.bing.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	re := regexp.MustCompile("data-ultra-definition-src=\".+?\\.jpg")
	htmlData, _ := ioutil.ReadAll(response.Body)
	imgURL = re.FindString(string(htmlData))
	imgURL = imgURL[28:]
	imgFilename = imgURL[6:]

	if 0 < len(imgURL) {
		log.Println("今日Bing桌面URL:" + imgURL)
		return "https://cn.bing.com/" + imgURL, imgFilename
	} else {
		return "", ""
	}
}
