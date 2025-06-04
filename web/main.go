package web

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/tkytel/tripd/handler"
)

//go:embed views/*
var viewsfs embed.FS

func Init() {
	engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	api := app.Group("/api")
	api.Get("/peers", handler.HandlePeers)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index", fiber.Map{})
	})

	log.Fatal(app.Listen(":3000"))
}
