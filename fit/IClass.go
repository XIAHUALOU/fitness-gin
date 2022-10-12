package FitGin

type IClass interface {
	Build(FitGin *FitGin) //参数和方法名必须一致
	Name() string
}
