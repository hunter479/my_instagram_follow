package instagram

func New(cookie string, profile userProfile) Instagram {
	var instance Instagram

	instance.Cookie = cookie
	instance.User = profile

	return instance
}
