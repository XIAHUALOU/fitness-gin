package Configuration

import (
	"github.com/XIAHUALOU/fitness-gin/goft"
	"github.com/XIAHUALOU/fitness-gin/tests/internal/classes"
)

type RouterConfig struct {
	FitGin     *FitGin.FitGin      `inject:"-"`
	IndexClass *classes.IndexClass `inject:"-"`
}

func NewRouterConfig() *RouterConfig {
	return &RouterConfig{}
}
func (this *RouterConfig) IndexRoutes() interface{} {
	this.FitGin.Handle("GET", "/a", this.IndexClass.TestA)
	this.FitGin.Handle("GET", "/b", this.IndexClass.TestA)
	this.FitGin.Handle("GET", "/void", this.IndexClass.IndexVoid)
	return FitGin.Empty
}
