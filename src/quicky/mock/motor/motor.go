

package motor

import "fmt"

type motorMock struct {
	name string
    pwm float32
    rpm float32
}

func Create(name string) motorMock {
	instance := motorMock{name: name, pwm: 0, rpm: 0}
	instance.run()
	return instance
}

func (r *motorMock) run() {
    
	fmt.Printf("\"%s\" has been created\n", r.name)
}

func (r *motorMock) DisplayStatus() {

	fmt.Printf("\"%s\" is spinning at %g RPM\n", r.name, r.rpm);
	// fmt.Println(r.name + ". Spinning at " + strconv.FormatFloat(r.rpm, 'E', -1, 32))

}