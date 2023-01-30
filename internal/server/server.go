package server

import (
	"flag"
	"os"
	"time"

	"go-videochat-conference/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
)

var (
	addr = flag.String("addr : ", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8081"
	}

	engine := html.New("../../views/", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.RoomWebSocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	app.Get("/room/:uuid/chat", handlers.Room)
	app.Get("/room/:uuid/chat/websocket", handlers.Room)
	app.Get("/room/:uuid/viewer/websocket", handlers.Room)
	app.Get("/stream/:suuid", handlers.Stream)
	app.Get("/stream/:suuid/websocket", handlers.Room)
	app.Get("/stream/:suuid/chat/websocket")
	app.Get("/stream/:suuid/viewer/websocket")
}
