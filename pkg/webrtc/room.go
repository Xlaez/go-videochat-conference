package webrtc

import (
	"go-videochat-conference/pkg/chat"
	"log"
	"sync"

	"github.com/fasthttp/websocket"
	"github.com/pion/webrtc/v3"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

func RoomConn(c *websocket.Conn, p *Peers) {
	var config webrtc.Configuration
	peerConnection, err := webrtc.NewPeerConnection(config)

	if err != nil {
		log.Print(err)
		return
	}

	newPeer := PeerConnectionState{
		PeerConnection: peerConnection,
		WebSocket:      &ThreadSafeWriter{},
		Conn:           c,
		Mutex:          sync.Mutex{},
	}
}
