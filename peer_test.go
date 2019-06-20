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
