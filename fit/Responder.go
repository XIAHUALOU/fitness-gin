package FitGin

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"sync"
)

var responderList []Responder
var once_resp_list sync.Once

func get_responder_list() []Responder {
	once_resp_list.Do(func() {
		responderList = []Responder{(StringResponder)(nil),
			(JsonResponder)(nil),
			(ViewResponder)(nil),
			(SqlResponder)(nil),
			(SqlQueryResponder)(nil),
			(VoidResponder)(nil),
		}
	})
	return responderList
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

func Convert(handler interface{}) gin.HandlerFunc {
	h_ref := reflect.ValueOf(handler)
	for _, r := range get_responder_list() {
		r_ref := reflect.TypeOf(r)
		if h_ref.Type().ConvertibleTo(r_ref) {
			return h_ref.Convert(r_ref).Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type StringResponder func(*gin.Context) string

func (self StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, getFairingHandler().handlerFairing(self, context).(string))
	}
}

type Json interface{}
type JsonResponder func(*gin.Context) Json

func (self JsonResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, getFairingHandler().handlerFairing(self, context))
	}
}

type SqlQueryResponder func(*gin.Context) Query

func (self SqlQueryResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		getQuery := getFairingHandler().handlerFairing(self, context).(Query)
		ret, err := queryForMapsByInterface(getQuery)
		if err != nil {
			panic(err)
		}
		context.JSON(200, ret)
	}
}

type SqlResponder func(*gin.Context) SimpleQuery

func (self SqlResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		getSql := getFairingHandler().handlerFairing(self, context).(SimpleQuery)
		ret, err := queryForMaps(string(getSql), nil, []interface{}{}...)
		if err != nil {
			panic(err)
		}
		context.JSON(200, ret)
	}
}

type Void struct{}
type VoidResponder func(ctx *gin.Context) Void

func (self VoidResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		getFairingHandler().handlerFairing(self, context)
	}
}

// Deprecated: 暂时不提供View的解析
type View string
type ViewResponder func(*gin.Context) View

func (self ViewResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.HTML(200, string(self(context))+".html", context.Keys)
	}
}
