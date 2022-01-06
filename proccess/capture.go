package getproc

import (
	"errors"
	"fmt"
	"github.com/lxn/win"
	"image"
	"image/color"
	"strconv"
	"strings"
	"unsafe"
)

type ColorArray struct {
	R int
	G int
	B int
}

var (
	Colorlist []string
	Colorarrays []ColorArray
    curcolor ColorArray
	allcolor = 10
	bodyhead = 5
	findcount = 0
)
func Screencapture(x,y,width,height int) (int,int) {
	//hwnd := win.GetDesktopWindow()

	hdc := win.GetDC(0)
	if hdc == 0 { fmt.Printf("Error code : 0002\n")
		return -1,-1}
	defer win.ReleaseDC(0, hdc)
	memory_device := win.CreateCompatibleDC(hdc)
	if memory_device == 0 { fmt.Printf("Error code : 0003\n")
		return -1,-1}
	defer win.DeleteDC(memory_device)
	bitmap := win.CreateCompatibleBitmap(hdc,int32(width), int32(height))
	if bitmap == 0 { fmt.Printf("Error code : 0004\n")
		return -1,-1}
	defer win.DeleteObject(win.HGDIOBJ(bitmap))
	var header win.BITMAPINFOHEADER
	header.BiSize = uint32(unsafe.Sizeof(header))
	header.BiPlanes = 1
	header.BiBitCount = 32
	header.BiWidth = int32(width)
	header.BiHeight = int32(-height)
	header.BiCompression = win.BI_RGB
	header.BiSizeImage = 0
	bitmapdatasize := uintptr(((int64(width) * int64(header.BiBitCount) + 31) / 32) * 4 * int64(height))
	hmem := win.GlobalAlloc(win.GMEM_MOVEABLE, bitmapdatasize)
	defer win.GlobalFree(hmem)
	memptr := win.GlobalLock(hmem)
	defer win.GlobalUnlock(hmem)
	old := win.SelectObject(memory_device, win.HGDIOBJ(bitmap))
	if old == 0 {
		fmt.Printf("Error code : 0005\n")
		return -1,-1
	}
	defer win.SelectObject(memory_device, old)
	if !win.BitBlt(memory_device,0,0, int32(width), int32(height), hdc, int32(x), int32(y), win.SRCCOPY){
		fmt.Printf("Error code : 0006\n")
		return -1,-1
	}
	if win.GetDIBits(hdc, bitmap, 0, uint32(height), (*uint8)(memptr), (*win.BITMAPINFO)(unsafe.Pointer(&header)), win.DIB_RGB_COLORS) == 0 {
		fmt.Printf("Error code : 0007\n")
		return -1,-1
	}
	var findx []int
	var findy []int
	maxfindx := 0
	maxfindy := 0
	//findxcount := 2
	find := false
	finddone := false
	src := uintptr(memptr)

	//ran := rand.Intn(100)
	//if Settings.Defalut.Headonly == "1" || Settings.Defalut.Bodyonly == "1"{
	//	for y := 0; y < height; y++ {
	//		if finddone {
	//			maxfindy = maxfindy + 1
	//			finddone = false
	//		}
	//		if maxfindy == 2 {
	//			break
	//		}
	//		maxfindx =  0
	//		for x := 0; x < width; x++ {
	//			if maxfindx >= Settings.Defalut.ColorCount {
	//				findxcount = 2
	//				finddone = true
	//				break
	//			}
	//			//v0 := *(*uint8)(unsafe.Pointer(src)) //B
	//			//v1 := *(*uint8)(unsafe.Pointer(src + 1)) //G
	//			//v2 := *(*uint8)(unsafe.Pointer(src + 2)) //R
	//
	//			curcolor = ColorArray{R:int(*(*uint8)(unsafe.Pointer(src + 2))),G:int(*(*uint8)(unsafe.Pointer(src + 1))), B:int(*(*uint8)(unsafe.Pointer(src)))}
	//			//currentcolor := color.RGBA{v2,v1,v0,255}
	//			//for _,n := range Colorlist{
	//			//curcolor = ColorArray{R:*(*int16)(unsafe.Pointer(src + 2)),G:*(*int16)(unsafe.Pointer(src + 1)), B:*(*int16)(unsafe.Pointer(src))}
	//			for _,n :=range Colorarrays{
	//				fmt.Printf("%v %v %d \n", curcolor, n, Settings.Defalut.ColorCount)
	//				if newCompare(curcolor, n, Settings.Advance.Shadeint){
	//					if findxcount <= 0 {
	//						findxcount = 2
	//					}
	//					if findxcount >= 2 {
	//						findy = append(findy, y)
	//						findx = append(findx, x)
	//						maxfindx = maxfindx + 1
	//					if x+2 < width {
	//						x += 2
	//						src += 8
	//					}
	//					find = true
	//				}
	//					findxcount = findxcount -1
	//					break
	//				}
	//			}
	//			src +=4
	//			//elapsedTime := time.Since(startTime)
	//			//fmt.Printf("실행시간: %s\n", elapsedTime)
	//		}
	//	}
	//} else {

		for y := 0; y < height; y++ {

			if finddone {
				maxfindy = maxfindy + 1
				finddone = false
			}
			if maxfindy == 2 {
				break
			}
			maxfindx =  0
			for x := 0; x < width; x++ {
				if maxfindx >= Settings.Defalut.ColorCount {
					//findxcount = 2
					finddone = true
					src += uintptr((width - x ) * 4)
					break
				}
				curcolor = ColorArray{R:int(*(*uint8)(unsafe.Pointer(src + 2))),G:int(*(*uint8)(unsafe.Pointer(src + 1))), B:int(*(*uint8)(unsafe.Pointer(src)))}
				for _,n :=range Colorarrays{
					//fmt.Printf("%v %v %d \n", curcolor, n, Settings.Defalut.ColorCount)
					if newCompare(curcolor, n, Settings.Advance.Shadeint){
							findx = append(findx, x)
							findy = append(findy, y)
						maxfindx = maxfindx + 1
						find = true
						break
					}
				}
				src +=4

				//elapsedTime := time.Since(startTime)
				//fmt.Printf("실행시간: %s\n", elapsedTime)
			}
		}
	//}
	//else if ran >= 70 {
	//	for y := 0; y < height; y++ {
	//		if finddone {
	//			break
	//		}
	//		maxfindx =  0
	//		for x := 0; x < width; x++ {
	//			if maxfindx >= 2 {
	//				findxcount = 2
	//				finddone = true
	//				break
	//			}
	//			curcolor = ColorArray{R:*(*uint8)(unsafe.Pointer(src + 2)),G:*(*uint8)(unsafe.Pointer(src + 1)), B:*(*uint8)(unsafe.Pointer(src))}
	//			for _,n := range Colorarrays{
	//
	//				if newCompare(curcolor, n, Settings.Advance.Shadeint){
	//					if findxcount <= 0 {
	//						findxcount = 2
	//					}
	//					if findxcount == 2 {
	//						findy = append(findy, y)
	//						findx = append(findx, x)
	//						maxfindx = maxfindx + 1
	//						if x+2 < width {
	//							x += 2
	//							src += 8
	//						}
	//						find = true
	//					}
	//					findxcount = findxcount -1
	//					break
	//				}
	//			}
	//			src +=4
	//			//elapsedTime := time.Since(startTime)
	//			//fmt.Printf("실행시간: %s\n", elapsedTime)
	//		}
	//	}
	//}
	//fmt.Printf("x : %v  y : %v \n", len(findx), len(findy))

	if find {
		//fmt.Printf("%d %d  x : %v  y : %v \n", average(findx), average(findy), findx, findy )

		return average(findx), average(findy)
	}
	return -1,-1

}

