package singbox

import (
	"fmt"

	"os"
	"os/exec"
	"syscall"

	"github.com/daifiyum/cat-box/config"
	"github.com/daifiyum/cat-box/utils"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/sys/windows"
)

var (
	cmd *exec.Cmd
)

func CheckConfig() error {
	cmd := exec.Command("./resources/core/sing-box.exe", "check", "-c", "./resources/core/config.json")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	if err != nil {
		output = output[21:]
		utils.ShowToast("cat-box", "错误", string(output))
		return err
	}
	return nil
}

func isProcessRunning(pid uint32) bool {
	handle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION, false, pid)
	if err != nil {
		return false
	}
	defer windows.CloseHandle(handle)

	var exitCode uint32
	err = windows.GetExitCodeProcess(handle, &exitCode)
	if err != nil {
		return false
	}

	return exitCode == 259
}

func Start() error {
	if cmd != nil && isProcessRunning(uint32(cmd.Process.Pid)) {
		return nil
	}

	err := CheckConfig()
	if err != nil {
		return err
	}

	cmd = exec.Command("./resources/core/sing-box.exe", "run", "-c", "./resources/core/config.json")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	err = cmd.Start()
	if err != nil {
		fmt.Printf("启动程序失败: %v\n", err)
		return err
	}
	return nil
}

func Stop() error {
	if cmd == nil {
		return nil
	}

	if isProcessRunning(uint32(cmd.Process.Pid)) {
		err := cmd.Process.Signal(syscall.SIGKILL)
		if err != nil {
			fmt.Printf("中断程序失败: %v\n", err)
			return err
		}
	} else {
		fmt.Println("程序已退出")
	}
	DisableProxy()
	cmd = nil
	return nil
}

func SaveConfig(data string) error {
	err := os.WriteFile(config.Config("CONFIG_PATH"), []byte(data), 0666)
	if err != nil {
		log.Error("write config failed", err)
		return err
	}
	return nil
}

func Reload(data string) error {
	err := SaveConfig(data)
	if err != nil {
		return err
	}
	err = Stop()
	if err != nil {
		log.Error("Failed to reload configuration")
		return err
	}
	err = Start()
	if err != nil {
		log.Error("Failed to reload configuration")
		return err
	}
	return nil
}
