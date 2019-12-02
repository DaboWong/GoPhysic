package application

import "gophy/src/graphics/window"

type Application interface {
	Init()
	Window() *window.Window
}
