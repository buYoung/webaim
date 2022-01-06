package getproc

import (
	"github.com/lxn/win"
	hook "github.com/robotn/gohook"
	"golang.org/x/sys/windows"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

var (
	mod 					= windows.NewLazyDLL("user32.dll")
	procGetWindowText   	= mod.NewProc("GetWindowTextW")
	procGetWindowTextLength = mod.NewProc("GetWindowTextLengthW")
	DrawRect   	= mod.NewProc("FrameRect")
	GetWindowInfo   	= mod.NewProc("GetWindowInfo")
	Mouseevent   	= mod.NewProc("GetAsyncKeyState")
	getdoubleclicktime = mod.NewProc("GetDoubleClickTime")
	Isstop = 0
	Settings = &Allsettings{}
	adtimer time.Time = time.Now()
)
type Userinfo struct {
	Id string `json:"id"`
}
type Hotkeys struct {
	On string `json:"on"`
	Off string `json:"off"`
	Red string `json:"red"`
	Blue string `json:"blue"`
	Solo string `json:"solo"`
	Exit string `json:"exit"`
	Autoon string `json:"autoon"`
	Autooff string `json:"autooff"`
	Work string `json:"work"`
	Autowork string `json:"autowork"`
	Autofire string `json:"autofire"`
	Bodyonly string `json:"bodyonly"`
	Headonly string `json:"headonly"`
}
type Rect struct {
	Width string `json:"width"`
	Heigh string `json:"heigh"`
	View string `json:"view"`
}
type Defalut struct {
	Lock string `json:"lock"`
	Lockstate uint8 `json:"lockstate"`
	Xx int `json:"xx"`
	Yy int `json:"yy"`
	Maxrec string `json:"maxrec"`
	Randomworkstatus bool `json:"randomworkstatus"`
	Randomwork int `json:"randomwork"`
	Delaystatus string `json:"delaystatus"`
	Delay int `json:"delay"`
	Autofirecount string  `json:"autofirecount"`
	Autodelaystatus string `json:"autodelaystatus"`
	Autodelay string `json:"autodelay"`
	Onlyauto string `json:"onlyauto"`
	X2 string `json:"x2"`
	Bodyonly string `json:"bodyonly"`
	Headonly string `json:"headonly"`
	Othercolor string `json:"othercolor"`
	ColorCount int `json:"color_count"`
	Ghost int `json:"ghost"`
}
type Colors struct {
	RedTeamBody []string `json:"red_team_body"`
	RedTeamHead []string `json:"red_team_head"`
	NRTB []ColorArray `json:"nrtb"`
	NRTH []ColorArray `json:"nrth"`
	BlueTeamBody []string `json:"blue_team_body"`
	BlueTeamHead []string `json:"blue_team_head"`
	NBTB []ColorArray `json:"nbtb"`
	NBTH []ColorArray `json:"nbth"`
	NOC []ColorArray `json:"noc"`

}
type Accel struct {
	Short int `json:"short"`
	Middle int `json:"middle"`
	Long int `json:"long"`
}
type Advance struct {
	Minx int `json:"minx"`
	Miny int `json:"miny"`
	Minaccel string `json:"minaccel"`
	Shade string `json:"shade"`
	Shadeint int `json:"shadeint"`
	Debug string `json:"debug"`
}
type Skinpath struct {
	Spath string `json:"spath"`
	Skinpath string `json:"skinpath"`
}
type Allsettings struct {
	Userinfo Userinfo `json:"userinfo"`
	Hotkeys Hotkeys `json:"hotkeys"`
	Rect Rect `json:"rect"`
	Defalut Defalut `json:"defalut"`
	Colors Colors `json:"colors"`
	Accel Accel `json:"accel"`
	Advance Advance `json:"advance"`
}

type Procinfo struct {
	HWND win.HWND
	X int32
	Y int32
	W int32
	H int32
	Ax int32
	Ay int32
	Rect win.RECT
}
type RECT struct {
	Left int32
	Top int32
	Right int32
	Bottom int32
}

type (
	HANDLE uintptr
	_HWND      HANDLE
)
var(
	CheckTabpress bool
	Checkifpress bool
	CheckifApress bool
	Keypress = false
	Checkkeys uint16
)
func Doubclicktime() uint {
	time,_,_ := getdoubleclicktime.Call()
	return uint(time)
}
func Gettabkey(ch chan int ,ch1 chan int){
	s := hook.Start()
	defer hook.End()
	status := false
	for {
		select {
		case <- ch:
			status = true
			CheckifApress = false
			Checkifpress = false
			Keypress = false
		case <- ch1:
			status = false
			CheckifApress =true
			Checkifpress = true
			Keypress = true
		case i := <-s:

			if i.Rawcode == 9 {
				switch i.Kind {
				case 3,4:
					CheckTabpress = false
				case 5:
					CheckTabpress= true
				}
			} else if status {
				if i.Button == 0 {
					continue
				}
				if i.Kind == hook.MouseHold && i.Button == Checkkeys {
					Keypress = true
				}
				if (i.Kind == hook.MouseUp || i.Kind == hook.MouseDown) && i.Button == Checkkeys {
					Keypress = false
				}
			}  else if i.Rawcode == 87 {
				if Settings.Defalut.Ghost == 1 {
					if time.Since(adtimer).Milliseconds() >= 100 {
						var input_a win.KEYBD_INPUT
						input_a.Type = win.INPUT_KEYBOARD
						input_a.Ki.WVk = 0x41
						win.SendInput(1, unsafe.Pointer(&input_a), int32(unsafe.Sizeof(win.KEYBD_INPUT{})))
						var input_d win.KEYBD_INPUT
						input_d.Type = win.INPUT_KEYBOARD
						input_d.Ki.WVk = 0x44
						win.SendInput(1, unsafe.Pointer(&input_d), int32(unsafe.Sizeof(win.KEYBD_INPUT{})))
						adtimer = time.Now()
					}
				}
			} else {
				continue
			}
		default:
			Checkifpress =true
			CheckifApress =true
		}
		time.Sleep(time.Nanosecond)
	}
}

func Getkeystates(ch chan int) {
	go func() {
		Isstop = 2
		s := hook.Start()
		defer hook.End()
		for {
			select {
			case i:= <- s:
				if i.Button == 0 {
					continue
				}
				if i.Kind == hook.MouseHold && i.Button == Checkkeys {
					Keypress = true
				}
				if (i.Kind == hook.MouseUp || i.Kind == hook.MouseDown) && i.Button == Checkkeys {
					Keypress = false
				}
			default:
				continue
			}
			select {
			case <-ch:
				Isstop = 0
			default:
				continue
			}
			if Isstop == 0{
				hook.End()
				Keypress = true
				break
			}
		}

	}()
}


func GetWindowTextLength(hwnd win.HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(
		uintptr(hwnd))
	return int(ret)
}

func GetWindowText(hwnd win.HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1
	buf := make([]uint16, textLen)
	procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func _DrawRect(hdc win.HDC, lprc *win.RECT, brush win.HGDIOBJ) int {
	ret, _, _ := DrawRect.Call(uintptr(hdc), uintptr(unsafe.Pointer(lprc)), uintptr(brush))
	return int(ret)
}
func getWindow(funcName string) uintptr {
	proc := mod.NewProc(funcName)
	hwnd, _, _ := proc.Call()
	return hwnd
}

func FindActiveProcess(GetProc chan Procinfo){
	var makez Procinfo
	var rect win.RECT
	var rwidth,rheight int

	makez.HWND = 0
	for{
		hwnd := win.GetForegroundWindow()
		if hwnd != 0 {
			win.GetWindowRect(hwnd, &rect)
			rect.Bottom = rect.Bottom - rect.Top
			rect.Right = rect.Right - rect.Left //창크기
			makez.HWND = hwnd
			makez.Rect = rect
			rwidth, _ = strconv.Atoi(Settings.Rect.Width)
			rheight, _ = strconv.Atoi(Settings.Rect.Heigh)
			makez.X = (rect.Left) + (rect.Right / 2) - (int32(rwidth) / 2)
			makez.Y = (rect.Top) + (rect.Bottom / 2) - (int32(rheight) / 2)
			makez.Ax = (rect.Left) + (rect.Right / 2) - 7
			makez.Ay = (rect.Top) + (rect.Bottom / 2) - 3
			makez.W = int32(rwidth)
			makez.H = int32(rheight)


			select {
			case GetProc <- makez:
			default:
				time.Sleep(1000 * time.Millisecond)
				continue
			}
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
