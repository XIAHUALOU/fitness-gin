package Configuration

import (
	"github.com/XIAHUALOU/fitness-gin/tests/internal/Services"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type MyConfig struct {
}

func NewMyConfig() *MyConfig {
	return &MyConfig{}
}
func (this *MyConfig) Test() *Services.TestService {
	return Services.NewTestService("mytest")
}
func (this *MyConfig) Naming() *Services.NameService {
	return Services.NewNameService("xiahualou")
}
func (this *MyConfig) GormDB() *gorm.DB {
	db, err := gorm.Open("mysql",
		"root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	return db
}
