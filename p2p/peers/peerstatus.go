package peers

import (
	pb "github.com/Qitmeer/qitmeer/p2p/proto/v1"
	"github.com/Qitmeer/qitmeer/p2p/qnr"
	"github.com/libp2p/go-libp2p-core/network"
	ma "github.com/multiformats/go-multiaddr"
	"time"
)

// PeerConnectionState is the state of the connection.
type PeerConnectionState int32

const (
	// PeerDisconnected means there is no connection to the peer.
	PeerDisconnected PeerConnectionState = iota
	// PeerDisconnecting means there is an on-going attempt to disconnect from the peer.
	PeerDisconnecting
	// PeerConnected means the peer has an active connection.
	PeerConnected
	// PeerConnecting means there is an on-going attempt to connect to the peer.
	PeerConnecting
)

// peerStatus is the status of an individual peer at the protocol level.
type peerStatus struct {
	address               ma.Multiaddr
	direction             network.Direction
	peerState             PeerConnectionState
	qnr                   *qnr.Record
	metaData              *pb.MetaData
	chainState            string
	chainStateLastUpdated time.Time
	badResponses          int
}
