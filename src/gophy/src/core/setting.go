package core

import (
	"time"
)

type TimeSetting struct {
	TimeScale         int64     //time scale of the entire scene
	UnScaledDeltaTime int64     //unscaled delta time (millisecond)
	TimeStartup       time.Time //time when the scene start up (milliscond)
	*time.Timer
}

//time per frame
func (self *TimeSetting) DeltaTime() int64 {
	return self.TimeScale * self.UnScaledDeltaTime
}

func (self *TimeSetting) resetTimer() {
	self.Timer = time.NewTimer(time.Millisecond + (time.Duration)(self.DeltaTime()))
}
