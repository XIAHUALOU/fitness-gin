package FitGin

type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc interface{}
}
type RoutesInfo []RouteInfo
type FitGinTree struct {
	trees methodTrees
}

func NewFitGinTree() *FitGinTree {
	tree := &FitGinTree{
		trees: make(methodTrees, 0, 9),
	}
	return tree
}
func (self *FitGinTree) addRoute(method, path string, handlers interface{}) {
	root := self.trees.get(method)
	if root == nil {
		root = new(node)
		root.fullPath = "/"
		self.trees = append(self.trees, methodTree{method: method, root: root})
	}
	root.addRoute(path, handlers)
}
func (self *FitGinTree) getRoute(httpMethod, path string) nodeValue {
	t := self.trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].method != httpMethod {
			continue
		}
		root := t[i].root
		// Find route in tree
		value := root.getValue(path, nil, false)
		if value.handlers != nil {
			return value
		}
	}
	return nodeValue{}
}
