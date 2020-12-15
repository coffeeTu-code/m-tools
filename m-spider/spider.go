package m_spider

import (
	"fmt"
	"log"
	"m-tools/m-spider/utility"
	"net/url"
	"sort"
	"strings"
)

func Spider(filePath, webPath string, urls map[string]string) {
	var targetUrls []string
	for targetUrl, _ := range urls {
		targetUrls = append(targetUrls, targetUrl)
	}
	sort.Strings(targetUrls)

	for order, targetUrl := range targetUrls {
		fmt.Println(" 次序：", order+1, "/", len(targetUrls), " ----- ", urls[targetUrl], " ----- ", targetUrl)
		job(filePath, webPath, targetUrl)
	}
}

func job(filePath, webPath, targetUrl string) {
	log.Println(webPath, targetUrl)

	urlObj, err := url.Parse(targetUrl)
	if err != nil {
		log.Println(err)
		return
	}
	switch {
	case strings.Contains(targetUrl, webPath+"/t"):
		mote_name, urls := utility.FindDocument(webPath, targetUrl, "t")
		fmt.Println("出境模特：", mote_name, len(urls))
		if mote_name == "" {
			return
		}
		download_path := strings.Replace(urlObj.EscapedPath()+"__"+mote_name, "/", "", -1)
		for i, val := range urls {
			fmt.Println("[", i+1, "/", len(urls), "]")
			download(filePath, val, download_path)
		}
	case strings.Contains(targetUrl, webPath+"/x"):
		mote_name, urls := utility.FindDocument(webPath, targetUrl, "x")
		fmt.Println("出境模特：", mote_name, len(urls))
		download_path := mote_name
		if download_path == "" {
			download_path = "default"
		}
		for i, val := range urls {
			fmt.Println("[", i+1, "/", len(urls), "]")
			download(filePath, val, download_path)
		}
	case strings.Contains(targetUrl, webPath+"/a"):
		download(filePath, targetUrl, "default")
	}
	log.Println("done...")
}

func download(filePath, url string, mote_name string) {
	document := utility.NewDocument(filePath, mote_name, url)
	if document != nil {
		document.FindAll()
		document.SaveContents(filePath)
	}
}
