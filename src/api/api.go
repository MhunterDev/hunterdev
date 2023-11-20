package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHome(c *gin.Context) {
	c.JSON(http.StatusOK, "Coming Soon")
}

func Router() {
	router := gin.Default()

	router.GET("/", handleHome)

	router.RunTLS(":80", "/usr/lib/mhdev/keychain/tls/CA.crt", "/usr/lib/mhdev/keychain/secret.pem")
}
