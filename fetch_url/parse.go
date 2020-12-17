package fetch_url

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
	"time"
)
const clientTimeout = time.Minute * 3
const maxDepth = 3
func fetch(url string) (*goquery.Document, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	client := &http.Client{}
	client.Timeout = clientTimeout
	resp, err := client.Do(req)
	if err != nil {
		//logs.Info("http连接失败！")
		err = fmt.Errorf("http连接失败！")
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		//logs.Info("Http返回状态为" + strconv.Itoa(resp.StatusCode))
		err = fmt.Errorf("Http返回状态为" + strconv.Itoa(resp.StatusCode))
		return nil, err
	}
	return goquery.NewDocumentFromReader(resp.Body)
}

func parseUrl(url string, useDepth, currDepth int, resultList *map[string]bool)  {
	document, err := fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	currDepth += 1
	//a链接
	document.Find("a").Each(func(i int, selection *goquery.Selection) {
		url, exists := selection.Attr("href")
		hasPrefix := strings.HasPrefix(url, "https://") || strings.HasPrefix(url,"http://")
		if exists && hasPrefix {
			(*resultList)[url] = true
		}
		if exists && hasPrefix && useDepth > currDepth {
			parseUrl(url, useDepth, currDepth, resultList)
		}
	})
	//iframe链接
	document.Find("iframe").Each(func(i int, selection *goquery.Selection) {
		url, exists := selection.Attr("src")
		hasPrefix := strings.HasPrefix(url, "https://") || strings.HasPrefix(url,"http://")
		if exists && hasPrefix {
			(*resultList)[url] = true
		}
		if exists && hasPrefix && useDepth > currDepth {
			parseUrl(url, useDepth, currDepth, resultList)
		}
	})

	//link
	document.Find("link").Each(func(i int, selection *goquery.Selection) {
		url, exists := selection.Attr("href")
		hasPrefix := strings.HasPrefix(url, "https://") || strings.HasPrefix(url,"http://")
		if exists && hasPrefix {
			(*resultList)[url] = true
		}
	})

	//脚本链接
	document.Find("script").Each(func(i int, selection *goquery.Selection) {
		url, exists := selection.Attr("src")
		hasPrefix := strings.HasPrefix(url, "https://") || strings.HasPrefix(url,"http://")
		if exists && hasPrefix {
			(*resultList)[url] = true
		}
	})

	//http[s]?://[a-zA-Z0-9\/.]

	//audio标签
	document.Find("audio ").Each(func(i int, selection *goquery.Selection) {
		url, exists := selection.Attr("src")
		hasPrefix := strings.HasPrefix(url, "https://") || strings.HasPrefix(url,"http://")
		if exists && hasPrefix {
			(*resultList)[url] = true
		}
	})

	//image标签
	document.Find("img").Each(func(i int, selection *goquery.Selection) {
		url, exists := selection.Attr("src")
		hasPrefix := strings.HasPrefix(url, "https://") || strings.HasPrefix(url,"http://")
		if exists && hasPrefix {
			(*resultList)[url] = true
		}
	})

	//video标签
	document.Find("video").Each(func(i int, selection *goquery.Selection) {
		url, exists := selection.Attr("src")
		hasPrefix := strings.HasPrefix(url, "https://") || strings.HasPrefix(url,"http://")
		if exists && hasPrefix {
			(*resultList)[url] = true
		}
	})
}

func Start(url string, useDepth int) *map[string]bool{
	t1 := time.Now()
	resultList := make(map[string]bool)
	if useDepth > maxDepth{
		useDepth = maxDepth
	}
	fmt.Println(100)
	parseUrl(url, useDepth, 0, &resultList)
	fmt.Println(time.Since(t1))
	return &resultList
}