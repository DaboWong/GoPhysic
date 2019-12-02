package setting

import "gophy/src/mathx/vec3"

type WorldSetting struct {
	//delta time for world
	deltaTime float32

	//gravity
	//DEFAULT = 0, -9.81, 0
	gravity vec3.Vector3

}

