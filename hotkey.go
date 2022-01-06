package main

import (
	"fmt"
	"github.com/MakeNowJust/hotkey"
	getproc "github.com/buYoung/webaim/proccess"
	"strconv"
	"strings"
)

func Sethotkey(index, key, keystr string) error {
	var err error
	if strings.Contains(key, "3/") {
		keys := strings.Split(key, "/")
		switch index {
		case "8":
			switch keys[1] {
			case "left":
				getproc.Checkkeys = 1
			case "right":
				getproc.Checkkeys = 2
			case "wheel":
				getproc.Checkkeys = 3

			}
			ch <- 0
			//hotkeyz.Mouseevents(a, getmousepress,ch, func() {
			//	getproc.Checkifpress = true
			//})
			getproc.Settings.Hotkeys.Work = fmt.Sprintf("%s|%s", key, keystr)
		case "9":
			switch keys[1] {
			case "left":
				getproc.Checkkeys = 1
			case "right":
				getproc.Checkkeys = 2
			case "wheel":
				getproc.Checkkeys = 3

			}
			ch <- 0
			getproc.Settings.Hotkeys.Autowork = fmt.Sprintf("%s|%s", key, keystr)
		case "10":
			switch keys[1] {
			case "left":
				getproc.Checkkeys = 1
			case "right":
				getproc.Checkkeys = 2
			case "wheel":
				getproc.Checkkeys = 3

			}
			ch <- 0

			getproc.Settings.Hotkeys.Autofire = fmt.Sprintf("%s|%s", key, keystr)
		}

	} else if strings.Contains(key, "/") {
		keys := strings.Split(key, "/")
		modif, _ := strconv.Atoi(keys[0])
		numkey, _ := strconv.Atoi(keys[1])

		switch index {
		case "0":
			if int32(onid) != -1 {
				onkey.Unregister(onid)
				onid = -1
			}
			onid, err = onkey.Register(hotkey.Modifier(uint32(modif)), uint32(numkey), func() {
				getproc.Onaim()
			})
			getproc.Settings.Hotkeys.On = fmt.Sprintf("%s|%s", key, keystr)

		case "1":
			if int32(offid) != -1 {
				offkey.Unregister(offid)
				offid = -1
			}
			offid, err = offkey.Register(hotkey.Modifier(uint32(modif)), uint32(numkey), func() {
				getproc.Offaim()
			})
			getproc.Settings.Hotkeys.Off = fmt.Sprintf("%s|%s", key, keystr)
		case "2":
			if int32(redid) != -1 {
				redkey.Unregister(redid)
				redid = -1
			}
			redid, err = redkey.Register(hotkey.Modifier(uint32(modif)), uint32(numkey), func() {
				getproc.RedColors()
			})
			getproc.Settings.Hotkeys.Red = fmt.Sprintf("%s|%s", key, keystr)
		case "3":
			if int32(blueid) != -1 {
				bluekey.Unregister(blueid)
				blueid = -1
			}
			blueid, err = bluekey.Register(hotkey.Modifier(uint32(modif)), uint32(numkey), func() {
				getproc.BlueColors()
			})
			getproc.Settings.Hotkeys.Blue = fmt.Sprintf("%s|%s", key, keystr)
		case "4":
			if int32(solosid) != -1 {
				solokey.Unregister(solosid)
				solosid = -1
			}
			solosid, err = solokey.Register(hotkey.Modifier(uint32(modif)), uint32(numkey), func() {
				getproc.MultipleColors()
			})
			getproc.Settings.Hotkeys.Solo = fmt.Sprintf("%s|%s", key, keystr)
		case "5":
			if int32(exitid) != -1 {
				exitkey.Unregister(exitid)
				exitid = -1
			}
			exitid, err = exitkey.Register(hotkey.Modifier(uint32(modif)), uint32(numkey), func() {
				getproc.Exitprogram()
			})
			getproc.Settings.Hotkeys.Exit = fmt.Sprintf("%s|%s", key, keystr)
		case "6":
			if int32(auonid) != -1 {
				auonkey.Unregister(auonid)
				auonid = -1
			}
			auonid, err = auonkey.Register(hotkey.Modifier(uint32(modif)), uint32(numkey), func() {
				getproc.Autoon()
			})
			getproc.Settings.Hotkeys.Autoon = fmt.Sprintf("%s|%s", key, keystr)
		case "7":
			if int32(auoffid) != -1 {
				auoffkey.Unregister(auoffid)
				auoffid = -1
			}
			auoffid, err = auoffkey.Register(hotkey.Modifier(uint32(modif)), uint32(numkey), func() {
				getproc.Autooff()
			})
			getproc.Settings.Hotkeys.Autooff = fmt.Sprintf("%s|%s", key, keystr)

		case "10":

			if int32(aufireid) != -1 {
				aufirekey.Unregister(aufireid)
				aufireid = -1
			}
			aufireid, err = aufirekey.Register(hotkey.Modifier(uint32(modif)), uint32(numkey), func() {
				getproc.Autofire()
			})

			getproc.Settings.Hotkeys.Autofire = fmt.Sprintf("%s|%s", key, keystr)
		case "11":
			if int32(bodyid) != -1 {
				bodykey.Unregister(bodyid)
				bodyid = -1
			}
			bodyid, err = bodykey.Register(hotkey.Modifier(uint32(modif)), uint32(numkey), func() {
				getproc.Bodyonly()
			})

			getproc.Settings.Hotkeys.Bodyonly = fmt.Sprintf("%s|%s", key, keystr)
		case "12":
			if int32(headid) != -1 {
				headkey.Unregister(headid)
				headid = -1
			}
			headid, err = headkey.Register(hotkey.Modifier(uint32(modif)), uint32(numkey), func() {
				getproc.Headonly()
			})

			getproc.Settings.Hotkeys.Headonly = fmt.Sprintf("%s|%s", key, keystr)
		}
	} else {
		//var resultkey uint32 =0
		//	hotkey.F1
		numkey, _ := strconv.Atoi(key)
		switch index {
		case "0":
			if int32(onid) != -1 {
				onkey.Unregister(onid)
				onid = -1
			}
			onid, err = onkey.Register(hotkey.Modifier(hotkey.None), uint32(numkey), func() {
				getproc.Onaim()
			})
			getproc.Settings.Hotkeys.On = fmt.Sprintf("%s|%s", key, keystr)
		case "1":
			if int32(offid) != -1 {
				offkey.Unregister(offid)
				offid = -1
			}
			offid, err = offkey.Register(hotkey.Modifier(hotkey.None), uint32(numkey), func() {
				getproc.Offaim()
			})
			getproc.Settings.Hotkeys.Off = fmt.Sprintf("%s|%s", key, keystr)
		case "2":
			if int32(redid) != -1 {
				redkey.Unregister(redid)
				redid = -1
			}
			redid, err = redkey.Register(hotkey.Modifier(hotkey.None), uint32(numkey), func() {
				getproc.RedColors()
			})
			getproc.Settings.Hotkeys.Red = fmt.Sprintf("%s|%s", key, keystr)
		case "3":
			if int32(blueid) != -1 {
				bluekey.Unregister(blueid)
				blueid = -1
			}
			blueid, err = bluekey.Register(hotkey.Modifier(hotkey.None), uint32(numkey), func() {
				getproc.BlueColors()
			})
			getproc.Settings.Hotkeys.Blue = fmt.Sprintf("%s|%s", key, keystr)
		case "4":
			if int32(solosid) != -1 {
				solokey.Unregister(solosid)
				solosid = -1
			}
			solosid, err = solokey.Register(hotkey.Modifier(hotkey.None), uint32(numkey), func() {
				getproc.MultipleColors()
			})
			getproc.Settings.Hotkeys.Solo = fmt.Sprintf("%s|%s", key, keystr)
		case "5":
			if int32(exitid) != -1 {
				exitkey.Unregister(exitid)
				exitid = -1
			}
			exitid, err = exitkey.Register(hotkey.Modifier(hotkey.None), uint32(numkey), func() {
				getproc.Exitprogram()
			})
			getproc.Settings.Hotkeys.Exit = fmt.Sprintf("%s|%s", key, keystr)
		case "6":
			if int32(auonid) != -1 {
				auonkey.Unregister(auonid)
				auonid = -1
			}
			auonid, err = auonkey.Register(hotkey.Modifier(hotkey.None), uint32(numkey), func() {
				getproc.Autoon()
			})
			getproc.Settings.Hotkeys.Autoon = fmt.Sprintf("%s|%s", key, keystr)
		case "7":
			if int32(auoffid) != -1 {
				auoffkey.Unregister(auoffid)
				auoffid = -1
			}
			auoffid, err = auoffkey.Register(hotkey.Modifier(hotkey.None), uint32(numkey), func() {
				getproc.Autooff()
			})
			getproc.Settings.Hotkeys.Autooff = fmt.Sprintf("%s|%s", key, keystr)

		case "10":
			if int32(aufireid) != -1 {
				aufirekey.Unregister(aufireid)
				aufireid = -1
			}
			aufireid, err = aufirekey.Register(hotkey.Modifier(hotkey.None), uint32(numkey), func() {
				getproc.Autofire()
			})
			getproc.Settings.Hotkeys.Autofire = fmt.Sprintf("%s|%s", key, keystr)
		case "11":
			if int32(bodyid) != -1 {
				bodykey.Unregister(bodyid)
				bodyid = -1
			}
			bodyid, err = bodykey.Register(hotkey.Modifier(hotkey.None), uint32(numkey), func() {
				getproc.Bodyonly()
			})

			getproc.Settings.Hotkeys.Bodyonly = fmt.Sprintf("%s|%s", key, keystr)
		case "12":
			if int32(headid) != -1 {
				headkey.Unregister(headid)
				headid = -1
			}
			headid, err = headkey.Register(hotkey.Modifier(hotkey.None), uint32(numkey), func() {
				getproc.Headonly()
			})

			getproc.Settings.Hotkeys.Headonly = fmt.Sprintf("%s|%s", key, keystr)
		}
	}
	return err
}
func unhoteky(index string) {
	switch index {
	case "0":
		if int32(onid) != -1 {
			onkey.Unregister(onid)
			onid = -1
		}
		getproc.Settings.Hotkeys.On = ""
	case "1":
		if int32(offid) != -1 {
			offkey.Unregister(offid)
			offid = -1
		}
		getproc.Settings.Hotkeys.Off = ""
	case "2":
		if int32(redid) != -1 {
			redkey.Unregister(redid)
			redid = -1
		}
		getproc.Settings.Hotkeys.Red = ""
	case "3":
		if int32(blueid) != -1 {
			bluekey.Unregister(blueid)
			blueid = -1
		}
		getproc.Settings.Hotkeys.Blue = ""
	case "4":
		if int32(solosid) != -1 {
			solokey.Unregister(solosid)
			solosid = -1
		}
		getproc.Settings.Hotkeys.Solo = ""
	case "5":
		if int32(exitid) != -1 {
			exitkey.Unregister(exitid)
			exitid = -1
		}
		getproc.Settings.Hotkeys.Exit = ""
	case "6":
		if int32(auonid) != -1 {
			auonkey.Unregister(auonid)
			auonid = -1
		}
		getproc.Settings.Hotkeys.Autoon = ""
	case "7":
		if int32(auoffid) != -1 {
			auoffkey.Unregister(auoffid)
			auoffid = -1
		}
		getproc.Settings.Hotkeys.Autooff = ""
	case "8":
		ch1 <- 1
		getproc.Settings.Hotkeys.Work = ""
	case "9":
		ch1 <- 1

		getproc.Settings.Hotkeys.Autowork = ""
	case "10":
		ch1 <- 1

		if int32(aufireid) != -1 {
			aufirekey.Unregister(aufireid)
			aufireid = -1
		}
		getproc.Settings.Hotkeys.Autofire = ""
	case "11":
		ch1 <- 1
		if int32(bodyid) != -1 {
			bodykey.Unregister(bodyid)
			bodyid = -1
		}
		getproc.Settings.Hotkeys.Bodyonly = ""
	case "12":
		ch1 <- 1
		if int32(headid) != -1 {
			headkey.Unregister(headid)
			headid = -1
		}
		getproc.Settings.Hotkeys.Headonly = ""
	}
}
