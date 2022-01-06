package main

import (
	"encoding/json"
	"github.com/MakeNowJust/hotkey"
	_ "github.com/buYoung/webaim/proccess"
	getproc "github.com/buYoung/webaim/proccess"
	"github.com/lxn/win"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

var (
	login                    = 0
	weapon                   = "0"
	chracter                 = "0"
	loginid                  = ""
	skinstate                = false
	version                  = 20
	onkey                    = hotkey.New()
	offkey                   = hotkey.New()
	redkey                   = hotkey.New()
	bluekey                  = hotkey.New()
	solokey                  = hotkey.New()
	exitkey                  = hotkey.New()
	auonkey                  = hotkey.New()
	auoffkey                 = hotkey.New()
	workkey                  = hotkey.New()
	lockkey                  = hotkey.New()
	auworkkey                = hotkey.New()
	aufirekey                = hotkey.New()
	bodykey                  = hotkey.New()
	headkey                  = hotkey.New()
	onid           hotkey.Id = -1
	offid          hotkey.Id = -1
	redid          hotkey.Id = -1
	blueid         hotkey.Id = -1
	solosid        hotkey.Id = -1
	exitid         hotkey.Id = -1
	auonid         hotkey.Id = -1
	auoffid        hotkey.Id = -1
	workid         hotkey.Id = -1
	lockid         hotkey.Id = -1
	auworkid       hotkey.Id = -1
	aufireid       hotkey.Id = -1
	bodyid         hotkey.Id = -1
	headid         hotkey.Id = -1
	numb                     = make(chan string)
	numb1                    = make(chan string)
	_getstatus     getproc.Procinfo
	_GetProc       = make(chan getproc.Procinfo, 1)
	curpos         win.POINT
	ch             = make(chan int, 1)
	ch1            = make(chan int, 1)
	getmousepress  = make(chan bool, 0)
	gettab         = make(chan bool, 0)
	autoclick = make(chan bool, 0)
	autoclickstate = false
	Checkifpress   = false
	CheckifApress  = false
	err            error
	ispress        = false
	chromepath string

)



func main() {
	runtime.GOMAXPROCS(0)
	getChromePath()
	firestStart()
	serverInitialize()
	b := getuserdata()
	if !b {
		return
	}
	print("시작\n")
	go getproc.Gettabkey(ch, ch1)
	go func() {
		getproc.FindActiveProcess(_GetProc)
	}()

	autoClickInitialize()
	findColorInScreen()
	autoClick()

	for {
		time.Sleep(time.Second)
		select {
		case <-getmousepress:
		default:

		}
	}
}

func getuserdata() bool {
	if !fileExists("setting.json") {
		file, err := os.Create("setting.json")
		if err != nil {
			log.Println("오류발생",err)
			return false
		}
		firestStart()
		err = json.NewEncoder(file).Encode(getproc.Settings)
		if err != nil {
			log.Println("세팅파일 저장실패", err)
			return false
		}
	}
	file, err := os.Open("setting.json")
	if err != nil {
		log.Println("세팅파일 열기 실패", err)
		return false
	}
	err = json.NewDecoder(file).Decode(getproc.Settings)
	if err != nil {
		log.Println("세팅파일 설정 실패", err)
		return false
	}
	if getproc.Settings.Defalut.ColorCount == 0 {
		getproc.Settings.Defalut.ColorCount = 1
	}
	if getproc.Settings.Hotkeys.On != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.On, "|")
		errh := Sethotkey("0", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Off != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Off, "|")
		errh := Sethotkey("1", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Red != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Red, "|")
		errh := Sethotkey("2", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Blue != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Blue, "|")
		errh := Sethotkey("3", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Solo != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Solo, "|")
		errh := Sethotkey("4", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Exit != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Exit, "|")
		errh := Sethotkey("5", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Autoon != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Autoon, "|")
		errh := Sethotkey("6", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Autooff != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Autooff, "|")
		errh := Sethotkey("7", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Work != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Work, "|")
		errh := Sethotkey("8", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Autowork != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Autowork, "|")
		errh := Sethotkey("9", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Autofire != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Autofire, "|")
		errh := Sethotkey("10", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Bodyonly != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Bodyonly, "|")
		errh := Sethotkey("11", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	if getproc.Settings.Hotkeys.Headonly != "" {
		splitdata := strings.Split(getproc.Settings.Hotkeys.Headonly, "|")
		errh := Sethotkey("12", splitdata[0], splitdata[1])
		if errh !=nil {
			println(errh.Error())
		}
	}
	colordepint, _ := strconv.Atoi(getproc.Settings.Advance.Shade)
	getproc.Settings.Advance.Shadeint = int(colordepint)
	switch getproc.Settings.Defalut.Lock {
	case "X축고정":
		getproc.Settings.Defalut.Lockstate = 0
	case "Y축고정":
		getproc.Settings.Defalut.Lockstate = 1
	default:
		getproc.Settings.Defalut.Lockstate = 2
	}

	return true
}

func saveuserdata() {
	file, err := os.Create("setting.json")
	if err != nil {
		log.Println("설정 파일 저장실패", err)
		return
	}
	json.NewEncoder(file).Encode(getproc.Settings)
	runtime.GC()
}
func resetcolor() {
	getproc.Settings.Colors.RedTeamBody = nil
	getproc.Settings.Colors.RedTeamHead = nil
	getproc.Settings.Colors.BlueTeamHead = nil
	getproc.Settings.Colors.BlueTeamBody = nil

	getproc.Settings.Colors.NRTB = nil
	getproc.Settings.Colors.NRTH = nil
	getproc.Settings.Colors.NBTB = nil
	getproc.Settings.Colors.NBTH = nil

	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{235, 20, 20})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{215, 20, 20})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{195, 20, 20})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{175, 20, 20})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{155, 20, 20})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{135, 20, 20})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{115, 20, 20})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{95, 20, 20})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{235, 30, 30})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{215, 30, 30})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{195, 30, 30})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{175, 30, 30})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{155, 30, 30})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{135, 30, 30})
	getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{115, 30, 30})

	getproc.Settings.Colors.NRTH = append(getproc.Settings.Colors.NRTH, getproc.ColorArray{235, 235, 0})
	getproc.Settings.Colors.NRTH = append(getproc.Settings.Colors.NRTH, getproc.ColorArray{215, 215, 0})
	getproc.Settings.Colors.NRTH = append(getproc.Settings.Colors.NRTH, getproc.ColorArray{195, 195, 0})
	getproc.Settings.Colors.NRTH = append(getproc.Settings.Colors.NRTH, getproc.ColorArray{175, 175, 0})
	getproc.Settings.Colors.NRTH = append(getproc.Settings.Colors.NRTH, getproc.ColorArray{155, 155, 0})
	getproc.Settings.Colors.NRTH = append(getproc.Settings.Colors.NRTH, getproc.ColorArray{135, 135, 0})
	getproc.Settings.Colors.NRTH = append(getproc.Settings.Colors.NRTH, getproc.ColorArray{115, 115, 0})
	getproc.Settings.Colors.NRTH = append(getproc.Settings.Colors.NRTH, getproc.ColorArray{95, 95, 0})
	getproc.Settings.Colors.NRTH = append(getproc.Settings.Colors.NRTH, getproc.ColorArray{75, 75, 0})

	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{20, 20, 235})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{20, 20, 215})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{20, 20, 195})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{20, 20, 175})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{20, 20, 155})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{20, 20, 135})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{20, 20, 115})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{20, 20, 95})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{40, 40, 235})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{40, 40, 215})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{40, 40, 195})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{40, 40, 175})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{40, 40, 155})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{40, 40, 135})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{50, 50, 235})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{50, 50, 215})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{50, 50, 195})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{50, 50, 175})
	getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{50, 50, 155})

	getproc.Settings.Colors.NBTH = append(getproc.Settings.Colors.NBTH, getproc.ColorArray{0, 235, 0})
	getproc.Settings.Colors.NBTH = append(getproc.Settings.Colors.NBTH, getproc.ColorArray{0, 215, 0})
	getproc.Settings.Colors.NBTH = append(getproc.Settings.Colors.NBTH, getproc.ColorArray{0, 195, 0})
	getproc.Settings.Colors.NBTH = append(getproc.Settings.Colors.NBTH, getproc.ColorArray{0, 175, 0})
	getproc.Settings.Colors.NBTH = append(getproc.Settings.Colors.NBTH, getproc.ColorArray{0, 155, 0})
	getproc.Settings.Colors.NBTH = append(getproc.Settings.Colors.NBTH, getproc.ColorArray{0, 135, 0})
	getproc.Settings.Colors.NBTH = append(getproc.Settings.Colors.NBTH, getproc.ColorArray{0, 115, 0})
	getproc.Settings.Colors.NBTH = append(getproc.Settings.Colors.NBTH, getproc.ColorArray{0, 95, 0})
	getproc.Settings.Colors.NBTH = append(getproc.Settings.Colors.NBTH, getproc.ColorArray{0, 75, 0})

}
func firestStart() {
	resetcolor()
	getproc.Settings.Advance.Shade = "20"
	getproc.Settings.Advance.Minaccel = "5"
	getproc.Settings.Advance.Minx = 5
	getproc.Settings.Advance.Miny = 5

	getproc.Settings.Accel.Short = 700
	getproc.Settings.Accel.Middle = 400
	getproc.Settings.Accel.Long = 200

	getproc.Settings.Defalut.Lock = "X축Y축고정"
	getproc.Settings.Defalut.Xx = 0
	getproc.Settings.Defalut.Yy = 0
	getproc.Settings.Defalut.Maxrec = "30"
	getproc.Settings.Defalut.Randomworkstatus = false
	getproc.Settings.Defalut.Randomwork = 0
	getproc.Settings.Defalut.Delaystatus = "0"
	getproc.Settings.Defalut.Delay = 0
	getproc.Settings.Defalut.Autofirecount = "0"
	getproc.Settings.Defalut.Autodelay = "0"
	getproc.Settings.Defalut.Autodelaystatus = "0"
	getproc.Settings.Defalut.Onlyauto = "0"
	getproc.Settings.Defalut.X2 = "0"

	getproc.Settings.Rect.View = "0"
	getproc.Settings.Rect.Width = "80"
	getproc.Settings.Rect.Heigh = "40"
}




