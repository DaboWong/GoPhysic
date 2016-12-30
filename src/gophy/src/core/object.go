package core

type Object struct {
	enabled bool
	started bool
}

//object interface implements
func (self *Object) isStarted() bool {
	return self.started
}
func (self *Object) IsEnabled() bool {
	return self.enabled
}
func (self *Object) SetEnable(enable bool) {
	if self.enabled != enable {
		self.enabled = enable
	}
}
func (self *Object) setStart() {
	self.started = true
}
