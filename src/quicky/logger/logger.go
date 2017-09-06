
package logger

import "fmt"

type Log struct {
	name string
    
}

func Create(name string, controller chan<- float64, telemetry <-chan float64) *brain {
	instance := brain{name: name, telemetry: telemetry, controller: controller}
	instance.run()
	return &instance
}

func (r *Log) write(topic string, message string, level int) {
    fmt.Printf("%s: %s", topic, message)
}
