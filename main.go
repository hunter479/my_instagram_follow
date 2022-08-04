package main

import (
	"encoding/json"
	"fmt"
	"my_instagram_follow/src/instagram"
	"my_instagram_follow/src/myexport"
	"os"
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
	initDirectory()
	var path string = fmt.Sprintf("record/%s", time.Now().Format("2006-01-02"))

	followers_list := instagram.GetFollowers(cookie)
	following_list := instagram.GetFollowing(cookie)
	followers_json, _ := json.MarshalIndent(followers_list, "", "    ")
	following_json, _ := json.MarshalIndent(following_list, "", "    ")
	myexport.MyWriteToFileJson(path+"/followers.json", followers_json)
	myexport.MyWriteToFileJson(path+"/following.json", following_json)
}

func main() {

	if (len(os.Args) != 2) || (os.Args[1] == "-h") {
		printHelp()
	} else {
		core(os.Args[1])
	}
}
