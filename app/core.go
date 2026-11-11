package app

import (
	"log/slog"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/kamiertop/ktm/assets"
	"github.com/kamiertop/ktm/component"
	"github.com/kamiertop/ktm/style"
)

const appName = "KTM"

type App struct {
	window *mygo.Window
	router *ui.Router
	tray   *mygo.Tray
}

func Run() (err error) {
	a := App{}
	mygo.App.WhenReady(func() {
		a.window = mygo.NewWindow(mygo.WindowOptions{
			Title:             appName,
			Width:             1200,
			Height:            880,
			MinWidth:          760,
			MinHeight:         600,
			TitleBarStyle:     mygo.TitleBarHidden,
			HideWindowButtons: true,
			Transparent:       true,
			Content:           ui.View(a.view),
			StateKey:          "main",
			URL:               "/",
		})
		a.window.OnMaximize(a.window.Invalidate)
		a.window.OnUnmaximize(a.window.Invalidate)

		// GNOME(Wayland)依赖 .desktop 文件中的图标来显示窗口图标
		if err = a.window.SetIcon(assets.AppIconPNG); err != nil {
			slog.Error("set window icon failed", "error", err)
			return
		}
		mygo.App.OnBeforeQuit(a.beforeQuit)

		if err = a.setTray(); err != nil {
			slog.Error("create system tray failed", "error", err)
			return
		}
	})

	return mygo.App.Run()
}

func (a *App) view(c *ui.Context) {
	theme := ui.DarkTheme()
	c.SetTheme(theme)

	radius := style.RadiusWindow
	borderWidth := style.BorderAppWidth
	if a.window.IsMaximized() {
		radius = 0
		borderWidth = 0
	}
	c.Root().
		Background(theme.Background).
		Radius(radius).
		Clip().
		Border(borderWidth, theme.Border)

	ui.Column(c).
		Fill().
		Children(func() {
			component.TitleBar(c, a.window)
		})
}

func (a *App) beforeQuit(e *mygo.QuitEvent) {

}
