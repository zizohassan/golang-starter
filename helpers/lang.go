package helpers

import (
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCurrentLang(g *gin.Context) string {
	return gotrans.DetectLanguage(g.GetHeader("Accept-Language"))
}

func GetCurrentLangFromHttp(r *http.Request) string {
	return r.Header.Get("Accept-Language")
}
