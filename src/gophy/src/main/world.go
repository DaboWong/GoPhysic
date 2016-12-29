package main

import (
	"gophy/src/core"
	"log"
	"math"
)

func main() {
	v := core.Vector3FromXYZ(0.5, 4, 1)

	log.Printf("%v, %v, %v", v.X, v.Y, v.Z)

	//add

	v2 := core.Vector3FromXYZ(10, 10.5, 20)

	res := v.Add(v2)
	log.Printf("%v, %v, %v", res.X, res.Y, res.Z)

	//mul

	var rang float32 = 3.5

	v3 := v.Multi(rang)

	log.Printf("%v, %v, %v", v3.X, v3.Y, v3.Z)

	v.MultiBy(rang)

	log.Printf("%v, %v, %v", v.X, v.Y, v.Z)

	vleft := core.Vector3FromXYZ(1, 0, 0)
	vright := core.Vector3FromXYZ(0, -1, 0)

	dot := vleft.Dot(vright)

	degree := math.Acos((float64)(dot))

	log.Print(degree * 180 / math.Pi)

	nor := core.Vector3FromXYZ(3, 4, 0)

	log.Print(nor.Magitude())
	log.Print(nor.Length())
	log.Print(nor.Normalize())
}
