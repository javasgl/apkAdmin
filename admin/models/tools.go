package models

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/astaxie/beego/context"
)

//md5
func String2md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

//generate token
func GenerateToken(username string) string {
	return String2md5(username)
}

//validate session
func ValidateToken(ctx *context.Context) bool {
	session := ctx.Input.CruSession.Get(SessionKey)

	if session != nil {
		return true
	}
	return false
}

//get image types
func GetImageType(contentType string) string {

	switch contentType {
	case "image/png":
		return "png"
	case "image/jpeg":
		fallthrough
	case "image/jpg":
		return "jpg"
	case "image/gif":
		return "gif"
	default:
		return ""
	}
}
