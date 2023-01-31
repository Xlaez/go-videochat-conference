package webrtc

import (
	"sync"

	"github.com/pion/webrtc/v3"
)

type Peers struct {
	ListLock     sync.RWMutex
	Connesctions []PeerConnectionState
	TrackLocals  map[string]*webrtc.TrackLocalStaticRTP
}

func (p *Peers) DispatchKeyFrame() {

}
