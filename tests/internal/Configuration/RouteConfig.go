package Configuration

import (
	"github.com/XIAHUALOU/fitness-gin/goft"
	"github.com/XIAHUALOU/fitness-gin/tests/internal/classes"
)

type RouterConfig struct {
	Goft       *goft.Goft          `inject:"-"`
	IndexClass *classes.IndexClass `inject:"-"`
}

func NewRouterConfig() *RouterConfig {
	return &RouterConfig{}
}
func (this *RouterConfig) IndexRoutes() interface{} {
	this.Goft.Handle("GET", "/a", this.IndexClass.TestA)
	this.Goft.Handle("GET", "/b", this.IndexClass.TestA)
	this.Goft.Handle("GET", "/void", this.IndexClass.IndexVoid)
	return goft.Empty
}
