package main

import (
	"encoding/json"
	"fmt"
	"my_instagram_follow/src/instagram"
	"my_instagram_follow/src/myexport"
	"os"
	"sync"
	"time"
)

func printHelp() {
	fmt.Println("Usage: ")
	fmt.Print("  ./my_instagram_follow <sessionid_cookie>\n\n")
	fmt.Print("Files are saved in `record/{todays_date}`")
}

func initDirectory() {
	var path string = fmt.Sprintf("record/%s", time.Now().Format("2006-01-02"))

	os.Mkdir("record", os.ModePerm)
	os.Mkdir(path, os.ModePerm)
}

func core(cookie string) {
	var wg sync.WaitGroup = sync.WaitGroup{}
	var path string = fmt.Sprintf("record/%s", time.Now().Format("2006-01-02"))

	initDirectory()
	wg.Add(2)
	go func() {
		fmt.Println("[ ] Fetching followers list")
		followers_list := instagram.GetFollowers(cookie)
		followers_json, _ := json.MarshalIndent(followers_list, "", "    ")
		myexport.MyWriteToFileJson(path+"/followers.json", followers_json)
		fmt.Println("[X] Followers list done!")
		defer wg.Done()

	}()
	go func() {
		fmt.Println("[ ] Fetching following list")
		following_list := instagram.GetFollowing(cookie)
		following_json, _ := json.MarshalIndent(following_list, "", "    ")
		myexport.MyWriteToFileJson(path+"/following.json", following_json)
		fmt.Println("[X] Following list done!")
		defer wg.Done()
	}()
	wg.Wait()
}

func main() {

	if (len(os.Args) != 2) || (os.Args[1] == "-h") {
		printHelp()
	} else {
		core(os.Args[1])
	}
}
