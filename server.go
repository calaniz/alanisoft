package main

import (
	"log/syslog"
	"net/http"
	"github.com/calaniz/alanisoft/handlers"
	"github.com/gorilla/mux"
)

import logging "github.com/gorilla/handlers"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{name:[a-zA-Z0-9@/\\-_\\.]+.(js|css|html|json|png|svg|tff|woff|eot)$}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/"+r.URL.Path)
	})

	return r
}

func main() {
	if logger, err := syslog.New(syslog.LOG_NOTICE|syslog.LOG_LOCAL1, "alanisoft-routing"); err != nil {
		logger.Err(err.Error())
	} else {
		r := NewRouter()
		r.Handle("/", handlers.AppHandler(handlers.GetMain)).Methods("GET")
		r.Handle("/blog", handlers.AppHandler(handlers.GetBlog)).Methods("GET")
		http.Handle("/", logging.CombinedLoggingHandler(logger, r))

		go func() {
			// Uncomment after we buy a TLS Cert
			// if err := http.ListenAndServeTLS(":10443", "./cert/cert.crt", "./cert/key.pem", nil); err != nil {
			// 	logger.Crit(err.Error())
			// }
		}()
		if err := http.ListenAndServe(":9001", nil); err != nil {
			logger.Crit(err.Error())
		}
	}
}