package timetool

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeAfterFunc(t *testing.T) {
	cnt := 0
	time.AfterFunc(3*time.Second, func() {
		cnt = 100
		fmt.Println("After 3s ", cnt)
	})

	if cnt == 0 {
		fmt.Println("1111", cnt)
		return
	}

	time.Sleep(5 * time.Second)
}
