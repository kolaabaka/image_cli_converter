package configuration

import "fmt"

type configFlags struct {
	RotateHorizontal, RotateVertical, GraySpectre bool
	FileName                                      string
}

func Config(argsList []string) (configuration configFlags) {
	//Optional program args
	for i, v := range argsList[:] {
		if i != 0 {
			switch v {
			case "rh":
				configuration.RotateHorizontal = true
			case "rv":
				configuration.RotateVertical = true
			case "g":
				configuration.GraySpectre = true
			case "p":
				configuration.FileName = argsList[i+1]
			default:
				if argsList[i-1] != "p" {
					fmt.Printf("Args '%v' doesn`t provided\n", v)
				}
			}
		}
	}
	return
}
