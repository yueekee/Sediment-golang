package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{
		"公司企业;公司;公司|购物服务;专卖店;专营店|购物服务;服装鞋帽皮具店;服装鞋帽皮具店",
		"公司企业;公司;广告装饰|购物服务;家居建材市场;厨卫市场",
		"公司企业;公司;商业贸易|购物服务;家居建材市场;厨卫市场",
		"公司企业;公司;公司|生活服务;旅行社;旅行社|购物服务;购物相关场所;购物相关场所",
	}

	for _, s2 := range s {
		if s2 == "" {
			continue
		}
		split := strings.Split(s2, "|")
		fmt.Println("split", split)
		fmt.Println("res:", strings.TrimPrefix(split[0], "公司企业;公司;"))
	}
}
