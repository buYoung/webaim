package getproc

import (
	"os"
)

var (
	Checkstate =false
	Checkautostate = false
	CheckRedBlue = uint8(255)
)

func Onaim() {
	Checkstate = true
	println("on")
}
func Offaim() {
	Checkstate = false
	println("off")
}

func RedColors() {
	CheckRedBlue = 1
	if Settings.Defalut.Headonly == "1" {
		//Colorlist = Settings.Colors.RedTeamHead[:]
		Colorarrays = Settings.Colors.NRTH
	} else if Settings.Defalut.Bodyonly == "1" {
		Colorarrays = Settings.Colors.NRTB
	} else {
		Colorarrays = Settings.Colors.NRTB
		for _, color := range Settings.Colors.NRTH{
			Colorarrays = append(Colorarrays, color)
		}
	}
}

func BlueColors() {
	CheckRedBlue = 2
	if Settings.Defalut.Headonly == "1" {
		//Colorlist = Settings.Colors.BlueTeamHead[:]
		Colorarrays = Settings.Colors.NBTH
	} else if Settings.Defalut.Bodyonly == "1" {
		Colorarrays = Settings.Colors.NBTB
	} else {
		Colorarrays = Settings.Colors.NBTB
		for _, color := range Settings.Colors.NBTH{
			Colorarrays = append(Colorarrays, color)
		}
	}
}

func MultipleColors() {
	CheckRedBlue = 3
	if Settings.Defalut.Headonly == "1" {
		Colorarrays = Settings.Colors.NBTH
		for _, color := range Settings.Colors.NRTH{
			Colorarrays = append(Colorarrays, color)
		}
	} else if Settings.Defalut.Bodyonly == "1" {
		Colorarrays = Settings.Colors.NBTB
		for _, color := range Settings.Colors.NRTB{
			Colorarrays = append(Colorarrays, color)
		}
	} else {
		Colorarrays = Settings.Colors.NRTB
		for _, color := range Settings.Colors.NBTB{
			Colorarrays = append(Colorarrays, color)
		}
		for _, color := range Settings.Colors.NRTH{
			Colorarrays = append(Colorarrays, color)
		}
		for _, color := range Settings.Colors.NBTH{
			Colorarrays = append(Colorarrays, color)
		}
	}
}
func Exitprogram() {
	os.Exit(0)
}

func Autoon() {
	println("a on")
	Checkautostate = true
}

func Autooff() {
	println("a off")
	Checkautostate = false
}

func Autofire() {
	CheckifApress  = false
}

func Bodyonly() {
	if Settings.Defalut.Bodyonly == "1" {
		Settings.Defalut.Bodyonly = "0"
	} else {
		Settings.Defalut.Bodyonly = "1"
	}
	Settings.Defalut.Headonly = "0"
	switch CheckRedBlue {
	case 1:
		if Settings.Defalut.Bodyonly == "0"{
			RedColors()
			return
		}
		Colorarrays = Settings.Colors.NRTB
	case 2:
		if Settings.Defalut.Bodyonly == "0"{
			BlueColors()
			return
		}
		Colorarrays = Settings.Colors.NBTB
	case 3:
		if Settings.Defalut.Bodyonly == "0"{
			MultipleColors()
			return
		}
		Colorarrays = Settings.Colors.NBTB
		for _, color := range Settings.Colors.NRTB{
			Colorarrays = append(Colorarrays, color)
		}
	}
}

func Headonly() {
	if Settings.Defalut.Headonly == "1" {
		Settings.Defalut.Headonly = "0"
	} else {
		Settings.Defalut.Headonly = "1"
	}
	Settings.Defalut.Bodyonly = "0"
	switch CheckRedBlue {
	case 1:
		if Settings.Defalut.Bodyonly == "0"{
			RedColors()
			return
		}
		Colorarrays = Settings.Colors.NRTH
	case 2:
		if Settings.Defalut.Bodyonly == "0"{
			BlueColors()
			return
		}
		Colorarrays = Settings.Colors.NBTH
	case 3:
		if Settings.Defalut.Bodyonly == "0"{
			MultipleColors()
			return
		}
		Colorarrays = Settings.Colors.NBTH
		for _, color := range Settings.Colors.NRTH{
			Colorarrays = append(Colorarrays, color)
		}
	}
}

func Mousetest(ch chan int) {
	Getkeystates(ch)
}