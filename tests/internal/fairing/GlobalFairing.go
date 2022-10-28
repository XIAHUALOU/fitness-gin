package fairing

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type GlobalFairing struct {
	DB *gorm.DB `inject:"-"`
}

func NewGlobalFairing() *GlobalFairing {
	return &GlobalFairing{}
}

func (self *GlobalFairing) OnRequest(ctx *gin.Context) error {
	ctx.Set("name", " global name ")
	return nil
}
func (self *GlobalFairing) OnResponse(ret interface{}) (interface{}, error) {
	//fmt.Println(self.DB)
	if str, ok := ret.(string); ok {
		str = str + "_global"
		return str, nil
	}
	return ret, nil
}
