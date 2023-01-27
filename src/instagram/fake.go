package instagram

import (
	"log"
	"my_instagram_follow/src/utils"
	"net/http"
)

func FakeUserProfile() userProfile {
	return userProfile{
		Username:        "foggyday_l",
		Id:              "3207377333",
		Profile_pic_url: "URL",
		Follow:          69,
		Followed_by:     69,
	}
}

func tempoInstagram() {
	var client = &http.Client{
		Timeout: 1000000000 * 10,
	}

	req, err := http.NewRequest("Get", "https://i.instagram.com/", nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	utils.DoSleepN(10)
	defer res.Body.Close()
}
