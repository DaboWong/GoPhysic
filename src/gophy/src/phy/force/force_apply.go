package force

import "gophy/src/mathx/vec3"

type ForceApplier interface {
	// apply force
	// force: the force will apply to the object
	// forceMode: force mode
	AddForce(force vec3.Vector3, forceMode Mode)
}
