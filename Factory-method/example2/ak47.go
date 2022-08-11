package main

// The type Ak47 is a struct with a field of type Gun.
// @property {Gun}  - Gun is a struct
type Ak47 struct {
	Gun
}

// NewAk47 returns a pointer to an Ak47 struct that has a Gun struct embedded in it.
func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name: "Ak47 gun",
			power: 4,
		},
	}
}