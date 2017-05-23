/*Package say is an example go package with no real functionality

It covers the basics of:
 - interfaces
 - type syntax
*/
package say

import (
	// "fmt" //unused imports are compile failures
	"testing"
)

func TestUsay_Say(t *testing.T) {
	//A map (of keys of type usay, of values string) is a hash Table
	tests := map[usay]string{
		0:  "Zero!!!!",
		1:  "Odd!",
		2:  "Even!",
		3:  "Odd!",
		4:  "Even!",
		5:  "Odd!",
		6:  "Even!",
		7:  "Odd!",
		8:  "Even!",
		9:  "Odd!",
		10: "Even!",
		11: "Too Big!!",
	}

	//this rendition of for, with the range keywork, which iterates through
	// all the keys in the above map in random order
	for key, value := range tests {
		if got := key.Say(); got != value {
			t.Errorf("Say() on %d failed:  Got %q, not %q as expected", key, got, value)
		}
	}
}

func TestI_Say(t *testing.T) {
	//A map (of keys of type usay, of values string) is a hash Table
	tests := map[I]string{
		0:  "0: Zero!!!!",
		1:  "1: Odd!",
		-1: "1: Odd!",
		2:  "2: Even!",
	}

	//this rendition of for, with the range keywork, which iterates through
	// all the keys in the above map in random order
	for key, value := range tests {
		if got := key.Say(); got != value {
			t.Errorf("Say() on %d failed:  Got %q, not %q as expected", key, got, value)
		}
	}
}
