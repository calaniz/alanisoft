package util

import (
	"log/syslog"
	"os"
	"github.com/drbawb/mustache"
)

var logger, _ = syslog.New(syslog.LOG_NOTICE | syslog.LOG_LOCAL1, "alanisoft-handler")
type MustacheContext map[string]interface{}

func GetConfigKey(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		val = def
	}
	return val
}

func NewMustacheContext(context []map[string]interface{}) map[string]interface{} {
	var ctx map[string]interface{} = nil
	if context != nil {
		ctx = context[0]
	} else {
		ctx = make(map[string]interface{}, 1)
	}

	return ctx
}

func T(t string) string {
	gopath := GetConfigKey("GOPATH", "/gopath")
	if _, err := os.Stat(gopath + "/src/app/public" + t); err == nil {
		return gopath + "/src/app/public" + t
	} else {
		return gopath + "/src/github.com/calaniz/alanisoft/public/js" + t
	}
}

func L(l string) string {
	gopath := GetConfigKey("GOPATH", "/gopath") 
	if _, err := os.Stat(gopath + "/src/app/public" + l); err == nil {
		return gopath + "/src/app/public" + l
	} else {
		return gopath + "/src/github.com/calaniz/alanisoft/public/js/layouts" + l
	}
}

func RenderIndex(content string, context ...map[string]interface{}) []byte {
	return []byte(mustache.RenderFileInLayout(T(content), L("/app.html"), NewMustacheContext(context)))
}
