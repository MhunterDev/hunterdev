package api

import (
	"net/http"

	logsalot "github.com/MhunterDev/hunterdev/src/base/logs"
	db "github.com/MhunterDev/hunterdev/src/db"
	"github.com/gofiber/fiber/v2"
)

func handleHome(c *fiber.Ctx) error {
	return c.Render("/workspaces/hunterdev/src/web/.public/html/welcome.html", fiber.Map{}, "html")
}

func handleAuth(c *fiber.Ctx) error {
	var u db.User
	c.BodyParser(&u)
	if db.AuthUser(u) != nil {
		c.Redirect("/", http.StatusForbidden)
	}
	return c.Render("/workspaces/hunterdev/src/web/.public/html/home.html", fiber.Map{}, "html")
}

func Router() {

	router := fiber.New()
	router.Get("/", handleHome)

	err := router.ListenTLS(":80", "/usr/lib/mhdev/keychain/tls/CA.crt", "/usr/lib/mhdev/keychain/tls/secret/CA.key")
	if err != nil {
		logsalot.ApiErr(err)
		return
	}
}
