package api

import (
	"html/template"

	logsalot "github.com/MhunterDev/hunterdev/src/base/logs"
	"github.com/gin-gonic/gin"
)

func handleHome(c *gin.Context) {
	template, err := template.ParseFiles(".public/welcome.html")
	if err != nil {
		logsalot.ApiErr(err)
	}
	template.Execute(c.Writer, nil)
}

func Router() {
	router := gin.Default()

	router.GET("/", handleHome)

	router.RunTLS(":80", "/usr/lib/mhdev/keychain/tls/CA.crt", "/usr/lib/mhdev/keychain/tls/secret/CA.key")
}
