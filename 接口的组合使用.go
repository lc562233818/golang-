package main

import "fmt"

type Flying interface {
	fly()
}
type Swimmer interface {
	swim()
}
type FishFly interface {
	Flying
	Swimmer
}
type Fish struct {
}

func (f Fish) fly() {
	fmt.Println("鱼会飞")
}
func (f Fish) swim() {
	fmt.Println("鱼会游泳")
}

func main() {
	var ff FishFly //定义接口
	f := Fish{}    //实例化struct
	ff = f
	ff.swim()
	ff.fly()
}
