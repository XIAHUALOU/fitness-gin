package Configuration

import (
	"github.com/XIAHUALOU/fitness-gin/fit"
	"github.com/XIAHUALOU/fitness-gin/tests/internal/classes"
)

type RouterConfig struct {
	FitGin     *FitGin.FitGin      `inject:"-"`
	IndexClass *classes.IndexClass `inject:"-"`
}

func NewRouterConfig() *RouterConfig {
	return &RouterConfig{}
}
func (self *RouterConfig) IndexRoutes() interface{} {
	self.FitGin.Handle("GET", "/a", self.IndexClass.TestA)
	self.FitGin.Handle("GET", "/b", self.IndexClass.TestA)
	self.FitGin.Handle("GET", "/void", self.IndexClass.IndexVoid)
	return FitGin.Empty
}
