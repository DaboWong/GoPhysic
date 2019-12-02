package force

//force type
type Mode int32

const (
	//linear force
	Linear Mode = iota

	//angular force
	Angular

	//impluse force mode
	Impulse

	//explosion force mode
	Explosion
)
