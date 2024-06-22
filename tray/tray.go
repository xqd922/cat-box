package tray

import (
	"fmt"
	"os/exec"

	"github.com/daifiyum/cat-box/singbox"
	"github.com/energye/systray"
)

var (
	IsProxy  bool
	SysProxy *systray.MenuItem
)

func SetServiceMode() {
	if SysProxy.Checked() {
		singbox.HandleProxyMode()
	} else {
		singbox.HandleTunMode()
	}
}

func GetIsProxy() bool {
	return IsProxy
}

func CreateItem() {
	proxyItem := systray.AddMenuItem("面板", "打开代理面板")
	proxyItem.SetIcon(HomeIcon)
	proxyItem.Click(func() {
		url := "http://localhost:9090/ui"
		err := exec.Command("cmd", "/c", "start", url).Run()
		if err != nil {
			fmt.Println("无法打开浏览器:", err)
		}
	})

	SubItem := systray.AddMenuItem("订阅", "打开订阅面板")
	SubItem.SetIcon(SubIcon)
	SubItem.Click(func() {
		url := "http://localhost:3000"
		err := exec.Command("cmd", "/c", "start", url).Run()
		if err != nil {
			fmt.Println("无法打开浏览器:", err)
		}
	})

	systray.AddSeparator()
	SysProxy = systray.AddMenuItemCheckbox("系统代理", "System Proxy", true)
	TunMode := systray.AddMenuItemCheckbox("TUN模式", "TUN Mode", false)

	SysProxy.Click(func() {
		if TunMode.Checked() {
			TunMode.Uncheck()
			SysProxy.Check()
		}
	})

	TunMode.Click(func() {
		if SysProxy.Checked() {
			SysProxy.Uncheck()
			TunMode.Check()
		}
	})

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("退出", "Quit the whole app")
	mQuit.Enable()
	mQuit.SetIcon(CloseIcon)
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

}
