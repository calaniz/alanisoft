package main

import (
	"log/syslog"
	"net/http"
	"github.com/calaniz/alanisoft/handlers"
	"github.com/calaniz/alanisoft/handlers/util"
	"github.com/gorilla/mux"
)

import logging "github.com/gorilla/handlers"

func NewRouter(logger *syslog.Writer) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{name:[a-zA-Z0-9@/\\-_\\.]+.(js|css|html|json|png|svg|tff|woff|eot)$}", func(w http.ResponseWriter, r *http.Request) {
		logger.Info(util.GetConfigKey("GOPATH", "/gopath") + "/src/github.com/calaniz/alanisoft/public"+r.URL.Path)
		http.ServeFile(w, r, util.GetConfigKey("GOPATH", "/gopath") + "/src/github.com/calaniz/alanisoft/public"+r.URL.Path)
	})

	return r
}

func main() {
	if logger, err := syslog.New(syslog.LOG_NOTICE|syslog.LOG_LOCAL1, "alanisoft-routing"); err != nil {
		logger.Err(err.Error())
	} else {
		r := NewRouter(logger)
		r.Handle("/", handlers.AppHandler(handlers.GetMain)).Methods("GET")
		http.Handle("/", logging.CombinedLoggingHandler(logger, r))
		go func() {
			if err := http.ListenAndServeTLS(":49155", util.GetConfigKey("GOPATH", "/gopath") + "/src/github.com/calaniz/alanisoft/ssl/alanisoft.com.cert",
				util.GetConfigKey("GOPATH", "/gopath") + "/src/github.com/calaniz/alanisoft/ssl/alanisoft.com.key", nil); err != nil {
				logger.Crit(err.Error())
			}
		}()
		if err := http.ListenAndServe(":49154", nil); err != nil {
			logger.Crit(err.Error())
		}
	}
}