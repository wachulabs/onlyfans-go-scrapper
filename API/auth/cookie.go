package auth

// "strings"

// Cookie store details of a cookie
type Cookie struct {
	AuthID     string `json:"auth_id"`
	Sess       string `json:"sess"`
	AuthHash   string `json:"auth_hash"`
	AuthUnique string `json:"auth_unique_"`
	AuthUID    string `json:"auth_uid_"`
}

//func ParseCookie(cookie string) Cookie{
