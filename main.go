package main

import (
	"github.com/daifiyum/cat-box/singbox"
	"github.com/daifiyum/cat-box/subservice"
	"github.com/daifiyum/cat-box/tray"
	_ "github.com/daifiyum/cat-box/utils"
	"github.com/energye/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	tray.InitTray()
	tray.CreateItem()
	subservice.SubService()
}

func onExit() {
	singbox.Stop()
}
