package instagram

import (
	"log"
	"strings"
)

func ParseUser(result []profile) map[string]profile {
	var parse map[string]profile = make(map[string]profile)

	for i := range result {
		parse[result[i].Pk_id] = result[i]
	}

	return parse
}

func ExtractUserIdFromCookie(cookie string) string {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal("Could not parse cookie")
		}
	}()

	return strings.Split(cookie, "%")[0]
}
