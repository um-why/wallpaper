## 程序配置文件详细说明

配置文件示例：

```json
{
  "sort": [
    "bing",
    "baidu",
    "zol",
    "sogou"
  ],
  "bing": {
    "mode": "random"
  },
  "baidu": {
    "word": [
      "壁纸",
      "壁纸 不同风格 唯美",
      "壁纸 卡通动漫 海贼王",
      "壁纸 动物 狗狗",
      "壁纸 影视 变形金刚"
    ],
    "download": false
  },
  "zol": {
    "sort": [
      "全部",
      "风景",
      "动漫",
      "美女",
      "创意"
    ],
    "download": true
  },
  "sogou": {
    "sort": [
      "美女",
      "动漫",
      "风景",
      "小清新",
      "动态"
    ],
    "download": true
  },
  "log": true
}
```
配置文件采用JSON格式，详细说明如下:

```
第一级设置项：

sort: 可设置为 baidu / bing / zol / sogou ， 壁纸来源
bing: sort = bing 时本项生效
baidu: sort = baidu 时本项生效
bing: sort = zol 时本项生效
sogou: sort = sogou 时本项生效
log: 可设置为 true / false ， 是否记录程序运行日志
```

```
第二级bing设置项:

bing.mode: 可设置为 today / random ；
该项等于 today 时，壁纸为bing当天的网站背景；
该项等于 random 时，壁纸为自程序运行以来程序所有下载的壁纸图片进行随机切换；
```

```
第二级baidu设置项:

baidu.word: 百度壁纸搜索关键词；
baidu.download: 可设置为 true / false，百度壁纸图片是否下载
```

关于百度壁纸搜索关键词的设置，您可以查看此处文档说明 [百度壁纸搜索关键词](baidu.md "百度壁纸关键词设置")，了解更多自主设置选项。

```
第二级zol设置项:

zol.sort: ZOL壁纸分类；
zol.download: 可设置为 true / false，ZOL壁纸图片是否下载
```

关于ZOL壁纸分类的设置，我也提供了文档说明，有兴趣可以查看 [ZOL壁纸分类](zol.md "ZOL壁纸分类设置")

```
第二级sogou设置项:
sogou.sort: SOGOU壁纸分类；
sogou.download: 可设置为 true / false，SOGOU壁纸图片是否下载
```

关于SOGOU壁纸分类的设置，有兴趣可以查看 [SOGOU壁纸分类](sogou.md "SOGOU壁纸分类设置")