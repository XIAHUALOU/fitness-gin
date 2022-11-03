package FitGin

import (
	"fmt"
	Injector "github.com/XIAHUALOU/fitness-ioc"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
	"strings"
	"sync"
)

type Bean interface {
	Name() string
}

var Empty = &struct{}{}
var innerRouter *FitGinTree // inner tree node . backup httpmethod and path
var innerRouter_once sync.Once

func getInnerRouter() *FitGinTree {
	innerRouter_once.Do(func() {
		innerRouter = NewFitGinTree()
	})
	return innerRouter
}

type FitGin struct {
	*gin.Engine
	g            *gin.RouterGroup // 保存 group对象
	exprData     map[string]interface{}
	currentGroup string // temp-var for group string
}

func Ignite(ginMiddlewares ...gin.HandlerFunc) *FitGin {
	g := &FitGin{Engine: gin.New(),
		exprData: map[string]interface{}{},
	}
	g.Use(ErrorHandler()) //强迫加载的异常处理中间件
	for _, handler := range ginMiddlewares {
		g.Use(handler)
	}
	config := InitConfig()
	Injector.BeanFactory.Set(g)      // inject self
	Injector.BeanFactory.Set(config) // add global into (new)BeanFactory
	Injector.BeanFactory.Set(NewGPAUtil())
	if config.Server.Html != "" {
		g.LoadHTMLGlob(config.Server.Html)
	}
	return g
}
func (self *FitGin) Launch() {
	var port int32 = 8080
	if config := Injector.BeanFactory.Get((*SysConfig)(nil)); config != nil {
		port = config.(*SysConfig).Server.Port
	}
	self.applyAll()
	getCronTask().Start()
	self.Run(fmt.Sprintf(":%d", port))
}
func (self *FitGin) LaunchWithPort(port int) {

	self.applyAll()
	getCronTask().Start()
	self.Run(fmt.Sprintf(":%d", port))
}
func (self *FitGin) Handle(httpMethod, relativePath string, handler interface{}) *FitGin {
	if h := Convert(handler); h != nil {
		methods := strings.Split(httpMethod, ",")
		for _, method := range methods {
			getInnerRouter().addRoute(method, self.getPath(relativePath), h) // for future
			self.g.Handle(method, relativePath, h)
		}

	}
	return self
}
func (self *FitGin) getPath(relativePath string) string {
	g := "/" + self.currentGroup
	if g == "/" {
		g = ""
	}
	g = g + relativePath
	g = strings.Replace(g, "//", "/", -1)
	return g
}
func (self *FitGin) HandleWithFairing(httpMethod, relativePath string, handler interface{}, fairings ...Fairing) *FitGin {
	if h := Convert(handler); h != nil {
		methods := strings.Split(httpMethod, ",")
		for _, f := range fairings {
			Injector.BeanFactory.Apply(f)
		}
		for _, method := range methods {
			getInnerRouter().addRoute(method, self.getPath(relativePath), fairings) //for future
			self.g.Handle(method, relativePath, h)
		}

	}
	return self
}

// 注册中间件
func (self *FitGin) Attach(f ...Fairing) *FitGin {
	for _, f1 := range f {
		Injector.BeanFactory.Set(f1)
	}
	getFairingHandler().AddFairing(f...)
	return self
}

func (self *FitGin) Beans(beans ...Bean) *FitGin {
	for _, bean := range beans {
		self.exprData[bean.Name()] = bean
		Injector.BeanFactory.Set(bean)
	}
	return self
}
func (self *FitGin) Config(cfgs ...interface{}) *FitGin {
	Injector.BeanFactory.Config(cfgs...)
	return self
}
func (self *FitGin) applyAll() {
	for t, v := range Injector.BeanFactory.GetBeanMapper() {
		if t.Elem().Kind() == reflect.Struct {
			Injector.BeanFactory.Apply(v.Interface())
		}
	}
}

func (self *FitGin) Mount(group string, classes ...IClass) *FitGin {
	self.g = self.Group(group)
	for _, class := range classes {
		self.currentGroup = group
		class.Build(self)
		//self.beanFactory.inject(class)
		self.Beans(class)
	}
	return self
}

// 0/3 * * * * *  //增加定时任务
func (self *FitGin) Task(cron string, expr interface{}) *FitGin {
	var err error
	if f, ok := expr.(func()); ok {
		_, err = getCronTask().AddFunc(cron, f)
	} else if exp, ok := expr.(Expr); ok {
		_, err = getCronTask().AddFunc(cron, func() {
			_, expErr := ExecExpr(exp, self.exprData)
			if expErr != nil {
				log.Println(expErr)
			}
		})
	}

	if err != nil {
		log.Println(err)
	}
	return self
}
