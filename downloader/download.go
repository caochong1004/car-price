package downloader

import (
	"github.com/axgle/mahonia"
	"github.com/caochong/studyGo/car-price/fake"
	"io"
	"log"
	"net/http"
)

func Get(url string) io.Reader  {
	client := http.Client{}
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		log.Printf("http.NewRequest err: %v\n", err)
	}

	req.Header.Add("User-Agent", fake.GetUserAgent())
	req.Header.Add("Referer","https://car.autohome.com.cn")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do err: %v\n", err)
	}
	mah := mahonia.NewDecoder("gbk")
	return mah.NewReader(resp.Body)
}
