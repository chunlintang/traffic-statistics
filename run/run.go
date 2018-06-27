package main

import (
	"flag"
)

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

			}
		}
	}
}

func main() {
	total := flag.Int("total", 100, "how many rows logs")
	// this nginx access is my computer's path,you should use your
	filePath := flag.String("filePath", "/usr/local/nginx/logs/dev.dig.com_access.log", "file path")
	flag.Parse()

	//fmt.Println(*total, *filePath)

	res := ruleResource()
	list := buildUrl(res)

	logStr := ...
	for i := 0; i <= *total; i++ {

	}
}
