package handlers

import (
	"net/http"
	"github.com/calaniz/alanisoft/handlers/util"
)

func GetBlog(w http.ResponseWriter, r *http.Request) *AppError {
	w.WriteHeader(http.StatusOK)
	w.Write(util.RenderIndex("/blog/blog.html"))
	return nil
}