func Autoclick(x,y int) bool {
	//hwnd := win.GetDesktopWindow()
	width := 10
	height := 15
	hdc := win.GetDC(0)
	if hdc == 0 { fmt.Printf("Error code : 0002\n")}
	defer win.ReleaseDC(0, hdc)
	memory_device := win.CreateCompatibleDC(hdc)
	if memory_device == 0 { fmt.Printf("Error code : 0003\n")}
	defer win.DeleteDC(memory_device)
	bitmap := win.CreateCompatibleBitmap(hdc,int32(width), int32(height))
	if bitmap == 0 { fmt.Printf("Error code : 0004\n")}
	defer win.DeleteObject(win.HGDIOBJ(bitmap))
	var header win.BITMAPINFOHEADER
	header.BiSize = uint32(unsafe.Sizeof(header))
	header.BiPlanes = 1
	header.BiBitCount = 32
	header.BiWidth = int32(width)
	header.BiHeight = int32(-height)
	header.BiCompression = win.BI_RGB
	header.BiSizeImage = 0
	bitmapdatasize := uintptr(((int64(width) * int64(header.BiBitCount) + 31) / 32) * 4 * int64(height))
	hmem := win.GlobalAlloc(win.GMEM_MOVEABLE, bitmapdatasize)
	defer win.GlobalFree(hmem)
	memptr := win.GlobalLock(hmem)
	defer win.GlobalUnlock(hmem)
	old := win.SelectObject(memory_device, win.HGDIOBJ(bitmap))
	if old == 0 {
		fmt.Printf("Error code : 0005\n")
	}
	defer win.SelectObject(memory_device, old)
	if !win.BitBlt(memory_device,0,0, int32(width), int32(height), hdc, int32(x), int32(y), win.SRCCOPY){
		fmt.Printf("Error code : 0006\n")
	}
	if win.GetDIBits(hdc, bitmap, 0, uint32(height), (*uint8)(memptr), (*win.BITMAPINFO)(unsafe.Pointer(&header)), win.DIB_RGB_COLORS) == 0 {
		fmt.Printf("Error code : 0007\n")
	}
	src := uintptr(memptr)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			curcolor = ColorArray{R:int(*(*uint8)(unsafe.Pointer(src + 2))),G:int(*(*uint8)(unsafe.Pointer(src + 1))), B:int(*(*uint8)(unsafe.Pointer(src)))}
			for _,n :=range Colorarrays{
				if newCompare(curcolor, n, Settings.Advance.Shadeint){
					return true
				}
			}
			src +=4
		}
	}
	return false
}

