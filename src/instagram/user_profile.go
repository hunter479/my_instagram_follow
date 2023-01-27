package instagram

import (
	"encoding/json"
	"fmt"
	"log"
	"my_instagram_follow/src/utils"
)

func GetUserById(userId string) userProfile {
	type user struct {
		Username string
	}
	type userReelNested struct {
		User user
	}
	type userNested struct {
		Reel userReelNested
	}
	type reelNested struct {
		User userNested
	}
	type reqOne struct {
		Data reelNested
	}
	var url string = fmt.Sprintf("https://.instagram.com/graphql/query/?query_hash=c9100bf9110dd6361671f113dd02e7d6&variables={\"user_id\":%s,\"include_chaining\":false,\"include_reel\":true,\"include_suggested_users\":false,\"include_logged_out_extras\":false,\"include_highlight_reels\":false,\"include_related_profiles\":false}", userId)
	req := utils.PrepareRequest("", url)
	res := utils.DoRequest(req)
	var data reqOne

	if err := json.Unmarshal(res, &data); err != nil {
		log.Fatal(err)
	}
	utils.DoSleepN(2)
	return GetUserByUsername(data.Data.User.Reel.User.Username)
}

func GetUserByUsername(username string) userProfile {
	type FollowCount struct {
		Count int
	}
	type userData struct {
		Username           string
		Id                 string
		Is_private         bool
		Profile_pic_url_hd string
		Edge_followed_by   FollowCount
		Edge_follow        FollowCount
	}
	type reqData struct {
		User userData
	}
	type reqProfile struct {
		Data   reqData
		Status string
	}

	var url string = fmt.Sprintf("https://i.instagram.com/api/v1/users/web_profile_info/?username=%s", username)
	req := utils.PrepareRequest("", url)
	res := utils.DoRequest(req)
	var data reqProfile
	var toReturn userProfile

	if err := json.Unmarshal(res, &data); err != nil {
		log.Fatal(err)
	}
	toReturn.Username = data.Data.User.Username
	toReturn.Id = data.Data.User.Id
	toReturn.Profile_pic_url = data.Data.User.Profile_pic_url_hd
	toReturn.Follow = data.Data.User.Edge_follow.Count
	toReturn.Followed_by = data.Data.User.Edge_followed_by.Count

	return toReturn
}
