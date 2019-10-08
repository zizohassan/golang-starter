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

func T(g *gin.Context , key ...string) string  {
	s := ""
	for _ , k := range key{
		s += gotrans.Tr(GetCurrentLang(g), k)
	}
	return s
}
