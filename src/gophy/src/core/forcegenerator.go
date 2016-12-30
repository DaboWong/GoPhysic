package core

type IForceGenerator interface {
	RegisterRigidbody(rigidbody *Rigidbody)
	UpdateForce(deltaTime float32)
}

//GravityGenerator
type GravityGenerator struct {
	Rigidbody *Rigidbody
}

func (self *GravityGenerator) RegisterRigidbody(rigidbody *Rigidbody) {

}

func (self *GravityGenerator) UpdateForce(deltaTime float32) {
	//self.Rigidbody.AddForce(Gravity.Multi(self.Rigidbody.Mass))
}

func UpdateGravity(deltaTime float32) {

}
func UpdateDamper(deltaTime float32) {

}
