package peer

import (
	"testing"
)

func TestNewRandomNodeID(t *testing.T) {
	id1 := NewRandomNodeID()
	id2 := NewRandomNodeID()
	if id1.String() == id2.String() {
		t.Fatal("no two ids should be equal")
	}
}

func TestPeer_StartServer(t *testing.T) {
	p := NewPeer(3000, []byte{127, 0, 0, 1})
	p.StartServer(nil, nil)
}