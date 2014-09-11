package handlers

import (
	"net/http"
	"github.com/calaniz/alanisoft/handlers/util"
)

func GetMain(w http.ResponseWriter, r *http.Request) *AppError {
	w.WriteHeader(http.StatusOK)
	w.Write(util.RenderIndex("/main/main.html"))
	return nil
}