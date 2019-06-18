package peer

import (

	"crypto/rand"

	"github.com/mr-tron/base58/base58"
)

type Peer struct {
	ID
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