package thread

import (
	fmt "fmt"
	"testing"
	"time"
)

func testcase1(stop chan struct{}) {
	testcase2(stop)
}
func testcase2(stop chan struct{}) {
	for true {
		select {
		case <-stop:
			fmt.Println("success exit")
			return
		default:
			fmt.Println("sleeping......")
			time.Sleep(time.Second * 2)
		}
	}
}

func TestThreadExit(t *testing.T) {
	th := NewThread(testcase1)
	th.Run()

	th2 := NewThread(testcase2)
	th2.Run()

	tt := time.NewTicker(5 * time.Second)
	for true {
		select {
		case <-tt.C:
			th.Stop()
			time.Sleep(1 * time.Second)

			th2.Stop()
			time.Sleep(3 * time.Second)
			return
		}
	}
}
