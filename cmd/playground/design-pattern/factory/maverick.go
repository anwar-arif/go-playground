package factory

type maverick struct {
	gun
}

func newMaverick() Gun {
	return &ak47{
		gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}
