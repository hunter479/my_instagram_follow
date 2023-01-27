package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func PrepareCookie(cookie string) string {
	return fmt.Sprintf("sessionid=%s", cookie)
}

func PrepareRequest(cookie string, url string) *http.Request {
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept-Language", "en-GB,en-Us;q=0.9,en;q=0.7")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("Referer", "https://i.instagram.com/")
	req.Header.Add("User-Agent", "Instagram 219.0.0.12.117 Android")
	req.Header.Add("cookie", cookie)

	return req
}

func DoRequest(req *http.Request) (bytes []byte) {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	} else if res.StatusCode != 200 {
		log.Fatalf("Response not 200.\nResponse code was %d\n", res.StatusCode)
	}
	bytes, errRead := io.ReadAll(res.Body)
	if errRead != nil {
		log.Fatal(errRead)
	}

	defer res.Body.Close()
	return bytes
}
