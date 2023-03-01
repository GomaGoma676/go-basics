package main

import (
	"fmt"
	"unsafe"
)

type controller interface {
	speedUp() int
	speedDown() int
}
type vehicle struct {
	speed       int
	enginePower int
}
type bycycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}
func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}
func (b *bycycle) speedUp() int {
	b.speed += 3 * b.humanPower
	return b.speed
}
func (b *bycycle) speedDown() int {
	b.speed -= 1 * b.humanPower
	return b.speed
}
func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}
func (v vehicle) String() string {
	return fmt.Sprintf("Vehicle current speed is %v (enginePower %v)", v.speed, v.enginePower)
}

func main() {
	v := &vehicle{0, 5}
	speedUpAndDown(v)
	b := &bycycle{0, 5}
	speedUpAndDown(b)
	fmt.Println(v)

	var i1 interface{}
	var i2 any
	fmt.Printf("%[1]v %[1]T %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("%[1]v %[1]T %v\n", i2, unsafe.Sizeof(i2))
	checkType(i2)
	i2 = 1
	checkType(i2)
	i2 = "hello"
	checkType(i2)

}
func checkType(i any) {
	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
