package utils

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/astaxie/beego/context"
)

func String2md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func GenerateToken(username string) string {
	return String2md5(username)
}

func ValidateToken(ctx *context.Context) bool {
	session := ctx.Input.CruSession.Get("apkuser")
	if session != nil {
		return true
	}
	return false
}
