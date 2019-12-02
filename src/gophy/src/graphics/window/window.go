package window

import "gophy/src/graphics/renderer"

type Window struct {
	width, height int32
	renderer      renderer.Renderer
}

