package message

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	got := Greeting("Sam", "Hi")
	wanted := "Hi, Sam"

	if got != wanted {
		t.Errorf("\ngot %s\nwanted %s", got, wanted)
	}
}
