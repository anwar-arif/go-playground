package factory

import "fmt"

func getGun(gunType string) (Gun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "maverick" {
		return newMaverick(), nil
	}
	return nil, fmt.Errorf("wrong gun type")
}
