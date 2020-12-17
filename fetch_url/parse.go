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
		err = fmt.Errorf("http连接失败！")
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
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
		ok := addToMap(selection, resultList, i, "href")
		if ok && useDepth > currDepth {
			parseUrl(url, useDepth, currDepth, resultList)
		}
	})
	//iframe链接
	document.Find("iframe").Each(func(i int, selection *goquery.Selection) {
		ok := addToMap(selection, resultList, i, "src")
		if ok && useDepth > currDepth {
			parseUrl(url, useDepth, currDepth, resultList)
		}
	})

	//link
	document.Find("link").Each(func(i int, selection *goquery.Selection) {
		addToMap(selection, resultList, i, "href")
	})

	//脚本链接
	document.Find("script").Each(func(i int, selection *goquery.Selection) {
		addToMap(selection, resultList, i,"src")
	})

	//http[s]?://[a-zA-Z0-9\/.]

	//audio标签
	document.Find("audio ").Each(func(i int, selection *goquery.Selection) {
		addToMap(selection, resultList, i, "src")
	})

	//image标签
	document.Find("img").Each(func(i int, selection *goquery.Selection) {
		addToMap(selection, resultList, i, "src")
	})

	//video标签
	document.Find("video").Each(func(i int, selection *goquery.Selection) {
		addToMap(selection, resultList, i, "src")
	})
}
func addToMap(selection *goquery.Selection, result *map[string]bool, _ int, attr string) bool {
	url, exists := selection.Attr(attr)
	hasPrefix := strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")
	if exists && hasPrefix {
		(*result)[url] = true
		return true
	}
	return false
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