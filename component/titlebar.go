package component

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/oneo-me/onevra/source/components"
	"github.com/oneo-me/onevra/source/icons/lucide"
)

func TitleBar(c *ui.Context, window *mygo.Window) {
	maximizeIcon := lucide.Square
	if window.IsMaximized() {
		maximizeIcon = lucide.Copy
	}

	ui.Row(c).
		Height(44).
		Shrink(0).
		AlignItems(ui.Center).
		Padding(0, 6, 0, 14).
		Gap(4).
		Children(func() {
			ui.Box(c).Grow(1).Height(44).DragWindow()
			titleBarButton(c.Key("minimize"), lucide.Minus, window.Minimize)
			titleBarButton(c.Key("maximize"), maximizeIcon, window.ToggleMaximize).MaximizeButton()
			titleBarButton(c.Key("close"), lucide.X, window.Close)
		})
}

type Icon = func() *ui.SVG

func titleBarButton(c *ui.Context, icon Icon, action func()) ui.Element {
	button := ui.ButtonBase(c).Size(32, 32).Radius(6).Center()
	if button.Hovered() {
		button.Background(c.Theme().SurfaceHover)
	}
	button.Children(func() {
		components.Icon(c, icon).Size(14, 14)
	})

	return button.OnClick(action)
}
