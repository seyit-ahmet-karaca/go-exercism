package elon

import "fmt"

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// TODO: define the 'Drive()' method
func (c *Car) Drive() {

	c.battery = c.battery - c.batteryDrain
	c.distance = c.distance + c.speed
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c Car) CanFinish(trackDistance int) bool {
	for c.distance < trackDistance && c.battery-c.batteryDrain >= 0 {
		c.Drive()
	}

	return c.distance >= trackDistance
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}
