package mathg

import (
	"math"
)

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

var (
	Vector3_Zero = Vector3{
		X: 0,
		Y: 0,
		Z: 0,
	}
)

func Vector3FromXYZ(x, y, z float32) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (self *Vector3) Negative() Vector3 {
	return Vector3{
		X: -self.X,
		Y: -self.Y,
		Z: -self.Z,
	}
}

func (self *Vector3) Add(rhs Vector3) Vector3 {
	return Vector3{
		X: self.X + rhs.X,
		Y: self.Y + rhs.Y,
		Z: self.Z + rhs.Z,
	}
}

func (self *Vector3) Minus(rhs Vector3) Vector3 {
	return Vector3{
		X: self.X - rhs.X,
		Y: self.Y - rhs.Y,
		Z: self.Z - rhs.Z,
	}
}

func (self *Vector3) Normalized() Vector3 {
	length := self.Length()
	return Vector3{
		X: self.X / length,
		Y: self.Y / length,
		Z: self.Z / length,
	}
}

func (self *Vector3) Normalize() Vector3 {
	length := self.Length()
	self.X /= length
	self.Y /= length
	self.Z /= length
	return *self
}

func (self *Vector3) Multi(value float32) Vector3 {
	return Vector3{
		X: self.X * value,
		Y: self.Y * value,
		Z: self.Z * value,
	}
}

func (self *Vector3) MultiBy(value float32) Vector3 {
	self.X *= value
	self.Y *= value
	self.Z *= value
	return *self
}

func (self *Vector3) Dot(rhs Vector3) float32 {
	return self.X*rhs.X + self.Y*rhs.Y + self.Z*rhs.Z
}

func (self *Vector3) Length() float32 {
	return (float32)(math.Sqrt((float64)(self.Magitude())))
}

func (self *Vector3) Magitude() float32 {
	return self.X*self.X + self.Y*self.Y + self.Z*self.Z
}

func (self *Vector3) Cross(rhs Vector3) Vector3 {
	return Vector3{
		X: self.X*rhs.X + self.X*rhs.Y + self.X*rhs.Z,
		Y: self.Y*rhs.X + self.Y*rhs.Y + self.Y*rhs.Z,
		Z: self.Z*rhs.Z + self.Z*rhs.Y + self.Z*rhs.Z,
	}
}

//TODO
func (self *Vector3) Equals(rhs *Vector3) bool {
	return false
}
