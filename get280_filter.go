package main

import (
	"time"

	"github.com/atotto/clipboard"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {

	now := time.Now().Format("200601")

	// ヘッドレスブラウザを起動する
	url := launcher.New().MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()

	// スクレイピング対象のページを指定する
	page := browser.MustPage("https://280blocker.net/files/280blocker_adblock_" + now + ".txt")
	// ページが完全にロードされるのを待つ
	page.MustWaitLoad()

	element := page.MustElement("pre")

	err := clipboard.WriteAll(element.MustText())
	if err != nil {
		return
	}

	page.MustClose()
}
