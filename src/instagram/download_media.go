package instagram

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"my_instagram_follow/src/myexport"
	"my_instagram_follow/src/utils"
	"os"
	"strings"
)

func DownloadUrl(filename string, filepath string, url string, instance Instagram) {
	var segment []string = nil
	var output string = fmt.Sprintf("%s/%s", filepath, filename)

	if filename == "" {
		tmp := strings.Split(strings.TrimRight(url, "/"), "/")
		segment = strings.Split(tmp[len(tmp)-1], "?")
		filename = segment[0]
		output = fmt.Sprintf("%s/%s", filepath, filename)
	}
	req := utils.PrepareRequest(instance.Cookie, url)
	res := utils.DoRequest(req)

	if err := os.WriteFile(output, res, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func ParseFeed() {
	// media type
	// 1 = single image
	// 2 = video
	// 8 = carousel

	type image_versions2 struct {
		Candidates []media
	}
	type carouselMedia struct {
		Id              string
		Media_type      int
		Image_versions2 image_versions2
	}
	type items struct {
		Taken_at         int
		Id               string
		Pk               int
		Device_timestamp int
		Media_type       int
		Image_versions2  image_versions2
		Carousel_media   []carouselMedia
		Video_versions   []media
	}
	type mediaFeed struct {
		Items []items
	}
	var datas mediaFeed

	jsonFile, err := os.Open("model_output.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bytes, &datas)

	toExport, err := json.Marshal(datas)
	if err != nil {
		log.Fatal(err)
	}

	myexport.MyWriteToFileJson("./test.json", toExport)
}

func DownloadMedia(instance Instagram) {

}
