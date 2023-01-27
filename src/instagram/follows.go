package instagram

import (
	"encoding/json"
	"fmt"
	"log"
	"my_instagram_follow/src/utils"
	"net/http"
)

func GetFollowed_by(instance Instagram, limit int) map[string]profile {
	var url string = fmt.Sprintf("https://i.instagram.com/api/v1/friendships/%s/followers/?count=0&search_surface=follow_list_page", instance.User.Id)
	var req *http.Request = utils.PrepareRequest(instance.Cookie, url)
	var l_users []profile = nil
	var tmp listUsers

	for len(l_users) < limit {
		utils.DoSleepRange(5, 10)
		if err := json.Unmarshal(utils.DoRequest(req), &tmp); err != nil {
			panic(err)
		} else {
			l_users = append(l_users, tmp.Users...)
		}
		if tmp.Next_max_id != "" {
			url = fmt.Sprintf("https://i.instagram.com/api/v1/friendships/%s/followers/?count=0&max_id=%s&search_surface=follow_list_page", instance.User.Id, tmp.Next_max_id)
			req = utils.PrepareRequest(instance.Cookie, url)
		} else {
			break
		}
		log.Println(len(tmp.Users))
	}

	return ParseUser(l_users)
}

func GetFollow(instance Instagram, limit int) map[string]profile {
	var url string = fmt.Sprintf("https://i.instagram.com/api/v1/friendships/%s/following/?count=0", instance.User.Id)
	var req *http.Request = utils.PrepareRequest(instance.Cookie, url)
	var l_users []profile = nil
	var tmp listUsers

	for len(l_users) < limit {
		if err := json.Unmarshal(utils.DoRequest(req), &tmp); err != nil {
			panic(err)
		} else {
			l_users = append(l_users, tmp.Users...)
		}
		if tmp.Next_max_id != "" {
			url = fmt.Sprintf("https://i.instagram.com/api/v1/friendships/%s/following/?count=0&max_id=%s", instance.User.Id, tmp.Next_max_id)
			req = utils.PrepareRequest(instance.Cookie, url)
			utils.DoSleepRange(5, 10)
		} else {
			break
		}
	}

	return ParseUser(l_users)
}
