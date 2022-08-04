package instagram

import (
	"fmt"
	"my_instagram_follow/src/utils"
	"net/http"
)

func GetFollowers(sessionid string) map[string]Profile {
	var url string = fmt.Sprintf("https://i.instagram.com/api/v1/friendships/%s/followers/?count=0", utils.GetUserId(sessionid))
	var response []byte = nil

	req, _ := http.NewRequest("Get", url, nil)
	req.Header.Add("cookie", utils.PrepareCookie(sessionid))
	req.Header.Add("User-Agent", "Instagram 219.0.0.12.117 Android")
	response = utils.DoRequest(req)

	return ParseResponse(response)
}

func GetFollowing(sessionid string) map[string]Profile {
	var url string = fmt.Sprintf("https://i.instagram.com/api/v1/friendships/%s/following/?count=0", utils.GetUserId(sessionid))
	var response []byte = nil

	req, _ := http.NewRequest("Get", url, nil)
	req.Header.Add("cookie", utils.PrepareCookie(sessionid))
	req.Header.Add("User-Agent", "Instagram 219.0.0.12.117 Android")
	response = utils.DoRequest(req)

	return ParseResponse(response)
}
