package tray

import (
	"fmt"
	"os/exec"
	"syscall"

	"github.com/daifiyum/cat-box/singbox"
	"github.com/daifiyum/cat-box/utils"

	"github.com/energye/systray"
)

var (
	IsProxy   bool
	mSysProxy *systray.MenuItem
)

func init() {
	err := utils.SetProcessDPIAware()
	if err != nil {
		fmt.Println("Failed to set process DPI aware:", err)
	}
}

func SetServiceMode() {
	if mSysProxy.Checked() {
		singbox.HandleProxyMode()
	} else {
		singbox.HandleTunMode()
	}
}

func GetIsProxy() bool {
	return IsProxy
}

func OpenBrowser(url string) {
	cmd := exec.Command("cmd", "/c", "start", url)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to open browser:", err)
	}
}

func CreateItem() {
	// items
	mHome := systray.AddMenuItem("面板", "打开代理面板")
	mHome.SetIcon(HomeIcon)

	mSub := systray.AddMenuItem("订阅", "打开订阅面板")
	mSub.SetIcon(SubIcon)

	systray.AddSeparator()

	mSysProxy = systray.AddMenuItemCheckbox("系统代理", "System Proxy", true)
	mTunMode := systray.AddMenuItemCheckbox("TUN模式", "TUN Mode", false)

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("退出", "Quit the whole app")
	mQuit.SetIcon(CloseIcon)

	// click
	mHome.Click(func() {
		OpenBrowser("http://localhost:9090/ui")
	})

	mSub.Click(func() {
		OpenBrowser("http://localhost:3000")
	})

	mSysProxy.Click(func() {
		if mTunMode.Checked() {
			mTunMode.Uncheck()
			mSysProxy.Check()
		}
	})

	mTunMode.Click(func() {
		if mSysProxy.Checked() {
			mSysProxy.Uncheck()
			mTunMode.Check()
		}
	})

	mQuit.Click(func() {
		systray.Quit()
	})
}

func InitTray() {
	systray.SetIcon(AppIcon)
	systray.SetTitle("cat-box")
	systray.SetTooltip("cat-box")

	systray.SetOnClick(func(menu systray.IMenu) {
		if IsProxy {
			singbox.Stop()
			systray.SetIcon(AppIcon)
			IsProxy = false
		} else {
			SetServiceMode()
			err := singbox.Start()
			if err != nil {
				return
			}
			systray.SetIcon(ProxyIcon)
			IsProxy = true
		}
	})

	CreateItem()
}
