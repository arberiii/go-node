package peer

import (

	"crypto/rand"
	"fmt"
	"log"
	"net"

	"github.com/mr-tron/base58/base58"
)

type Peer struct {
	ID
	Addr []byte
	Port int
}

const IDLength = 20

type ID []byte

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

func (p *Peer) StartServer(handle func([]byte), reply func(*net.UDPAddr)) {
	addr := &net.UDPAddr{IP: p.Addr, Port: p.Port, Zone: ""}
	ServerConn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Panic(err)
	}
	// handle the error
	defer ServerConn.Close()
	buffer := make([]byte, 1024)

	for {
		_, remoteAddr, err := ServerConn.ReadFromUDP(buffer)
		fmt.Println(err)
		// handle the message
		handle(buffer)
		if reply != nil {
			reply(remoteAddr)
		}
	}
}

func (p *Peer) Send(b []byte, raddr *net.UDPAddr) error {
	ClientConn, err := net.DialUDP("udp", &net.UDPAddr{IP: p.Addr, Port: p.Port, Zone: ""}, raddr)
	if err != nil {
		return err
	}
	defer ClientConn.Close()

	ClientConn.Write(b)

	return nil
}