package main

import (
	"log"
	"math"
	"reflect"
)

type Comp struct {
	Key   int32
	Value string
}

func main() {

	//log.Println("testRelfect")
	//testRelfect()

	log.Println("test world")
	testWorld()

	//test float
	//var f1 float32 = 0.000001234
	//var f2 float32 = 0.0000012341
	//log.Println(f1 == f2)
}

func testRelfect() {
	cp := &Comp{
		Key:   1,
		Value: "none",
	}
	var inter interface{} = *cp
	ty := reflect.TypeOf(inter)

	log.Println(ty.Name())
	log.Println(ty.Kind().String())
}

func testWorld() {
	world := core.NewWorld()

	world.Startup()

	world.NewGameObject()
	world.Run()
}

func testVector() {
	//unit test
	v := mathg.Vector3FromXYZ(0.5, 4, 1)

	log.Printf("%v, %v, %v", v.X, v.Y, v.Z)

	//add

	v2 := mathg.Vector3FromXYZ(10, 10.5, 20)

	res := v.Add(v2)
	log.Printf("%v, %v, %v", res.X, res.Y, res.Z)

	//mul

	var rang float32 = 3.5

	v3 := v.Multi(rang)

	log.Printf("%v, %v, %v", v3.X, v3.Y, v3.Z)

	v.MultiBy(rang)

	log.Printf("%v, %v, %v", v.X, v.Y, v.Z)

	vleft := mathg.Vector3FromXYZ(1, 0, 0)
	vright := mathg.Vector3FromXYZ(0, -1, 0)

	dot := vleft.Dot(vright)

	degree := math.Acos((float64)(dot))

	log.Print(degree * 180 / math.Pi)

	nor := mathg.Vector3FromXYZ(3, 4, 0)

	log.Print(nor.Magnitude() == 25)
	log.Print(nor.SqrtMagnitude() == 5)
	log.Print(nor.Normalize())
}
