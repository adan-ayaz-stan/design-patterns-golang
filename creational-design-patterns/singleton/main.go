package main

import (
	"fmt"
	"sync"
)

type MyClass struct {
	atrrib string
}

func (c *MyClass) SetAttr(v string) {
	c.atrrib = v
}

func (c *MyClass) GetAttr() string {
	return c.atrrib
}

var once sync.Once
var instance *MyClass

func GetMyClass() *MyClass {
	once.Do(func() {
		instance = &MyClass{"first"}
	})

	return instance
}

func main() {
	a := GetMyClass()
	a.SetAttr("second")
	fmt.Println(a.GetAttr())
	b := GetMyClass()
	fmt.Println(b.GetAttr())
}
