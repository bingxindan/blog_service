package main

import (
	"blog_service/model"
	"fmt"
)

func main() {
	fmt.Println(11)
	obj := new(model.BlogModel)
	ret, err := obj.FindOne(1)
	fmt.Println(ret, "##", err)
}
