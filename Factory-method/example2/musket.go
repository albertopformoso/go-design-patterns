package main

// A musket is a gun that has a barrel and a trigger.
// @property {Gun}  - Gun is a struct type
type musket struct {
	Gun
}

// `newMusket()` is a function that returns a pointer to a `musket` struct that implements the `IGun`
// interface
func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