func serverInitialize() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		webServerStart(dir)
	}()
}
func autoClickInitialize() {
	go func() {
		for {
			select {
			case _getstatus = <-_GetProc:
			case autoclickstate = <- autoclick:
			default:
			}
			time.Sleep(time.Nanosecond)
			//fmt.Printf("%v \n",autoclickstate)
			autoclickstate = false
		}
	}()
}

func mouseMove(xx,yy int32 ) {
	var input win.MOUSE_INPUT
	input.Type = win.INPUT_MOUSE
	input.Mi = win.MOUSEINPUT{Dx: int32(xx), Dy: int32(yy), DwFlags: win.MOUSEEVENTF_MOVE, DwExtraInfo: uintptr(0)}
	win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
}

func findColorInScreen() {
	go func() {
		getproc.Checkifpress = true
		getproc.CheckTabpress = true
		getproc.Keypress = true
		for {
			if getproc.Settings.Defalut.Onlyauto == "1" {
				time.Sleep(time.Nanosecond)
				continue
			}
			if getproc.Settings.Defalut.Randomworkstatus {
				a := rand.Intn(100)
				if a >= getproc.Settings.Defalut.Randomwork {
					time.Sleep(time.Nanosecond)
					continue
				}
			}

			if _getstatus.HWND != 0 && login == 1 && getproc.Keypress && getproc.Checkstate && getproc.CheckTabpress {
				x, y := getproc.Screencapture(int(_getstatus.X), int(_getstatus.Y), int(_getstatus.W), int(_getstatus.H))
				if x != -1 || y != -1 {
					win.GetCursorPos(&curpos)
					m, _ := strconv.Atoi(getproc.Settings.Advance.Minaccel)
					s := getproc.Settings.Accel.Short
					mm := getproc.Settings.Accel.Middle
					l := getproc.Settings.Accel.Long

					xx := x + int(_getstatus.X) - int(curpos.X)
					if xx >= m && xx <= 30 {
						xx = int(float32(xx) * (float32(s) / 1000))
					} else if xx >= 31 && xx <= 80 {
						xx = int(float32(xx) * (float32(mm) / 1000))
					} else if xx >= 81 {
						xx = int(float32(xx) * (float32(l) / 1000))
					} else if xx <= -m && xx >= -30 {
						xx = int(float32(xx) * (float32(s) / 1000))
					} else if xx <= -31 && xx >= -80 {
						xx = int(float32(xx) * (float32(mm) / 1000))
					} else if xx <= -81 {
						xx = int(float32(xx) * (float32(l) / 1000))
					}

					yy := y + int(_getstatus.Y) - int(curpos.Y)
					if yy >= m && yy <= 10 {
						yy = int(float32(yy) * (float32(s) / 1000))
					} else if yy >= 11 && yy <= 80 {
						yy = int(float32(yy) * (float32(mm) / 1000))
					} else if yy >= 31 && yy <= 80 {
						yy = int(float32(yy) * (float32(mm) / 1000))
					} else if yy >= 81 {
						yy = int(float32(yy) * (float32(l) / 1000))
					} else if yy <= -m && yy >= -30 {
						yy = int(float32(yy) * (float32(s) / 1000))
					} else if yy <= -31 && yy >= -80 {
						yy = int(float32(yy) * (float32(mm) / 1000))
					} else if yy <= -81 {
						yy = int(float32(yy) * (float32(l) / 1000))
					}

					switch getproc.Settings.Defalut.Lockstate {
					case 0:
						yy = 0
					case 1:
						xx = 0
					default:
						xx = xx
						yy= yy
					}

					xx += getproc.Settings.Defalut.Xx
					yy +=  getproc.Settings.Defalut.Yy

					mouseMove(int32(xx), int32(yy))
				}

				if getproc.Settings.Defalut.Delaystatus == "1" && getproc.Settings.Defalut.Delay != 0 {
					time.Sleep(time.Duration(getproc.Settings.Defalut.Delay) * time.Millisecond)
				} else {
					time.Sleep(time.Nanosecond)
				}
			} else {
				time.Sleep(time.Nanosecond)
			}
		}
	}()
}
func autoClick() {
	go func() {
		getproc.CheckTabpress = true
		var checksleepval = uint32(0)
		var checkcount = 0
		for {
			if getproc.Settings.Defalut.Autodelaystatus == "1"  {
				if getproc.Settings.Defalut.Autodelay != "0" {
					checksleepval = 50
				}
				dtime, _ := strconv.Atoi(getproc.Settings.Defalut.Autodelay)
				checksleepval = uint32(dtime)
				ased, _ := strconv.Atoi(getproc.Settings.Defalut.Autofirecount)
				checkcount = ased
			}
			if getproc.Checkautostate {
				if _getstatus.HWND != 0 && login == 1  && getproc.CheckTabpress {
					if getproc.Autoclick(int(_getstatus.Ax), int(_getstatus.Ay)) {
						var input win.MOUSE_INPUT
						input.Type = win.INPUT_MOUSE

						for i := 0; i<checkcount; i++ {
							input.Mi = win.MOUSEINPUT{DwFlags: win.MOUSEEVENTF_LEFTDOWN , DwExtraInfo: uintptr(0),Time:checksleepval}
							win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
						}
						input.Mi = win.MOUSEINPUT{DwFlags:  win.MOUSEEVENTF_LEFTUP , DwExtraInfo: uintptr(0)}
						win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(win.MOUSE_INPUT{})))
					}
				}
			} else {
				time.Sleep(time.Second)
			}
			time.Sleep(time.Nanosecond)

		}
	}()
}