package instagram

import (
	"encoding/json"
	"strconv"
)

type Profile struct {
	Pk       int
	Username string
}

type List_Users struct {
	Users []Profile
}

func ParseResponse(obj []byte) map[string]Profile {
	var result List_Users
	var err error = json.Unmarshal(obj, &result)
	var parse map[string]Profile = make(map[string]Profile)

	if err != nil {
		panic(err)
	}
	for i := range result.Users {
		parse[strconv.Itoa(result.Users[i].Pk)] = result.Users[i]
	}

	return parse
}
