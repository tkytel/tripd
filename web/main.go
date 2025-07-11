package web

import (
	"embed"
	"log"
	"net/http"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/tkytel/tripd/handler"
)

//go:embed static/*
var staticfs embed.FS

//go:embed views/*
var viewsfs embed.FS

func Init() {
	engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	prometheus := fiberprometheus.New("tripd")
	prometheus.SetIgnoreStatusCodes([]int{401, 403, 404})
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	api := app.Group("/api")
	api.Get("/peers", handler.HandlePeers)
	api.Get("/about", handler.HandleAbout)
	api.Get("/metrics", handler.HandleMetrics)

	app.Get("/static/*", func(c *fiber.Ctx) error {
		filePath := c.Params("*")
		data, err := staticfs.ReadFile("static/" + filePath)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.Type(filePath).Send(data)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		if handler.Ready {
			return c.Render("views/index", fiber.Map{})
		} else {
			return c.SendStatus(503)
		}
	})

	log.Fatal(app.Listen(":3000"))
}
