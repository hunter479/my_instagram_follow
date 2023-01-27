package main

import (
	"encoding/json"
	"fmt"
	"log"
	"my_instagram_follow/src/instagram"
	"my_instagram_follow/src/myexport"
	"my_instagram_follow/src/utils"
	"os"
	"sync"
	"time"
)

func printHelp() {
	fmt.Print("Usage:\n\n")
	fmt.Print("./my_instagram_follow <sessionid_cookie>\n\n")
	fmt.Print("-i <id>: specify id")
	fmt.Print("Files are saved in `record/{todays_date}`")
}

func initDirectory(dirName string) string {
	var path string = fmt.Sprintf("record/%s/%s", time.Now().Format("2006-01-02"), dirName)

	os.MkdirAll(path, os.ModePerm)
	os.MkdirAll(path+"/archive", os.ModePerm)
	return path
}

func fetchFollow(target instagram.Instagram) {
	var wg sync.WaitGroup = sync.WaitGroup{}

	wg.Add(2)
	go func() {
		log.Println("[ ] Fetching followed_by list")
		followers_list := instagram.GetFollowed_by(target, 1000)
		followers_json, err := json.MarshalIndent(followers_list, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		myexport.MyWriteToFileJson(target.ExportPath+"/followers.json", followers_json)
		log.Println("[X] Followers list done!")
		defer wg.Done()

	}()
	go func() {
		log.Println("[ ] Fetching follow list")
		following_list := instagram.GetFollow(target, 210)
		following_json, _ := json.MarshalIndent(following_list, "", "    ")
		myexport.MyWriteToFileJson(target.ExportPath+"/following.json", following_json)
		log.Println("[X] Following list done!")
		defer wg.Done()
	}()
	wg.Wait()
}

func core(cookie string) {
	var target instagram.Instagram = instagram.New(utils.PrepareCookie(cookie), instagram.FakeUserProfile())
	target.ExportPath = initDirectory(target.User.Username)

	// fetchFollow(target)
	instagram.ParseFeed()
}

func main() {

	if (len(os.Args) != 2) || (os.Args[1] == "-h") {
		printHelp()
	} else {
		core(os.Args[1])
	}
}
