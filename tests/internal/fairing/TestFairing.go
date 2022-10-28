package fairing

import (
	"github.com/gin-gonic/gin"
)

type TestFairing struct{}

func NewTestFairing() *TestFairing {
	return &TestFairing{}
}

func (self *TestFairing) OnRequest(ctx *gin.Context) error {
	if name, exists := ctx.Get("name"); exists {
		name = "self is " + name.(string)
		ctx.Set("name", name)
	}
	return nil
}
func (self *TestFairing) OnResponse(ret interface{}) (interface{}, error) {
	if str, ok := ret.(string); ok {
		str = "test_" + str
		return str, nil
	}
	return ret, nil
}
