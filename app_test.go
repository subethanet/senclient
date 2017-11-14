package senclient

import (
	"testing"
	"fmt"
	"time"
)

func TestAppStart(t *testing.T) {
	app1 := Create(4242)
	fmt.Println(app1)

	time.Sleep(10 * time.Millisecond)

	app2 := Create(4243)
	app2.ConnectAndListen("127.0.0.1", 4242)
}