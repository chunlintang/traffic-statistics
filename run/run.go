package main

import (
	"flag"
	"strings"
	"strconv"
	"net/url"
	"io/ioutil"
	"math/rand"
	"time"
)

var uaList = []string{
	// Chrome｜谷歌浏览器
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.87 Safari/537.36",
	// Firefox｜火狐浏览器
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:46.0) Gecko/20100101 Firefox/46.0",
	// Opera｜欧朋浏览器
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.87 Safari/537.36 OPR/37.0.2178.32",
	// Safari｜苹果浏览器
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.57.2 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
	// 360安全浏览器/360极速浏览器
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36",
	// 微软 Edge 浏览器
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Safari/537.36 Edge/13.10586",
	// Internet Explorer 11 浏览器
	"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko",
	// Internet Explorer 10 浏览器
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)",
	// Internet Explorer 9 浏览器
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)",
	// Internet Explorer 8 浏览器
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0)",
	// 百度浏览器
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.106 BIDUBrowser/8.3 Safari/537.36",
	// 遨游浏览器
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Maxthon/4.9.2.1000 Chrome/39.0.2146.0 Safari/537.36",
	// QQ浏览器
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36 Core/1.47.277.400 QQBrowser/9.4.7658.400",
	// UC浏览器电脑版
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 UBrowser/5.6.12150.8 Safari/537.36",
	// 搜狗浏览器
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36 SE 2.X MetaSr 1.0",
	// 猎豹浏览器
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.154 Safari/537.36 LBBROWSER",
	// 世界之窗浏览器
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 Safari/537.36 TheWorld 7",
}

type resource struct {
	url    string
	target string
	start  int
	end    int
}

func ruleResource() []resource {
	var res []resource
	// index page
	r1 := resource{
		url:    "http://dev.dotexample.com/example/views",
		target: "",
		start:  0,
		end:    0,
	}
	// list page
	r2 := resource{
		url:    "http://dev.dotexample.com/example/views/list.html",
		target: "",
		start:  0,
		end:    0,
	}
	// detail page
	r3 := resource{
		url:    "http://dev.dotexample.com/example/views/detail.html",
		target: "",
		start:  0,
		end:    0,
	}
	// profile page
	r4 := resource{
		url:    "http://dev.dotexample.com/example/views/profile.html",
		target: "",
		start:  0,
		end:    0,
	}
	res = append(append(append(append(res, r1), r2), r3), r4)

	return res
}

func buildUrl(res []resource) []string {
	var list []string

	for _, r := range res {
		if len(r.target) == 0 {
			list = append(list, r.url)
		} else {
			for i := r.start; i <= r.end; i++ {
				urlStr := strings.Replace(r.url, r.target, strconv.Itoa(i), -1)
				list = append(list, urlStr)
			}
		}
	}

	return list
}

func makeLog(current, refer, ua string) string {
	u := url.Values{}
	u.Set("time", "1")
	u.Set("url", current)
	u.Set("refer", refer)
	u.Set("ua", ua)
	paramsStr := u.Encode()

	template := "127.0.0.1 - - [26/Jun/2018:23:17:44 +0800] \"GET /dig?{$paramsStr} HTTP/1.1\" 200 43 \"-\" \"{$ua}\" \"-\"\"
	log := strings.Replace(template, "{$paramsStr}", paramsStr, -1)
	log = strings.Replace(log, "{$ua}", ua, -1)

	return log
}

func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return max
	}
	return r.Intn(max-min) + min
}

func main() {
	total := flag.Int("total", 100, "how many rows logs")
	// this nginx access is my computer's path,you should use your
	filePath := flag.String("filePath", "/usr/local/nginx/logs/dev.dig.com_access.log", "file path")
	flag.Parse()

	//fmt.Println(*total, *filePath)

	res := ruleResource()
	list := buildUrl(res)

	for i := 0; i <= *total; i++ {
		currentUrl := list[randInt(0, len(list)-1)]
		referUrl := list[randInt(0, len(list)-1)]
		ua := uaList[randInt(0, len(uaList)-1)]

		logStr := makeLog(currentUrl, referUrl, ua)
		ioutil.WriteFile(*filePath, []byte(logStr), 0644)
	}
}
