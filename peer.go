package peer

import (
	"crypto/rand"
	"log"
	"net"

	"github.com/mr-tron/base58/base58"
)

// Peer is composed of three parts: ID which is a uniquely identifier of a peer,
// Addr is the UDP address, Port is the port which the peer is listening
type Peer struct {
	ID
	Addr []byte
	Port int
}

const IDLength = 20

type ID []byte

// NewRandomNodeID generates a random ID
func NewRandomNodeID() ID {
	buffer := make([]byte, IDLength)
	_, err := rand.Read(buffer)
	if err != nil {
		panic(err)
	}
	var id []byte
	for _, b := range buffer {
		id = append(id, b)
	}
	return id
}

func (id ID) String() string {
	return base58.Encode([]byte(id))
}

func NewPeer(port int, addr []byte) *Peer {
	p := Peer{}
	p.ID = NewRandomNodeID()
	p.Port = port
	p.Addr = addr

	return &p
}

// StartServer listen and serves UDP request directed to the peer. It takes two functions as parameters
// handle(which handles the new requests) and periodicTask(its a place holder for the callee to use for periodic task)
func (p *Peer) StartServer(handle func([]byte, *net.UDPConn,*net.UDPAddr) error, periodicTask func(conn *net.UDPConn)) {
	addr := &net.UDPAddr{IP: p.Addr, Port: p.Port, Zone: ""}
	ServerConn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Panic(err)
	}
	// handle the error
	defer ServerConn.Close()
	buffer := make([]byte, 1024)

	periodicTask(ServerConn)

	for {
		_, remoteAddr, err := ServerConn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		// handle the message
		err = handle(buffer,ServerConn, remoteAddr)
		if err != nil {
			log.Println(err)
		}
	}
}
