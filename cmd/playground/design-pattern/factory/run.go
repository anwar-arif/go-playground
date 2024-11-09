package factory

import "fmt"

func printDetails(gun Gun) {
	fmt.Println("gun: ", gun.getName())
	fmt.Println("power: ", gun.getPower())
}

func Run() {
	ak47, _ := getGun("ak47")
	maverick, _ := getGun("maverick")
	printDetails(ak47)
	printDetails(maverick)
}
