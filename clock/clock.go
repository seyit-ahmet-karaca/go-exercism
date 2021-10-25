package clock

import "fmt"

type Clock int

const minutesOfADay = 1440
const minutesOfAHour = 60

// New create new Clock type from given hours and minutes
func New(hour, minute int) Clock {

	minutes := (hour*minutesOfAHour + minute) % minutesOfADay
	// a == (a / b * b) + a % b
	if minutes < 0 {
		minutes = minutesOfADay + minutes
	}
	return Clock(minutes)
}

// Add increase minute to the clock given minute
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

// Subtract decrease minute to the clock given minute
func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

// String returns "HH:mm" date format
func (c Clock) String() string {
	hour := int(c) / 60
	minute := int(c) % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}
