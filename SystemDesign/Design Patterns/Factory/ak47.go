package main

type Ak47 struct {
	Gun
}

// The actual object of Ak47
func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
