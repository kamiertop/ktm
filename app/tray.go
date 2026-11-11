package app

import (
	"github.com/egoist/mygo"

	"github.com/kamiertop/ktm/assets"
)

func (a *App) setTray() error {
	menu := mygo.NewMenu([]*mygo.MenuItem{
		{Label: "显示窗口", Click: a.menuFunc(a.window.Show)},
		{Label: "隐藏窗口", Click: a.menuFunc(a.window.Hide)},

		mygo.Separator(),

		{Label: "退出", Click: a.menuFunc(mygo.App.Quit)},
	})
	tray, err := mygo.NewTray(mygo.TrayOptions{
		ID:             appName,
		Icon:           assets.TrayIconSVG,
		IconIsTemplate: true,
		Menu:           menu,
	})
	if err != nil {
		return err
	}
	// Linux: 双击唤出
	tray.OnClick(a.window.Show)

	a.tray = tray

	return nil
}

func (a *App) menuFunc(fn func()) func(*mygo.MenuItem, *mygo.Window) {
	return func(item *mygo.MenuItem, window *mygo.Window) {
		fn()
	}
}
