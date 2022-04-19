package WINAPI

import (
	"syscall"
	"unsafe"
)

var (
	user32, _      = syscall.LoadLibrary("user32.dll")
	messageBox, _  = syscall.GetProcAddress(user32, "MessageBoxW")
	postMessage, _ = syscall.GetProcAddress(user32, "PostMessageW")
	findWindow, _ = syscall.GetProcAddress(user32, "FindWindowW")
	getDlgItem, _ = syscall.GetProcAddress(user32, "GetDlgItem")
	isWindowEnabled, _ = syscall.GetProcAddress(user32, "IsWindowEnabled")
	sendMessage, _ = syscall.GetProcAddress(user32, "SendMessageW")


	EmSetsel      = 177
	EmReplacesel  = 194
	WmLbuttondown = 513
	WmLbuttonup   = 514
)

func SendMessage(hWnd, Msg, wParam, lParam uintptr) (result uintptr) {
	//start pos , end pos
	var nargs uintptr = 4
	ret, _, callErr := syscall.Syscall6(sendMessage,
		nargs,
		hWnd,
		Msg,
		wParam,
		lParam,
		0,
		0)
	checkError(callErr)
	result = ret
	return
}

func IsWindowEnabled(hWnd uintptr) (result uintptr) {
	var nargs uintptr = 1
	ret, _, callErr := syscall.Syscall(isWindowEnabled,
		nargs,
		hWnd,
		0,
		0,)
	checkError(callErr)
	result = ret
	return
}

func GetDlgItem(hDlg, nIDDlgItem uintptr) (result uintptr) {
	var nargs uintptr = 2
	ret, _, callErr := syscall.Syscall(getDlgItem,
		nargs,
		hDlg,
		nIDDlgItem,
		0)
	checkError(callErr)
	result = ret
	return
}

func FindWindow(caption string) (result uintptr) {
	var nargs uintptr = 2
	ret, _, callErr := syscall.Syscall(findWindow,
		nargs,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		0)
	checkError(callErr)
	result = ret
	return
}

func PostMessage(hWnd , Msg, wParam, lParam uintptr) (result int) {
	var nargs uintptr = 4
	ret, _, callErr := syscall.Syscall6(postMessage,
		nargs,
		hWnd,
		Msg,
		wParam,
		lParam,
		0,
		0)
	checkError(callErr)
	result = int(ret)
	return
}

func MessageBox(caption, text string, style uintptr) (result int) {
	var nargs uintptr = 4
	ret, _, callErr := syscall.Syscall6(messageBox,
		nargs,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		style,
		0,
		0)
	checkError(callErr)
	result = int(ret)
	return
}


func checkError(e syscall.Errno) {
	if e != 0 {
		panic(e)
	}
}