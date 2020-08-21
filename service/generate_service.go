package service

import (
	"autobullshiter/model"
	"fmt"
	"math/rand"
	"time"
)

// GenerateService 生成字符串服务
type GenerateService struct {
	Prefix1 string `form:"prefix1" json:"prefix1" binding:"required,min=1,max=64"`
	Prefix2 string `form:"prefix2" json:"prefix2" binding:"required,min=1,max=64"`
	Length  uint16 `form:"length" json:"length" binding:"required,min=1,max=32700"`
}

// Generate 生成字符串
func (service *GenerateService) Generate() interface{} {
	var str string
	t := []string{service.Prefix1, service.Prefix2}

	for i := 1; i <= int(service.Length); i++ {
		str += t[0]
		surfixes, err := model.DB.LRange(fmt.Sprintf("('%s', '%s')", t[0], t[1]), 0, -1).Result()
		if err != nil {
			return err
		}
		t[0] = t[1]
		rand.Seed(time.Now().Unix())
		i := rand.Int() % len(surfixes)
		t[1] = surfixes[i]
	}

	return str

}
