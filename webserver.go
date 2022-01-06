package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/buYoung/webaim/aes256"
	getproc "github.com/buYoung/webaim/proccess"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

func webServerStart(dir string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Static("/_inc", "template/_inc")
	router.Static("/webfonts", "template/webfonts")
	router.Static("/icons", "template/icons")
	router.Static("/css", "template/css")
	router.Static("/js", "template/js")
	router.LoadHTMLGlob("template/*.html")
	router.GET("/", func(c *gin.Context) {
		logininfo := ""
		if fileExists(path.Join(dir, "save.ini")) {
			d, _ := ioutil.ReadFile(path.Join(dir, "save.ini"))
			loadinfo := aes256.Decrypt(string(d), "qa1we89b4q89r1q6w16q#!$124")
			logininfo = base64.StdEncoding.EncodeToString([]byte(loadinfo))
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"login":    login,
			"weapon":   weapon,
			"chracter": chracter,
			"id":       loginid,
			"saveinfo": logininfo,
		})
	})
	router.GET("/index.html", func(c *gin.Context) {
		logininfo := ""
		if fileExists(path.Join(dir, "save.ini")) {
			d, _ := ioutil.ReadFile(path.Join(dir, "save.ini"))
			loadinfo := aes256.Decrypt(string(d), "qa1we89b4q89r1q6w16q#!$124")
			logininfo = base64.StdEncoding.EncodeToString([]byte(loadinfo))
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"login":    login,
			"weapon":   weapon,
			"chracter": chracter,
			"id":       getproc.Settings.Userinfo.Id,
			"saveinfo": logininfo,
		})
	})
	router.GET("/set1-1.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "set1-1.html", gin.H{
			"login":    login,
			"on":       getproc.Settings.Hotkeys.On,
			"off":      getproc.Settings.Hotkeys.Off,
			"red":      getproc.Settings.Hotkeys.Red,
			"blue":     getproc.Settings.Hotkeys.Blue,
			"solo":     getproc.Settings.Hotkeys.Solo,
			"exit":     getproc.Settings.Hotkeys.Exit,
			"autoon":   getproc.Settings.Hotkeys.Autoon,
			"autooff":  getproc.Settings.Hotkeys.Autooff,
			"work":     getproc.Settings.Hotkeys.Work,
			"autowork": getproc.Settings.Hotkeys.Autowork,
			"autofire": getproc.Settings.Hotkeys.Autofire,
			"body":     getproc.Settings.Hotkeys.Bodyonly,
			"head":     getproc.Settings.Hotkeys.Headonly,
		})
	})
	router.GET("/set1-2.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "set1-2.html", gin.H{
			"login":  login,
			"width":  getproc.Settings.Rect.Width,
			"height": getproc.Settings.Rect.Heigh,
			"look":   getproc.Settings.Rect.View,
		})
	})
	router.GET("/set1-3.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "set1-3.html", gin.H{
			"login":            login,
			"lock":             getproc.Settings.Defalut.Lock,
			"x":                getproc.Settings.Defalut.Xx,
			"y":                getproc.Settings.Defalut.Yy,
			"maxrec":           getproc.Settings.Defalut.Maxrec,
			"randomwork":       getproc.Settings.Defalut.Randomwork,
			"randomworkstatus": getproc.Settings.Defalut.Randomworkstatus,
			"delay":            getproc.Settings.Defalut.Delay,
			"delaystatus":      getproc.Settings.Defalut.Delaystatus,
			"autofirecount":    getproc.Settings.Defalut.Autofirecount,
			"autodelay":        getproc.Settings.Defalut.Autodelay,
			"autodelaystatus":  getproc.Settings.Defalut.Autodelaystatus,
			"onlyauto":         getproc.Settings.Defalut.Onlyauto,
			"x2":               getproc.Settings.Defalut.X2,
			"head":             getproc.Settings.Defalut.Headonly,
			"body":             getproc.Settings.Defalut.Bodyonly,
			"other":            getproc.Settings.Defalut.Othercolor,
			"count":            getproc.Settings.Defalut.ColorCount,
			"ghost":            getproc.Settings.Defalut.Ghost,
		})
	})
	router.GET("/set2-1.html", func(c *gin.Context) {
		res2B, _ := json.Marshal(getproc.Settings.Colors)
		c.HTML(http.StatusOK, "set2-1.html", gin.H{"login": login,
			"colors": string(res2B)})
	})
	router.GET("/set2-2.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "set2-2.html", gin.H{
			"login":  login,
			"short":  getproc.Settings.Accel.Short,
			"middle": getproc.Settings.Accel.Middle,
			"long":   getproc.Settings.Accel.Long,
		})
	})
	router.GET("/set2-3.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "set2-3.html", gin.H{
			"login":     login,
			"dontmovex": getproc.Settings.Advance.Minx,
			"dontmovey": getproc.Settings.Advance.Miny,
			"minaccel":  getproc.Settings.Advance.Minaccel,
			"colordep":  getproc.Settings.Advance.Shade,
		})
	})

	router.POST("/unhotkey", func(c *gin.Context) {
		index := c.PostForm("index")
		unhoteky(index)
		saveuserdata()

		c.JSON(http.StatusOK, "complete")
	})
	router.POST("/hotkey", func(c *gin.Context) {
		index := c.PostForm("index")
		key := c.PostForm("keycode")
		keystr := c.PostForm("key")
		err := Sethotkey(index, key, keystr)
		if err != nil {
			println(err)
			if strings.Contains(err.Error(), "failed to register hotkey") {
				c.PureJSON(http.StatusOK, `{"error":1,"result": "이미 등록된 키조합입니다.","detail" : "`+keystr+` 키조합는 이미 등록되어있습니다"}`)
			}
			recover()
		} else {
			saveuserdata()
			c.PureJSON(http.StatusOK, `{"error":0,"result": "키조합이 등록되었습니다.","detail" : "`+keystr+` 키조합이 등록되어있습니다"}`)
		}

	})
	router.POST("/rect", func(c *gin.Context) {
		width := c.PostForm("width")
		height := c.PostForm("height")
		look := c.PostForm("look")
		returnval := ""
		if width != "" {
			getproc.Settings.Rect.Width = width
			returnval = `{"error":0,"result": "넓이 설정완료","detail" : "넓이 ` + width + ` 설정완료 "}`
		}
		if height != "" {
			getproc.Settings.Rect.Heigh = height
			returnval = `{"error":0,"result": "높이 설정완료","detail" : "높이 ` + height + ` 설정완료 "}`
		}
		if look != "" {
			getproc.Settings.Rect.View = look
			returnval = `{"error":0,"result": "인식범위표시 설정완료","detail" : "인식범위표시 설정완료 "}`
		}
		saveuserdata()

		c.PureJSON(http.StatusOK, returnval)
	})
	router.POST("/defalut", func(c *gin.Context) {
		lock := c.PostForm("lock")
		x := c.PostForm("x")
		y := c.PostForm("y")
		maxrec := c.PostForm("maxrec")
		randomwork := c.PostForm("randomwork")
		randomworkstatus := c.PostForm("randomworkstatus")
		delay := c.PostForm("delay")
		delaystatus := c.PostForm("delaystatus")
		autofirecount := c.PostForm("autofirecount")
		autodelay := c.PostForm("autodelay")
		autodelaystatus := c.PostForm("autodelaystatus")
		onlyauto := c.PostForm("onlyauto")
		x2 := c.PostForm("x2")
		head := c.PostForm("headonly")
		body := c.PostForm("bodyonly")
		other := c.PostForm("othercoloronly")
		colorcount := c.PostForm("colorcount")
		ghost := c.PostForm("ghost")
		returnval := ""
		if lock != "" {
			getproc.Settings.Defalut.Lock = lock
			switch getproc.Settings.Defalut.Lock {
			case "X축고정":
				getproc.Settings.Defalut.Lockstate = 0
			case "Y축고정":
				getproc.Settings.Defalut.Lockstate = 1
			default:
				getproc.Settings.Defalut.Lockstate = 2
			}

			returnval = `{"error":0,"result": "고정방식 설정완료","detail" : "고정방식 ` + lock + ` 설정완료"}`
		}
		if x != "" {
			zb, _ := strconv.Atoi(x)
			getproc.Settings.Defalut.Xx = zb
			returnval = `{"error":0,"result": "x축 위치 설정완료","detail" : "x축 위치 ` + x + ` 설정완료"}`
		}
		if y != "" {
			zb, _ := strconv.Atoi(y)
			getproc.Settings.Defalut.Yy = zb
			returnval = `{"error":0,"result": "y축 위치 설정완료","detail" : "y축 위치 ` + y + ` 설정완료"}`
		}
		if maxrec != "" {
			getproc.Settings.Defalut.Maxrec = maxrec
			returnval = `{"error":0,"result": "최대인식거리 설정완료","detail" : "최대인식거리 ` + maxrec + ` 설정완료"}`
		}
		if randomwork != "" {
			a, _ := strconv.Atoi(randomwork)
			getproc.Settings.Defalut.Randomwork = a
			returnval = `{"error":0,"result": "랜덤작동 설정완료","detail" : "랜덤작동 ` + randomwork + `% 설정완료"}`
		}
		if randomworkstatus != "" {
			if randomworkstatus == "0" {
				getproc.Settings.Defalut.Randomworkstatus = false
			} else {
				getproc.Settings.Defalut.Randomworkstatus = true
			}
			returnval = `{"error":0,"result": "랜덤작동상태 설정완료","detail" : "랜덤작동상태 설정완료"}`
		}
		if delay != "" {
			val, _ := strconv.Atoi(delay)
			getproc.Settings.Defalut.Delay = val
			returnval = `{"error":0,"result": "딜레이 설정완료","detail" : "딜레이 ` + delay + ` 설정완료"}`
		}
		if delaystatus != "" {
			getproc.Settings.Defalut.Delaystatus = delaystatus
			returnval = `{"error":0,"result": "딜레이상태 설정완료","detail" : "딜레이상태 설정완료"}`
		}
		if autofirecount != "" {
			getproc.Settings.Defalut.Autofirecount = autofirecount
			returnval = `{"error":0,"result": "오토샷 발사횟수 설정완료","detail" : "오토샷 발사횟수 ` + autofirecount + ` 설정완료"}`
		}
		if autodelay != "" {
			getproc.Settings.Defalut.Autodelay = autodelay
			returnval = `{"error":0,"result": "오토샷 딜레이 설정완료","detail" : "오토샷 딜레이 ` + autodelay + ` 설정완료"}`
		}
		if autodelaystatus != "" {
			getproc.Settings.Defalut.Autodelaystatus = autodelaystatus
			returnval = `{"error":0,"result": "오토샷 딜레이 설정완료","detail" : "오토샷 딜레이상태 설정완료"}`
		}
		if onlyauto != "" {
			getproc.Settings.Defalut.Onlyauto = onlyauto
			returnval = `{"error":0,"result": "오토샷만 사용하기 설정완료","detail" : "오토샷만 사용하기 설정완료"}`
		}
		if x2 != "" {
			getproc.Settings.Defalut.X2 = x2
			returnval = `{"error":0,"result": "2배속 작동 설정완료","detail" :  "2배속 작동설정완료"}`
		}
		if body != "" {
			getproc.Settings.Defalut.Bodyonly = body
			returnval = `{"error":0,"result": "몸전용 설정완료","detail" :  "몸전용 작동설정완료"}`
		}
		if head != "" {
			getproc.Settings.Defalut.Headonly = head
			returnval = `{"error":0,"result": "머리전용 설정완료","detail" :  "머리전용 작동설정완료"}`
		}
		if other != "" {
			getproc.Settings.Defalut.Othercolor = head
			returnval = `{"error":0,"result": "기타색 설정완료","detail" :  "기타색 작동설정완료"}`
		}
		if colorcount != "" {
			a, _ := strconv.Atoi(colorcount)
			getproc.Settings.Defalut.ColorCount = a
			returnval = `{"error":0,"result": ` + colorcount + `" 설정완료","detail" :  "색상찾기갯수 설정완료"}`
		}
		if ghost != "" {
			a, _ := strconv.Atoi(ghost)
			getproc.Settings.Defalut.Ghost = a
			returnval = `{"error":0,"result": ` + ghost + `" 설정완료","detail" :  "a/d키 누르기 설정완료"}`
		}

		saveuserdata()

		c.PureJSON(http.StatusOK, returnval)
	})
	router.POST("/color", func(c *gin.Context) {
		color := c.PostForm("color")
		colorsplit := strings.Split(color, "|")
		getproc.Settings.Colors.NRTB = nil
		getproc.Settings.Colors.NRTH = nil
		getproc.Settings.Colors.NBTB = nil
		getproc.Settings.Colors.NBTH = nil
		getproc.Settings.Colors.NOC = nil

		err = json.Unmarshal([]byte(colorsplit[0]), &getproc.Settings.Colors.NRTB)
		err = json.Unmarshal([]byte(colorsplit[1]), &getproc.Settings.Colors.NRTH)
		err = json.Unmarshal([]byte(colorsplit[2]), &getproc.Settings.Colors.NBTB)
		err = json.Unmarshal([]byte(colorsplit[3]), &getproc.Settings.Colors.NBTH)
		err = json.Unmarshal([]byte(colorsplit[4]), &getproc.Settings.Colors.NOC)
		returnval := `{"error":0,"result": "인식색상 변경완료","detail" : "인식색상 변경완료 "}`
		if err != nil {
			returnval = `{"error":1,"result": "인식색상 변경실패","detail" : "` + err.Error() + `"}`
		} else {
			saveuserdata()
		}

		c.PureJSON(http.StatusOK, returnval)
	})
	router.POST("/colorstate", func(c *gin.Context) {
		indexs := c.PostForm("index")
		colorrgb := c.PostForm("rgb")

		var result string
		indexnum, _ := strconv.Atoi(indexs)
		var R = int(0)
		var G = int(0)
		var B = int(0)
		if indexnum >= 6 && indexnum <= 10 {
			var getsplit = strings.Split(colorrgb, ",")
			numintr, _ := strconv.Atoi(getsplit[0])
			numintg, _ := strconv.Atoi(getsplit[1])
			numintb, _ := strconv.Atoi(getsplit[2])
			R = int(numintr)
			G = int(numintg)
			B = int(numintb)
		}
		switch indexnum {
		case 1: //색초기화
			resetcolor()
			result = `{"error":0,"result": "색상태","detail" : "모든색상 초기화 완료"}`
		case 2: //레드몸 색상제거
			getproc.Settings.Colors.NRTB = nil
			result = `{"error":0,"result": "색상태","detail" : "레드몸색상 전부 삭제완료"}`
		case 3: //레드머리 색상제거
			getproc.Settings.Colors.NRTH = nil
			result = `{"error":0,"result": "색상태","detail" : "레드머리색상 전부 삭제완료"}`
		case 4: //블루몸 색상제거
			getproc.Settings.Colors.NBTB = nil
			result = `{"error":0,"result": "색상태","detail" : "블루몸색상 전부 삭제완료"}`
		case 5: //블루머리 색상제거
			getproc.Settings.Colors.NBTH = nil
			result = `{"error":0,"result": "색상태","detail" : "블루머리색상 전부 삭제완료"}`
		case 6: //레드몸 색상추가
			//getproc.Settings.Colors.RedTeamBody = append(getproc.Settings.Colors.RedTeamBody,colorrgb)
			getproc.Settings.Colors.NRTB = append(getproc.Settings.Colors.NRTB, getproc.ColorArray{R, G, B})
			result = `{"error":0,"result": "색상태","detail" : "레드몸색상 추가완료"}`
		case 7: //레드머리 색상추가
			//getproc.Settings.Colors.RedTeamHead = append(getproc.Settings.Colors.RedTeamHead,colorrgb)
			getproc.Settings.Colors.NRTH = append(getproc.Settings.Colors.NRTH, getproc.ColorArray{R, G, B})
			result = `{"error":0,"result": "색상태","detail" : "레드머리색상 추가완료"}`
		case 8: //블루몸 색상추가
			//getproc.Settings.Colors.BlueTeamBody = append(getproc.Settings.Colors.BlueTeamBody,colorrgb)
			getproc.Settings.Colors.NBTB = append(getproc.Settings.Colors.NBTB, getproc.ColorArray{R, G, B})
			result = `{"error":0,"result": "색상태","detail" : "블루몸색상 추가완료"}`
		case 9: //블루머리 색상추가
			//getproc.Settings.Colors.BlueTeamHead = append(getproc.Settings.Colors.BlueTeamHead,colorrgb)
			getproc.Settings.Colors.NBTH = append(getproc.Settings.Colors.NBTH, getproc.ColorArray{R, G, B})
			result = `{"error":0,"result": "색상태","detail" : "블루머리색상 추가완료"}`
		case 10: //기타색상 색상추가
			//getproc.Settings.Colors.BlueTeamHead = append(getproc.Settings.Colors.BlueTeamHead,colorrgb)
			getproc.Settings.Colors.NOC = append(getproc.Settings.Colors.NOC, getproc.ColorArray{R, G, B})
			result = `{"error":0,"result": "색상태","detail" : "블루머리색상 추가완료"}`
		case 11: //기타색상 색상제거
			getproc.Settings.Colors.NOC = nil
			result = `{"error":0,"result": "색상태","detail" : "블루머리색상 전부 삭제완료"}`
		}
		saveuserdata()
		c.PureJSON(http.StatusOK, result)
	})

	router.POST("/accel", func(c *gin.Context) {
		timecheck := time.Now()

		_short := c.PostForm("short")
		_middle := c.PostForm("middle")
		_long := c.PostForm("long")
		_calc := c.PostForm("calc")
		returnval := ""
		log.Println(time.Since(timecheck))
		if _short != "" {
			val, _ := strconv.Atoi(_short)
			getproc.Settings.Accel.Short = val
			returnval = `{"error":0,"result": "가속도 설정완료","detail" : "가까이있을때 ` + _calc + ` 설정완료"}`
		}
		if _middle != "" {
			val, _ := strconv.Atoi(_middle)
			getproc.Settings.Accel.Middle = val
			returnval = `{"error":0,"result": "가속도 설정완료","detail" : "중간일때 ` + _calc + ` 설정완료"}`
		}
		if _long != "" {
			val, _ := strconv.Atoi(_long)
			getproc.Settings.Accel.Long = val
			returnval = `{"error":0,"result": "가속도 설정완료","detail" : "멀리있을때 ` + _calc + ` 설정완료"}`
		}
		log.Println(time.Since(timecheck))
		saveuserdata()
		log.Println(time.Since(timecheck))
		c.PureJSON(http.StatusOK, returnval)
	})
	router.POST("/advence", func(c *gin.Context) {

		dontmovex := c.PostForm("dontmovex")
		dontmovey := c.PostForm("dontmovey")
		minaccel := c.PostForm("minaccel")
		colordep := c.PostForm("colordep")
		returnval := ""
		if dontmovex != "" {
			val, _ := strconv.Atoi(dontmovex)
			getproc.Settings.Advance.Minx = val
			returnval = `{"error":0,"result": "고급설정 설정완료","detail" : "인식안하기 X축 ` + dontmovex + `픽셀 설정완료"}`
		}
		if dontmovey != "" {
			val, _ := strconv.Atoi(dontmovex)
			getproc.Settings.Advance.Miny = val
			returnval = `{"error":0,"result": "고급설정 설정완료","detail" : "인식안하기 y축 ` + dontmovey + `픽셀 설정완료"}`
		}
		if minaccel != "" {
			getproc.Settings.Advance.Minaccel = minaccel
			returnval = `{"error":0,"result": "고급설정 설정완료","detail" : "가속최저거리 ` + minaccel + ` 설정완료"}`
		}
		if colordep != "" {
			colordepint, _ := strconv.Atoi(colordep)
			getproc.Settings.Advance.Shade = colordep
			getproc.Settings.Advance.Shadeint = int(colordepint)
			returnval = `{"error":0,"result": "고급설정 설정완료","detail" : "색상범위 ` + colordep + ` 설정완료"}`
		}
		saveuserdata()

		c.PureJSON(http.StatusOK, returnval)
	})

	ranport, _ := net.Listen("tcp", ":0")
	_, port, _ := net.SplitHostPort(ranport.Addr().String())
	println("현재 버전 : ", version)
	fmt.Printf("프로그램 설정할경우 http://127.0.0.1:%s 로 들어가면 설정가능합니다.\n단, 꼭 크롬브라우저를 사용해야합니다.\n", port)

	//browser.OpenURL("http://127.0.0.1:" + port)
	go func() {
		exec.Command(chromepath, fmt.Sprintf("-incognito http://127.0.0.1:%s", port)).Run()
	}()
	http.Serve(ranport, router)
	//_ = router.Run(":5000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
