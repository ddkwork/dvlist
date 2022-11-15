package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/src/fynelib/fyneTheme"
	"github.com/hujun-open/dvlist"
)

func main() {
	data := newObject()
	list, err := dvlist.NewDVList(data)
	if !mylog.Error(err) {
		return
	}
	a := app.NewWithID("com.mitmproxy")
	//a.SetIcon(fyne.NewStaticResource("mitm", asserts.MitmBuf))
	fyneTheme.Dark()
	w := a.NewWindow("mitmproxy")
	w.Resize(fyne.NewSize(2000, 680))
	w.SetMaster()
	w.CenterOnScreen()
	w.SetContent(list)
	w.ShowAndRun()
}
