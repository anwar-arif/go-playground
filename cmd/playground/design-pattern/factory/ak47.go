package factory

type ak47 struct {
	gun
}

func newAk47() Gun {
	return &ak47{
		gun{
			name:  "Ak47 gun",
			power: 4,
		},
	}
}
