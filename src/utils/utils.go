package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func PrepareCookie(cookie string) string {
	return fmt.Sprintf("sessionid=%s", cookie)
}

func GetUserId(cookie string) string {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal("Could not parse cookie")
		}
	}()

	return strings.Split(cookie, "%")[0]
}

func DoRequest(req *http.Request) (bytes []byte) {
	var client = &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	} else if res.Header.Get("Content-type") != "application/json; charset=utf-8" {
		fmt.Fprint(os.Stderr, "Response is not a json. Please check your sessionid cookie.\n")
		os.Exit(1)
	}
	bytes, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		log.Fatal(errRead)
	}
	defer res.Body.Close()

	return
}
