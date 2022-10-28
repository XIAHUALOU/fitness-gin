package Services

import "fmt"

type NameService struct {
	MyName string
}

func NewNameService(myName string) *NameService {
	return &NameService{MyName: myName}
}
func (self *NameService) ShowName() {
	fmt.Println(self.MyName)
}
