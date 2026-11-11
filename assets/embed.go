package assets

import (
	_ "embed"
)

var (
	//go:embed icon.png
	AppIconPNG []byte

	//go:embed tray.png
	TrayIconSVG []byte
)
