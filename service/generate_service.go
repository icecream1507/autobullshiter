package service

import (
	"autobullshiter/model"
	"fmt"
	"math/rand"
	"time"

	"github.com/yanyiwu/gojieba"
)

// GenerateService 生成字符串服务
type GenerateService struct {
	PreString string `form:"prestr" json:"prestr" binding:"required,min=1,max=256"`
	Length    uint16 `form:"length" json:"length" binding:"required,min=1,max=32700"`
}

// Generate 生成字符串
func (service *GenerateService) Generate() interface{} {
	var str string
	jieba := gojieba.NewJieba()
	defer jieba.Free()

	w := jieba.Cut(service.PreString, true)
	pre := []string{w[len(w)-3], w[len(w)-2], w[len(w)-1]}

	for i := 1; i <= int(service.Length); i++ {
		surfixes, err := model.DB.LRange(fmt.Sprintf("('%s', '%s', '%s')", pre[0], pre[1], pre[2]), 0, -1).Result()
		if err != nil {
			return err
		}
		if len(surfixes) == 0 {
			break
		}

		pre[0] = pre[1]
		pre[1] = pre[2]
		rand.Seed(time.Now().Unix())
		j := rand.Int() % len(surfixes)
		pre[2] = surfixes[j]

		str += pre[2]
	}

	return service.PreString + str
}
