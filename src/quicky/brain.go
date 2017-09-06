
package brain

import "fmt"

type brain struct {
	name string
    telemetry <-chan float64
    controller <-chan float64
}

func Create(name string, telemetry <-chan float64, controller <-chan float64) *brain {
	instance := brain{name: name, telemetry: telemetry, controller: controller}
	instance.run()
	return &instance
}

func (r *brain) run() {
    fmt.Printf("\"%s\" has been created\n", r.name)
	go func() {
		for {
			telemetry := <-r.telemetry
			fmt.Printf("\"%s\" telemetry %g\n", r.name, telemetry);
		}
	}()
}

func (r *brain) DisplayStatus() {

	fmt.Printf("\"%s\" is thinking\n", r.name);
}