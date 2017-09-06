

package motor

import "fmt"
import "time"


type motorMock struct {
	name string
    voltage <-chan float64
    rpm float64
    rpmOutput chan<- float64
}

func Create(name string, voltage <-chan float64, rpmOutput chan<- float64) *motorMock {
	instance := motorMock{name: name, voltage: voltage, rpm: 0, rpmOutput: rpmOutput}
	instance.run()
	return &instance
}

func (r *motorMock) run() {
    fmt.Printf("\"%s\" has been created\n", r.name)
	go func() { 
		for {
			voltage := <-r.voltage
			r.rpm = voltage * 1.5
			fmt.Println(r.rpm);
		}
	}()
	go func() { 
		for {
			r.rpmOutput <- r.rpm
			time.Sleep(time.Second);
		}
	}()
}

func (r *motorMock) DisplayStatus() {

	fmt.Printf("\"%s\" is spinning at %g RPM\n", r.name, r.rpm);
	// fmt.Println(r.name + ". Spinning at " + strconv.FormatFloat(r.rpm, 'E', -1, 32))

}