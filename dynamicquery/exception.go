package dynamicquery

import (
	"time"
)

// Exception ...
type Exception struct {
	When time.Time
	What string
}

// NewException ...
func NewException(message string) Exception {
	err := Exception{}
	err.What = message
	err.When = time.Now()
	return err
}

func (err Exception) Error() string {
	return err.What
}
