package golimiter

import (
	"testing"
	"time"
)

func TestFoo(t *testing.T) {
	lim := NewLimiter(5, 2)
	lim2 := NewLimiter(5, 2)

	lim.Allow("user")
	lim2.Allow("user")
	if lim.Allow("user") != true {
		t.Error("Should be true")
	}

	if lim2.Allow("user") != true {
		t.Error("Second instance of limit should not be limited at this point")
	}

	if lim.Allow("user") != false {
		t.Error("Limit is not working")
	}

	if lim.Allow("user") != false {
		t.Error("Limit is not working")
	}

	time.Sleep(1 * time.Second)

	if lim.Allow("user") != true {
		t.Error("Rate should be released")
	}
}
