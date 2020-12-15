package utility

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// 请求文件
func Download(url string, filepath, filename string) {
	body, err := HTTPResponse(url)
	if body == nil || err != nil {
		log.Println(err)
		return
	}
	defer body.Close()
	data, _ := ioutil.ReadAll(body)

	p := filepath + "/" + filename + ".jpg"
	f, err := os.Create(p)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	f.Write(data)
}

// 判断文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 判断HTTP Response有效性
func HTTPResponse(url string) (io.ReadCloser, error) {
	var (
		resp *http.Response
		err  error
	)

	for retry := 1; retry > 0; retry-- {
		resp, err = http.Get(url)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("status code error:", resp.StatusCode, resp.Status)
		return nil, err
	}

	return resp.Body, nil
}
