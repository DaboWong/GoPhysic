package object

type Layer struct {
	// layer Name
	layerName string

	// entity
	entity Entity

	// int value
	layerInt int32
}

func (self *Layer) Entity() Entity {
	return self.entity
}
