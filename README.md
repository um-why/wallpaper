# wallpaper
一个简单的go程序，从Bing、Baidu、ZOL抓取图片并设置为壁纸

## 简介

1. 设置桌面壁纸，最少系统资源占用；绿色、无广告、无监控；
2. bing壁纸质量好；本程序自动下载，长期积累下来，(即便没有网络)您的壁纸就是完美极致的风景大片；
3. baidu壁纸自动适配本机屏幕分辨率；(不建议下载，部分图片可能模糊)
4. zol壁纸自动挑选适配本机屏幕分辨率图片；(部分图片可能宽度不够)

## 启动壁纸切换

#### 1. 获取程序

- 去 [https://gitee.com/um-why/wallpaper/releases](https://gitee.com/um-why/wallpaper/releases) 直接下载构建后的合适版本的zip压缩包。

- 或克隆源代码 `git clone https://gitee.com/um-why/wallpaper.git` ，自己构建 `go build -ldflags "-H windowsgui" main.go` 。

- 将下载的zip压缩包，解压到任意文件夹中(建议解压到非系统盘文件目录中)。

#### 2. 程序配置

配置文件是 `config.json` ，和主程序在同级目录中。

配置文件示例：

```json
{
  "sort": "bing",
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
    "download": false
  },
  "log": true
}
```
配置文件采用JSON格式，详细说明如下:

```
第一级设置项：

sort: 可设置为 baidu / bing / zol ， 壁纸来源
bing: sort = bing 时本项生效
baidu: sort = baidu 时本项生效
bing: sort = zol 时本项生效
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

关于百度壁纸搜索关键词的设置，您可以查看此处文档说明 [百度壁纸搜索关键词](readme/baidu.md "百度壁纸关键词设置")，了解更多自主设置选项。

```
第二级zol设置项:

zol.sort: ZOL壁纸分类；
zol.download: 可设置为 true / false，ZOL壁纸图片是否下载
```

关于ZOL壁纸分类的设置，我也提供了文档说明，有兴趣可以查看 [ZOL壁纸分类](readme/zol.md "ZOL壁纸分类设置")

#### 3. 程序运行

- 运行本程序，即可完成壁纸的更改；程序无任何窗口，壁纸设置后即自动退出，无后台守护；
- 您可以根据自己对壁纸切换频次的选择，自主设定程序的运行频次；
```
如果您喜欢每日更换一次桌面壁纸；建议设置为桌面壁纸来源于bing、不随机即可，程序配置如下：

{
  "sort": "bing",
  "bing": {
    "mode": "today"
  }
}
```
也可考虑开启随机，经过一小段时间的积累，您就可以拥有一个自主的非常漂亮的桌面壁纸素材库。开启随机，只需要将此处程序配置中的 `"mode": "today"` 修改为 `"mode": "random"`

------

```
如果您喜欢一日更换多次桌面壁纸；可设置桌面壁纸来源于 baidu 或 zol，程序配置如下：

{
  "sort": "baidu",
  "baidu": {
    "word": [
      "壁纸"
    ],
    "download": false
  }
}

{
  "sort": "zol",
  "zol": {
    "sort": [
      "全部"
    ],
    "download": false
  }
}
```

#### 4. 自动更换壁纸

- 您可以**将程序快捷方式加入系统的启动项**，既可实现开机自动更换桌面壁纸；

- 也可**将程序加入系统计划任务，自由设置运行间隔**，即实现更绚丽的桌面壁纸更换。

------

快去试试吧！喜欢您能喜欢这个程序。