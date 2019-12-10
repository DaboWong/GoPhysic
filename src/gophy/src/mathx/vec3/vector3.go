package vec3

import (
	"math"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func (self *Vector3) Negative() {
	self.X = -self.X
	self.Y = -self.Y
	self.Z = -self.Z
}

func (self *Vector3) Add(rhs Vector3) Vector3 {
	self.X += rhs.X
	self.Y += rhs.Y
	self.Z += rhs.Z
	return *self
}

func (self *Vector3) Sub(rhs Vector3) Vector3 {
	self.X -= rhs.X
	self.Y -= rhs.Y
	self.Z -= rhs.Z
	return *self
}

func (self Vector3) Normalized() Vector3 {
	length := self.SqrtMagnitude()
	return Vector3{
		X: self.X / length,
		Y: self.Y / length,
		Z: self.Z / length,
	}
}

func (self *Vector3) Normalize() Vector3 {
	length := self.SqrtMagnitude()
	self.X /= length
	self.Y /= length
	self.Z /= length
	return *self
}

func (self *Vector3) Multi(value float64) Vector3 {
	self.X *= value
	self.Y *= value
	self.Z *= value
	return *self
}

func (self Vector3) SqrtMagnitude() float64 {
	return math.Sqrt(self.Magnitude())
}

func (self Vector3) Magnitude() float64 {
	return self.X*self.X + self.Y*self.Y + self.Z*self.Z
}
