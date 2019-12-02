package object

const (
	//最大支持32个层级 大概
	// max layer is 32 , maybe...
	MaxLayer = 32
)

//层级管理器
type LayerManager struct {
	stringLayer map[string]int32
	intLayer    map[int32]string

	//各个层级之间的碰撞关系 是否需要确定？
	collisionSetting [][]bool
}

func NewLayerManager() *LayerManager {
	lm := &LayerManager{
		stringLayer:      make(map[string]int32),
		intLayer:         make(map[int32]string),
		collisionSetting: nil,
	}

	lm.collisionSetting = make([][]bool, MaxLayer)
	for x := 0; x < MaxLayer; x++ {
		c := make([]bool, MaxLayer)
		lm.collisionSetting[x] = c
	}

	return lm
}

func (self *LayerManager) LayerInt(layer string) int32 {
	return 0
}

func (self *LayerManager) LayerName(layer int32) string {
	return ""
}

func (self *LayerManager) AddLayer(layer string) int32 {
	return 0
}

func (self *LayerManager) RemoveLayer(layer string) {

}
