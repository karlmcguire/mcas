package mcas

import (
	"fmt"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	for i := 0; i < 16; i++ {
		go func() {
			fmt.Println(PID())
		}()
	}
	time.Sleep(time.Second * 2)
}
