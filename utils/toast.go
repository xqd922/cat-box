package utils

import (
	"fmt"
	"path/filepath"
	"syscall"
	"unsafe"
)

// 加载 DLL 并查找指定的函数, toast.dll采用rust基于win api编写
func loadToastDLL(procName string) (*syscall.DLL, *syscall.Proc, error) {
	// 加载库
	lib, err := syscall.LoadDLL("./resources/libs/toast.dll")
	if err != nil {
		return nil, nil, fmt.Errorf("error loading DLL: %v", err)
	}

	// 查找函数
	proc, err := lib.FindProc(procName)
	if err != nil {
		lib.Release()
		return nil, nil, fmt.Errorf("error finding %s function: %v", procName, err)
	}

	return lib, proc, nil
}

// ShowToast 显示一个带有给定标题和消息的 toast 通知
func ShowToast(aumid, title, message string) error {
	lib, toastProc, err := loadToastDLL("toast")
	if err != nil {
		return err
	}
	defer lib.Release()

	// 准备输入参数
	cAumid, err := syscall.BytePtrFromString(aumid)
	if err != nil {
		return fmt.Errorf("error converting title to C string: %v", err)
	}
	cTitle, err := syscall.BytePtrFromString(title)
	if err != nil {
		return fmt.Errorf("error converting title to C string: %v", err)
	}
	cMessage, err := syscall.BytePtrFromString(message)
	if err != nil {
		return fmt.Errorf("error converting message to C string: %v", err)
	}

	// 调用函数
	ret, _, _ := toastProc.Call(uintptr(unsafe.Pointer(cAumid)), uintptr(unsafe.Pointer(cTitle)), uintptr(unsafe.Pointer(cMessage)))
	if ret == 0 {
		return fmt.Errorf("failed to show toast")
	}

	return nil
}

// HKEY_CURRENT_USER\Software\Classes\AppUserModelId\<id>
func RegisterToast(id, name, path string) error {
	lib, registerProc, err := loadToastDLL("register_toast_notification")
	if err != nil {
		return err
	}
	defer lib.Release()

	// 准备输入参数
	cId, err := syscall.BytePtrFromString(id)
	if err != nil {
		return fmt.Errorf("error converting title to C string: %v", err)
	}
	cName, err := syscall.BytePtrFromString(name)
	if err != nil {
		return fmt.Errorf("error converting message to C string: %v", err)
	}
	cPath, err := syscall.BytePtrFromString(path)
	if err != nil {
		return fmt.Errorf("error converting message to C string: %v", err)
	}

	// 调用函数
	ret, _, _ := registerProc.Call(uintptr(unsafe.Pointer(cId)), uintptr(unsafe.Pointer(cName)), uintptr(unsafe.Pointer(cPath)))
	if ret == 0 {
		return fmt.Errorf("failed to register toast")
	}

	return nil
}

// 初始化，注册通知ID, 以显示通知图标和名称
func init() {
	toastIcon, err := filepath.Abs("./resources/icons/box.ico")
	if err != nil {
		fmt.Println(err)
	}
	err = RegisterToast("cat-box", "cat-box", toastIcon)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Register toast successfully")
	}
}
