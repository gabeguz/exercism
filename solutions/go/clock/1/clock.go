package clock

import "fmt"

const testVersion = 4

// You can find more details and hints in the test file.

type Clock struct {
	minute int
	hour   int
}

// New creates a new clock with an initial time
func New(hour, minute int) Clock {
	// convert hours to minutes
	h := 0
	m := hour * 60
	// add remaining minutes
	m = m + minute

	// convert minutes back to hours

	// default case for minutes
	if m < 60 && m > -1 {
		m = m
	}

	// sixty minutes is next hour
	if m == 60 {
		m = 0
		h++
	}

	// minutes roll over
	if m > 60 {
		h = h + m/60
		m = m % 60
	}

	// negative minutes roll over
	if m < -60 {
		h = h + m/60
		m = m % 60
	}

	// negative minutes
	if m < 0 && m > -60 {
		h = h - 1
		m = 60 + m
	}

	// negative hours roll over
	if h < -24 {
		h = h % 24
	}

	// negative hours
	if h < 0 && h > -24 {
		h = 24 + h
	}

	// midnight is zero and hours roll over
	if h >= 24 {
		h = h % 24
	}

	return Clock{int(m), int(h)}
}

func (c Clock) String() string {
	s := fmt.Sprintf("%02d:%02d", c.hour, c.minute)
	return s
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return New(c.hour, c.minute)
}