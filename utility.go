package main

import (
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
	"time"
)

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(
		chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func dirExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func getChromePath() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\chrome.exe`, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
	if err != nil {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", `https://www.google.co.kr/chrome/?brand=CHBD&gclid=CjwKCAjwk6P2BRAIEiwAfVJ0rBX_43JBD2UJq0yxVl7Gx6p4rPAxPqXVKOtXUV8vxaNaKhNtMzAbzxoCVV8QAvD_BwE&gclsrc=aw.ds`).Start()

		println("크롬 브라우저를 설치해주시기바랍니다.")
		time.Sleep(5 * time.Second)
		os.Exit(50)
		return
	}

	chromepath, _, err = k.GetStringValue("")
	k.Close()
}
func makepath(oldname, newname string) error {
	err = os.Symlink(oldname, newname)
	return err
}
