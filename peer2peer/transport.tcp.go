package peer2peer

import (
	"fmt"

	"net"
	"sync"
)

type TcpPeer struct {
	// the underlying connection of the peer
	conn net.Conn
	// if we dial out and retrieve a connection -> outbound = true
	// if we accept and retrieve a connection -> outbound = false
	outbound bool
}

type TcpTransport struct {
	listenAddress string
	listener      net.Listener

	mutex sync.RWMutex
	peers map[net.Addr]Peer
}

func CreateTcpPeer(conn net.Conn, outbound bool) *TcpPeer {
	return &TcpPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func CreateTcpTransport(port string) *TcpTransport {
	return &TcpTransport{
		listenAddress: port,
	}
}

func (t *TcpTransport) Listen() error {
	ln, err := net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	t.listener = ln
	go t.createAcceptLoop()

	return nil
}

func (t *TcpTransport) createAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		go t.handleConnection(conn)
	}
}

func (t *TcpTransport) handleConnection(conn net.Conn) {
	fmt.Printf("new incomming connection %+v\n", conn)
}
