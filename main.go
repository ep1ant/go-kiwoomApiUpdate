package main

import (
	"fmt"
	"kiwoomUpdate/WINAPI"
	"kiwoomUpdate/inputOutputGo"
	"os/exec"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("https://github.com/sn1ezny")
	launchKiwoomAPI()

	for {
		if WINAPI.FindWindow("Open API Login") == 0 {
			fmt.Println("로그인창 대기중")
			callSleep(500)
		} else {
			callSleep(1000)
			go loginStart()
			break
		}
	}

	for {
		if WINAPI.FindWindow("opstarter") == 0 {
			fmt.Println("확인 대기중")
			callSleep(1000)
		} else {
			hwnd := WINAPI.FindWindow("opstarter")
			ClickButton(WINAPI.GetDlgItem(hwnd, 0x2))
			fmt.Println("종료 완료")
			callSleep(500)
			break
		}
	}

	callSleep(2000)

	if WINAPI.FindWindow("업그레이드 확인") != 0 {
		hwnd := WINAPI.FindWindow("업그레이드 확인")
		ClickButton(WINAPI.GetDlgItem(hwnd,0x1))
	} else {
		for {
			if WINAPI.FindWindow("Open API Login") == 0 {
				fmt.Println("창 찾는중")
				callSleep(100)
			} else {
				hwnd := WINAPI.FindWindow("Open API Login")
				ClickButton(WINAPI.GetDlgItem(hwnd,0x2))
				exitKiwoomAPI()
				callSleep(100)
				break
			}
		}
	}
}

func loginStart() {
	hwnd := WINAPI.FindWindow("Open API Login")
	if WINAPI.IsWindowEnabled(WINAPI.GetDlgItem(hwnd,0x3EA))!=1 {
		ClickButton(WINAPI.GetDlgItem(hwnd,0x3ED))
	}
	id,pw,signPw := inputOutputGo.OpenFile()
	enterKeys(WINAPI.GetDlgItem(hwnd,0x3E8),stringToUintPtr(id))
	enterKeys(WINAPI.GetDlgItem(hwnd,0x3E9),stringToUintPtr(pw))
	enterKeys(WINAPI.GetDlgItem(hwnd,0x3EA),stringToUintPtr(signPw))
	ClickButton(WINAPI.GetDlgItem(hwnd,0x1))
}

func enterKeys(hwnd, data uintptr) {
	WINAPI.SendMessage(hwnd, uintptr(WINAPI.EmSetsel),0,0xFFFFFFFF)
	// uintptr 은 음수를 받지않기때문에 0xffffffff로 처리
	WINAPI.SendMessage(hwnd, uintptr(WINAPI.EmReplacesel),0, data)
	callSleep(300)
}

func ClickButton(btn_hwnd uintptr) {
	WINAPI.PostMessage(btn_hwnd, uintptr(WINAPI.WmLbuttondown),0,0)
	callSleep(100)
	WINAPI.PostMessage(btn_hwnd, uintptr(WINAPI.WmLbuttonup),0,0)
	callSleep(300)
}

func stringToUintPtr(text string) (result uintptr) {
	result = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text)))
	return
}

func launchKiwoomAPI() {
	commandExecute("start C:\\OpenAPI\\opstarter.exe  /API")
}

func exitKiwoomAPI() {
	commandExecute("taskkill /f /im opstarter.exe")
}

func commandExecute(command string) {
	cmd := exec.Command("cmd", "/c", command)
	_ = cmd.Start()
	callSleep(200)
}

func callSleep(SleepTime time.Duration) {
	time.Sleep(SleepTime * time.Millisecond)
}