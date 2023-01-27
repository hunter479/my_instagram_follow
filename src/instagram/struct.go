package instagram

type profile struct {
	Pk_id    string
	Username string
}

type media struct {
	Url    string
	Width  int
	Height int
}

type userProfile struct {
	Username        string
	Id              string
	Profile_pic_url string
	Private         bool
	Followed_by     int
	Follow          int
}

type listUsers struct {
	Users       []profile
	Next_max_id string
}

type Instagram struct {
	Cookie     string
	ExportPath string
	User       userProfile
}