func Compare(cur color.RGBA, dest string, shade int) bool{
	splitz := strings.Split(dest,",")
	r,_ := strconv.Atoi(splitz[0])
	g,_ := strconv.Atoi(splitz[1])
	b,_ := strconv.Atoi(splitz[2])

	return (intabs(int(cur.R) - r) <= shade && intabs(int(cur.G) - g) <= shade && intabs(int(cur.B) - b) <= shade)
}

func newCompare(cur ColorArray, dest ColorArray, shade int) bool{
	return (uintabs(cur.R - dest.R) <= shade && uintabs(cur.G - dest.G) <= shade && uintabs(cur.B - dest.B) <= shade)
}
func intabs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}
func uintabs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}
func addcolorlist(clist []color.RGBA, col color.RGBA)  []color.RGBA{
	clist = append(clist, col)
	return  clist
}
func average(xs[]int)  int {
	total:=0
	for _,v:=range xs {
		total +=  v
	}
	if total == 0 {
		return 0
	}
	return total/int(len(xs))
}

func createimg(rect image.Rectangle) (img *image.RGBA, e error) {
	img = nil
	e = errors.New("이미지픽셀 생성실패")

	defer func() {
		err := recover()
		if err == nil{
			e = nil
		}
	}()
	img = image.NewRGBA(rect)
	return img, e
}
