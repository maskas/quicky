// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

import "fmt"
import "time"
import motorMock "./src/quicky/mock/motor"

func main() {
    fmt.Println("Program has been started");


    leftMotorVoltage := make(chan float64)
	motorMock := motorMock.Create("Left motor", leftMotorVoltage)
    
    motorMock.DisplayStatus()
    leftMotorVoltage <- 5.0
    time.Sleep(time.Second);   
    motorMock.DisplayStatus()
    time.Sleep(time.Second);   
    motorMock.DisplayStatus()
        

  
}
