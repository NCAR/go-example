package busy

import (
	"testing"
)

func TestBasic(t *testing.T) {
	//Basic(1, 2)   // will deadlock: not enough writers
	// Basic(2, 1)   // will deadlock: not enough readers
	Basic(5000, 5000) // works
}

func TestMapRace(t *testing.T) {
	MapRace(5) //wont work reliably due to race condition
}
