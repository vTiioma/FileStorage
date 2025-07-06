package peer2peer

// represents the remote node
type Peer interface{}

// anything that handles communiction between nodes in the network
// can take the form of: TCP, UDP, WebSockets, etc.
type Transport interface {
	Listen() error
}
