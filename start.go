
package main

import "fmt"
import "time"
import motorMock "./src/quicky/mock/motor"
import brain "./src/quicky/brain"
import logger ".src/quicky/logger"

func main() {
    fmt.Println("Program has been started");

    leftMotorVoltage := make(chan float64)
    leftMotorRpm := make(chan float64)

    logger := Logger{}

    brain.Create("Motorbrain", leftMotorVoltage, leftMotorRpm)

    motorMock := motorMock.Create("Left motor", leftMotorVoltage, leftMotorRpm)

    
    motorMock.DisplayStatus()
    leftMotorVoltage <- 5.2
    time.Sleep(time.Second*10);   
}
