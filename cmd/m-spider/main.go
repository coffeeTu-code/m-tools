package main

import (
	"fmt"
	m_spider "m-tools/m-spider"
	"m-tools/m-spider/configer"
)

func main() {
	fmt.Println("Hello Spider")

	m_spider.Spider(configer.FilePath, configer.WebPath, configer.AllDoc)
}
