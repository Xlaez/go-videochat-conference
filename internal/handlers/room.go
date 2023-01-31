package handlers

import (
	"fmt"

	rtc "go-videochat-conference/pkg/webrtc"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}

func RoomWebSocket(c *fiber.Ctx) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}

	_, _, room := createOrGetRoom(uuid)
	rtc.RoomConn(c, room.Peers)

}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}

	uuid, suuid, _ := createOrGetRoom(uuid)
}

func createOrGetRoom(uuid string) (string, string, Room) {

}

func RoomViewerWebSocket(c *websocket.Conn)          {}
func roomViewerConn(c *websocket.Conn, p *rtc.Peers) {}

type websocketMsg struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
