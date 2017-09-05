// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

// import "fmt"
// import "time"
import motorMock "./src/quicky/mock/motor"

func main() {

	motorMock := motorMock.Create("Left motor")
    motorMock.DisplayStatus()

  
}
