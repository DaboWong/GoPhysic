package collider

type Sphere struct {
	*Collider
}

func (self *Sphere) Radius() float32 {
	return self.size.X / 2
}
