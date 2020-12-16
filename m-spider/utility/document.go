package utility

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"os"
	"path"
	"strings"
)

type document struct {
	url      string
	filepath string             //保存content的文件路径
	content  map[string]content //key=src
	pages    map[string]pages   //key=page
}

type content struct {
	src string
	alt string
}

type pages struct {
	page string
	href string
}

func NewDocument(filePath, mote_name, urlStr string) *document {
	urlObj, err := url.Parse(urlStr)
	if urlStr == "" || err != nil {
		log.Println(err)
		return nil
	}

	//判断文件夹是否存在,并创建
	filepath := path.Join(filePath+mote_name, strings.Replace(urlObj.EscapedPath(), "a/", "", -1))
	if !PathExists(filepath) {
		if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
			log.Println(err)
			return nil
		}
	} else {
		log.Println("document 已存在", urlStr, ",save path = ", filepath)
		return nil
	}

	log.Println("new document = ", urlStr, ",save path = ", filepath)
	return &document{
		url:      urlObj.String(),
		filepath: filepath,
		content:  map[string]content{},
		pages:    map[string]pages{},
	}
}

func FindDocument(webPath, url string, feature string) (mote string, a []string) {

	var urls = map[string]bool{url: true}
	if doc := newDocument(url); doc != nil {
		// Find all page
		doc.Find("center").Each(func(i int, s *goquery.Selection) {
			s.Find("a").Each(func(i int, s *goquery.Selection) {
				href, _ := s.Attr("href")
				if strings.Contains(href, feature+"/") {
					if !strings.Contains(href, "http") {
						href = webPath + href
					}
					urls[href] = true
				}
			})
		})
		// Find mote 模特
		doc.Find("div").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			if val, exist := s.Attr("class"); !exist || val != "left" {
				return
			}
			s.Find("img").Each(func(i int, s *goquery.Selection) {
				alt, _ := s.Attr("alt")
				mote = strings.Split(strings.Replace(strings.Replace(alt, " ", "", -1), "/", "", -1), "、")[0]
			})
		})
	}

	var documents = map[string]bool{}
	for k, _ := range urls {
		doc := newDocument(k)
		if doc == nil {
			continue
		}
		doc.Find("div").Each(func(i int, s *goquery.Selection) {
			if val, exist := s.Attr("class"); !exist || val != "hezi" {
				return
			}
			s.Find("a").Each(func(i int, s *goquery.Selection) {
				href, _ := s.Attr("href")
				if strings.Contains(href, "a") {
					documents[href] = true
				}
			})
		})
	}

	var ret []string
	for k, _ := range documents {
		ret = append(ret, k)
	}
	return mote, ret
}

// 请求html页面
func newDocument(url string) (document *goquery.Document) {
	body, err := HTTPResponse(url)
	if body == nil || err != nil {
		log.Println(err)
		return
	}
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println(err)
		return
	}
	return doc
}


func (this *document) FindAll() {
	for this.url != "" {
		this.Find()

		if next := this.Next(); next == "" || next == this.url {
			this.url = ""
		} else {
			this.url = next
		}
	}
}

func (this *document) Find() {
	if doc := newDocument(this.url); doc != nil {
		this.FindContents(doc)
		this.FindPages(doc)
	}
}

func (this *document) Next() string {
	if len(this.pages) == 0 {
		return ""
	}
	page, ok := this.pages["下一页"]
	if !ok {
		return ""
	}
	return page.href
}

func (this *document) FindContents(document *goquery.Document) {
	if this.content == nil {
		this.content = make(map[string]content)
	}

	// Find the review items
	document.Find("div").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		if val, exist := s.Attr("class"); !exist || val != "content" {
			return
		}
		s.Find("img").Each(func(i int, s *goquery.Selection) {
			src, _ := s.Attr("src")
			alt, _ := s.Attr("alt")
			class, _ := s.Attr("class")
			if class != "tupian_img" {
				return
			}
			this.content[src] = content{
				src: src,
				alt: strings.Replace(strings.Replace(alt, " ", "", -1), "/", "", -1),
			}
		})
	})
}

func (this *document) FindPages(document *goquery.Document) {
	if this.pages == nil {
		this.pages = make(map[string]pages)
	}

	// Find the review items
	document.Find("div").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		if val, exist := s.Attr("id"); !exist || val != "pages" {
			return
		}
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			page := s.Text()
			this.pages[page] = pages{
				page: page,
				href: href,
			}
		})
	})
}

func (this *document) SaveContents(filePath string) {
	fmt.Print("content: ", len(this.content), ", download...")
	var i, percent = 0, 0
	for _, content := range this.content {
		xiezhen := strings.Split(this.filepath, "/")[len(strings.Split(filePath, "/"))]
		content.alt = xiezhen + "__" + content.alt
		Download(content.src, this.filepath, content.alt)

		i++
		if tmp := i * 100 / len(this.content); tmp%10 == 0 && tmp != percent {
			percent = tmp
			fmt.Print(">>", percent, "% ")
		}
	}
	fmt.Println()
}
