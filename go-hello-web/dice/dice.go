// Package dice exists purely for demonstrative purposes.
// It just provides a mock SDK that can be imported.
// In reality one might use an actual AWS Go SDK service, for example.
// Pretend that in reality this object calls some API service with each function.
package dice

import "math/rand"

// New returns a dice rolling service.
func New() *Dice {
	return &Dice{}
}

// Dice is just an example service SDK object that returns dice rolls.
// It can handle multiple die of various sizes.
type Dice struct{}

// Roll6 returns a random number between 1-6 (inclusive)
func (d *Dice) Roll6() (int, error) {
	return rand.Intn(6) + 1, nil
}

// Roll20 returns a random number between 1-20 (inclusive)
func (d *Dice) Roll20() (int, error) {
	return rand.Intn(20) + 1, nil
}
