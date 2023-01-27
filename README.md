# my_instagram_follow
A simple learning project done in golang to scrap instagram profile.
Because of Instagram's limit, follows and followed_by lists are limited by default.

## Requirements
Golang is required | https://go.dev/doc/install

You will also need your Instagram sessionId cookie.

### How to get your sessionid cookie
```
On Instagram's website:
[Firefox]
DevTool(f12) -> Storage -> Cookies -> "www.instagram.com" -> copy the value column associated with the "sessionid" column.

[Google Chrome]
DevTool(f12) -> Application -> Storage -> Cookies -> "www.instagram.com" -> copy the value column associated with the "sessionid" column.
```

## Build
```
git clone https://github.com/hunter479/my_instagram_follow.git
cd my_instagram_follow
go build .
```

## Usage
```
./my_instagram_follow <sessionid_cookie>
```
The lists will be saved in the directory "record/{username}/{todays_date}"

## Legal Disclaimer
The use is the responsibility of the end user. Developers assume no liability and are not responsible for any misuse or damage caused.

## To do
- Creation of a function to compare the differences in time between the records
- Add config file to not pass the cookie in CLI
- Use config file or ENV variable to chose where the record are saved.
- Flag to choose the record directory
- Prepare releases so as not to be forced to build the project