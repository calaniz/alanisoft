package util

import (
	"log/syslog"
	"github.com/drbawb/mustache"
)

var logger, _ = syslog.New(syslog.LOG_NOTICE | syslog.LOG_LOCAL1, "alanisoft-handler")
type MustacheContext map[string]interface{}

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
	return "public/js" + t
}

func L(l string) string {
	return "public/js/layouts" + l
}

func RenderIndex(content string, context ...map[string]interface{}) []byte {
	return []byte(mustache.RenderFileInLayout(T(content), L("/app.html"), NewMustacheContext(context)))
}
