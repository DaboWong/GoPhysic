package core

type ForceGeneratorManager struct {
	forceElementMap map[IForceGenerator]IForceGenerator
}

func NewForceGenManager() *ForceGeneratorManager {
	return &ForceGeneratorManager{
		forceElementMap: make(map[IForceGenerator]IForceGenerator),
	}
}

func (self *ForceGeneratorManager) Add(fgen IForceGenerator) bool {
	if _, ok := self.forceElementMap[fgen]; !ok {
		return false
	}
	self.forceElementMap[fgen] = fgen
	return true
}

func (self *ForceGeneratorManager) Remove(fgen IForceGenerator) {
	delete(self.forceElementMap, fgen)
}

func (self *ForceGeneratorManager) Update(deltaTime float32) {
	for k, _ := range self.forceElementMap {
		if k != nil {
			k.UpdateForce(deltaTime)
		}
	}
}
