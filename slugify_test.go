package slugify

import (
	"testing"
)

func TestS(t *testing.T) {
	if S("Hello World") != "hello-world" {
		t.Fail()
	}
	if S("Hello           World????????") != "hello-world" {
		t.Fail()
	}

	if S("what's my name again?") != "whats-my-name-again" {
		t.Fail()
	}

	// all runes-to-be-replaced are contained here
	// to generate:
	//	s := ""
	//	for r, _ := range replacements {
	//		s += string(r)
	//	}
	//	fmt.Println(s)
	if S("äìĩőôüűũãëẽóûâêòúáéùñàèøïçß'æîöíõ") != "aiioouuuaeeouaeouaeunaeoicssaeioio" {
		t.Fail()
	}
}
