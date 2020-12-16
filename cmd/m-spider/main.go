package main

import (
	"log"
	m_spider "m-tools/m-spider"
	"m-tools/m-spider/configer"
)

func main() {
	log.Println("Hello Spider")

	m_spider.Spider(configer.FilePath, configer.WebPath, configer.AllDoc)
}
